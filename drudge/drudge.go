// Package drudge defines works, the dependencies between them, and their
// dependencies on software operating systems and hardware architectures, then
// does selected works and their dependencies in dependency order, where works
// depended on by others are done first. A work is a function that returns an
// error. It is identified by a unique combination of a verb and an object, such
// as install foo or upgrade bar.
package drudge

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

// Files is the path to the files directory in this package.
var Files string

// Home is the home directory path of the current user.
var Home string

var systems = map[System]struct{}{System(runtime.GOARCH): {}, System(runtime.GOOS): {}}

func dir(path string) (bool, error) {
	var f, err = os.Lstat(path)

	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return f.IsDir(), nil
}

func exists(path string) (bool, error) {
	var _, err = os.Lstat(path)

	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func file(path string) (bool, error) {
	var f, err = os.Lstat(path)

	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return !f.IsDir(), nil
}

func init() {
	var u, err = user.Current()

	if err != nil {
		panic(err)
	}

	var _, f, _, _ = runtime.Caller(1)

	Files = filepath.Join(filepath.Dir(filepath.Dir(f)), "files")
	Home = u.HomeDir
}

func topological(unsorted map[*work]struct{}) ([]*work, error) {
	if len(unsorted) == 0 {
		return nil, nil
	}

	var indep []*work

	for t := range unsorted {
		if len(t.afterworks) == 0 {
			indep = append(indep, t)
		}
	}

	var sorted []*work

	for len(indep) > 0 {
		var t = indep[0]

		indep = indep[1:]
		sorted = append(sorted, t)

		for b := range t.beforeworks {
			if _, ok := unsorted[b]; !ok {
				continue
			}

			delete(b.afterworks, t)

			if len(b.afterworks) == 0 {
				indep = append(indep, b)
			}
		}
	}

	if len(sorted) != len(unsorted) {
		for _, t := range sorted {
			delete(unsorted, t)
		}

		var join []string

		for t := range unsorted {
			join = append(join, fmt.Sprintf("%s %s", t.verb, t.object))
		}

		sort.Strings(join)

		var plural string

		if len(join) > 1 {
			plural = "s"
		}

		return nil, fmt.Errorf("dependency cycle for work%s %s", plural, strings.Join(join, ", "))
	}

	return sorted, nil
}

func symlink(path string) (bool, error) {
	var fi, err = os.Lstat(path)

	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	if fi.Mode()&os.ModeSymlink != os.ModeSymlink {
		return false, nil
	}

	return true, nil
}

// Worker does work defined as functions and identified by verbs and objects.
type Worker struct {
	// Dry is whether changes should be made.
	Dry bool

	// Output is the log output. Defaults to os.Stdout.
	Output io.Writer

	// Quiet is whether the log should be printed.
	Quiet bool

	// Verbose is whether the full log should be printed.
	Verbose bool

	verbobjects map[string]map[string]*work
}

// Backup copies paths to the backup directory under the files directory.
func (w *Worker) Backup(paths ...string) error {
	var b = filepath.Join(Files, "backup")

	if d, err := dir(b); err != nil {
		return err
	} else if !d {
		if err := w.Run("mkdir", "-p", b); err != nil {
			return err
		}
	}

	for _, p := range paths {
		var from = filepath.Join(Home, p)
		var to = filepath.Join(b, p)
		var todir = filepath.Dir(to)

		if d, err := dir(todir); err != nil {
			return err
		} else if !d {
			if err := w.Run("mkdir", "-p", todir); err != nil {
				return err
			}
		}

		if err := w.Run("cp", "-f", from, to); err != nil {
			return err
		}
	}

	return nil
}

// Do does the work identified by verb and objects.
func (w *Worker) Do(verb string, objects ...string) error {
	w.Log("dry is %t", w.Dry)
	w.Log("quiet is %t", w.Quiet)
	w.Log("verbose is %t", w.Verbose)
	w.resolve()

	var visit, err = w.search(verb, objects)

	if err != nil {
		return err
	}

	var unsorted = map[*work]struct{}{}

	for len(visit) > 0 {
		var t = visit[0]

		visit = visit[1:]

		if _, ok := unsorted[t]; ok {
			continue
		}

		unsorted[t] = struct{}{}

		for a := range t.afterworks {
			if _, ok := unsorted[a]; ok {
				continue
			}

			visit = append(visit, a)
		}
	}

	sorted, err := topological(unsorted)

	if err != nil {
		return err
	}

	for _, t := range sorted {
		w.Log("%s %s", verb, t.object)

		if err := t.do(); err != nil {
			return err
		}
	}

	if !w.Quiet && !w.Verbose {
		fmt.Println()
	}

	return nil
}

// Log logs with format and args only if Quiet is false and Verbose is true.
func (w *Worker) Log(format string, args ...interface{}) {
	var o = w.Output

	if o == nil {
		o = os.Stdout
	}

	if !w.Quiet && w.Verbose {
		fmt.Fprintf(o, "%s: %s\n", os.Args[0], fmt.Sprintf(format, args...))
	}
}

// Pack restores the original Files for paths.
func (w *Worker) Pack(paths ...string) error {
	for _, p := range paths {
		var from = filepath.Join(Home, p)
		var to = filepath.Join(Files, p)
		var e, err = exists(from)

		if err != nil {
			return err
		}

		if !e {
			return fmt.Errorf("%s does not exist", from)
		}

		f, err := file(from)

		if err != nil {
			return err
		}

		if !f {
			return fmt.Errorf("%s is not a file", from)
		}

		s, err := symlink(from)

		if err != nil {
			return err
		}

		if !s {
			return fmt.Errorf("%s is not a symlink", from)
		}

		l, err := os.Readlink(from)

		if err != nil {
			return err
		}

		if l != to {
			return fmt.Errorf("%s does not symlink %s", from, to)
		}

		var backupfile = filepath.Join(Files, "backup", p)

		e, err = exists(backupfile)

		if err != nil {
			return err
		}

		if e {
			w.Log("%s exists", backupfile)
			w.Log("rename %s to %s", backupfile, from)

			if err := os.Rename(backupfile, from); err != nil {
				return err
			}
		} else {
			w.Log("%s does not exist", backupfile)
		}

		w.Log("remove %s", from)

		if err := os.Remove(from); err != nil {
			return err
		}

		w.Progress()
	}

	return nil
}

// Progress prints a period with no newline to indicate progress if Quiet and
// Verbose are false.
func (w *Worker) Progress() {
	if !w.Quiet && !w.Verbose {
		fmt.Print(".")
	}
}

// Run executes the command and arguments. It returns early if Dry is true.
func (w *Worker) Run(command string, args ...string) error {
	w.Log("command: %s %s", command, strings.Join(args, " "))

	if w.Dry {
		return nil
	}

	var c = exec.Command(command, args...)
	var bs, err = c.CombinedOutput()

	w.Progress()

	if err != nil {
		return fmt.Errorf("%s: %s", err, string(bs))
	}

	return nil
}

// RunMany makes many Run calls and stops after the first error.
func (w *Worker) RunMany(commands [][]string) error {
	for _, c := range commands {
		if err := w.Run(c[0], c[1:]...); err != nil {
			return err
		}
	}

	return nil
}

// RunOutput executes the command and arguments and returns the combined
// standard and error outputs.
func (w *Worker) RunOutput(command string, args ...string) (string, error) {
	w.Log("command: %s %s", command, strings.Join(args, " "))

	var c = exec.Command(command, args...)
	var bs, err = c.CombinedOutput()
	var s = string(bs)

	if err != nil {
		return "", fmt.Errorf("%s: %s", err, s)
	}

	return s, nil
}

// RunTry executes the command and arguments and returns whether it succeeds.
func (w *Worker) RunTry(command string, args ...string) bool {
	w.Log("command: %s %s", command, strings.Join(args, " "))

	var c = exec.Command(command, args...)

	c.Run()

	return c.ProcessState.Success()
}

// Stow creates symbolic links from paths relative to the repository Files
// directory to paths relative to the Home directory and backs up the original
// Files.
func (w *Worker) Stow(paths ...string) error {
	for _, p := range paths {
		var from = filepath.Join(Home, p)
		var to = filepath.Join(Files, p)
		var e, err = exists(from)

		if err != nil {
			return err
		}

		if e {
			var s, err = symlink(from)

			if err != nil {
				return err
			}

			if s {
				var l, err = os.Readlink(from)

				if err != nil {
					return err
				}

				if l == to {
					w.Log("already symlinked from %s to %s", from, to)

					continue
				}
			}

			var backupfile = filepath.Join(Files, "backup", p)
			var backupdir = filepath.Dir(backupfile)
			d, err := dir(backupdir)

			if err != nil {
				return err
			}

			if !d {
				if err := os.MkdirAll(backupdir, 0700); err != nil {
					return err
				}

				w.Log("made directory %s", backupdir)
			}

			if err := os.Rename(from, backupfile); err != nil {
				return err
			}

			w.Log("moved file from %s to %s", from, backupfile)
		}

		if err := os.Symlink(to, from); err != nil {
			return err
		}

		w.Log("symlinked file from %s to %s", from, to)
		w.Progress()
	}

	return nil
}

// Work will call do for verb and object augmented by options.
func (w *Worker) Work(verb, object string, do func() error, options ...Option) {
	if w.verbobjects == nil {
		w.verbobjects = map[string]map[string]*work{}
	}

	if w.verbobjects[verb] == nil {
		w.verbobjects[verb] = map[string]*work{}
	}

	if _, ok := w.verbobjects[verb][object]; ok {
		panic(object)
	}

	var t = &work{
		afternames:  map[string][]string{},
		afterworks:  map[*work]struct{}{},
		beforenames: map[string][]string{},
		beforeworks: map[*work]struct{}{},
		do:          do,
		object:      object,
		systems:     map[System]struct{}{},
		verb:        verb,
	}

	for _, o := range options {
		o(t)
	}

	w.verbobjects[verb][object] = t
}

func (w *Worker) resolve() error {
	for _, os := range w.verbobjects {
		for _, t := range os {
			for v, os := range t.afternames {
				for _, o := range os {
					var before, ok = w.verbobjects[v][o]

					if !ok {
						return fmt.Errorf("work %s %s is undefined", v, o)
					}

					t.afterworks[before] = struct{}{}
					before.beforeworks[t] = struct{}{}
				}
			}

			for v, os := range t.beforenames {
				for _, o := range os {
					var after, ok = w.verbobjects[v][o]

					if !ok {
						return fmt.Errorf("work %s %s is undefined", v, o)
					}

					t.beforeworks[after] = struct{}{}
					after.afterworks[t] = struct{}{}
				}
			}
		}
	}

	return nil
}

func (w *Worker) search(verb string, objects []string) ([]*work, error) {
	var all = w.verbobjects[verb]

	if len(all) == 0 {
		return nil, fmt.Errorf("verb %s is invalid", verb)
	}

	var visit []*work

	if len(objects) == 0 {
		for _, t := range all {
			var compat = true

			for s := range t.systems {
				if _, ok := systems[s]; !ok {
					compat = false

					break
				}
			}

			if compat {
				visit = append(visit, t)
			}
		}
	} else {
		for _, o := range objects {
			var t, ok = all[o]

			if !ok {
				return nil, fmt.Errorf("object %s is invalid", o)
			}

			var ss []string

			for s := range t.systems {
				ss = append(ss, string(s))
			}

			sort.Strings(ss)

			for _, s := range ss {
				if _, ok := systems[System(s)]; !ok {
					return nil, fmt.Errorf("work %s %s requires system %s", verb, o, s)
				}
			}

			visit = append(visit, t)
		}
	}

	return visit, nil
}

// Option defines work and system dependencies.
type Option func(*work)

// After adds a dependency to a verb and object.
func After(verb, object string) Option {
	return func(t *work) {
		t.afternames[verb] = append(t.afternames[verb], object)
	}
}

// Before adds a dependency from a verb and object.
func Before(verb, object string) Option {
	return func(t *work) {
		t.beforenames[verb] = append(t.beforenames[verb], object)
	}
}

// Require adds a dependency on one or more Systems.
func Require(ss ...System) Option {
	return func(t *work) {
		for _, s := range ss {
			t.systems[s] = struct{}{}
		}
	}
}

// System is a hardware architecture or software operating system dependency.
type System string

// The software operating systems.
const (
	Android   System = "android"
	Darwin           = "darwin"
	Dragonfly        = "dragonfly"
	Freebsd          = "freebsd"
	Linux            = "linux"
	Nacl             = "nacl"
	Netbsd           = "netbsd"
	Openbsd          = "openbsd"
	Plan9            = "plan9"
	Solaris          = "solaris"
	Windows          = "windows"
	Zos              = "zos"
)

// The hardware architectures.
const (
	Amd64       System = "amd64"
	Amd64p32           = "amd64p32"
	Arm                = "arm"
	Arm64              = "arm64"
	Arm64be            = "arm64be"
	Armbe              = "armbe"
	I386               = "386"
	Mips               = "mips"
	Mips64             = "mips64"
	Mips64le           = "mips64le"
	Mips64p32          = "mips64p32"
	Mips64p32le        = "mips64p32le"
	Mipsle             = "mipsle"
	Ppc                = "ppc"
	Ppc64              = "ppc64"
	Ppc64le            = "ppc64le"
	S390               = "s390"
	S390x              = "s390x"
	Sparc              = "sparc"
	Sparc64            = "sparc64"
)

type work struct {
	afternames  map[string][]string
	afterworks  map[*work]struct{}
	beforenames map[string][]string
	beforeworks map[*work]struct{}
	do          func() error
	object      string
	systems     map[System]struct{}
	verb        string
}
