package drudge

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestWorker_basic(t *testing.T) {
	t.Parallel()

	var w = Worker{Quiet: true}
	var aa, ab, x bool

	w.Work("a", "a", func() error {
		aa = true

		return nil
	})

	w.Work("a", "b", func() error {
		ab = true

		return nil
	})

	w.Work("", "", func() error {
		x = true

		return nil
	})

	if err := w.Do("a", "a"); err != nil {
		t.Errorf("actual %v, expected nil", err)
	}

	if err := w.Do("a", "b"); err != nil {
		t.Errorf("actual %v, expected nil", err)
	}

	if err := w.Do("", ""); err != nil {
		t.Errorf("actual %v, expected nil", err)
	}

	for i, b := range []bool{aa, ab, x} {
		if !b {
			t.Errorf("%d: actual false, expected true", i)
		}
	}

	aa, ab = false, false

	if err := w.Do("a", "a", "b"); err != nil {
		t.Errorf("actual %v, expected nil", err)
	}

	for i, b := range []bool{aa, ab} {
		if !b {
			t.Errorf("%d: actual false, expected true", i)
		}
	}
}

func TestWorker_panic(t *testing.T) {
	t.Parallel()

	var w = Worker{Quiet: true}

	var panicked = func(f func()) (panicked bool) {
		defer func() {
			if v := recover(); v != nil {
				panicked = true
			}
		}()

		f()

		return false
	}

	w.Work("a", "a", nil)

	if !panicked(func() { w.Work("a", "a", nil) }) {
		t.Errorf("actual no panic, expected panic")
	}

	if !panicked(func() { w.Do("a", "a") }) {
		t.Errorf("actual no panic, expected panic")
	}

	w.Work("a", "b", func() error {
		panic("a b")
	})

	if !panicked(func() { w.Do("a", "b") }) {
		t.Errorf("actual no panic, expected panic")
	}
}

func TestWorker_Work_cycledeps(t *testing.T) {
	t.Parallel()

	var w = Worker{Quiet: true}

	w.Work("a", "a", func() error { return nil }, After("a", "a"))

	if err := w.Do("a", "a"); err == nil {
		t.Errorf("actual nil, expected not nil")
	} else if a, e := err.Error(), "dependency cycle for work a a"; a != e {
		t.Errorf("actual %s, expected %s", a, e)
	}

	w = Worker{Quiet: true}

	w.Work("a", "a", func() error { return nil }, After("a", "b"))
	w.Work("a", "b", func() error { return nil }, After("a", "a"))

	if err := w.Do("a", "a"); err == nil {
		t.Errorf("actual nil, expected not nil")
	} else if a, e := err.Error(), "dependency cycle for works a a, a b"; a != e {
		t.Errorf("actual %s, expected %s", a, e)
	}

	if err := w.Do("a", "b"); err == nil {
		t.Errorf("actual nil, expected not nil")
	} else if a, e := err.Error(), "dependency cycle for works a a, a b"; a != e {
		t.Errorf("actual %s, expected %s", a, e)
	}

	w = Worker{Quiet: true}

	w.Work("a", "a", func() error { return nil }, After("a", "b"))
	w.Work("a", "b", func() error { return nil }, After("a", "c"))
	w.Work("a", "c", func() error { return nil }, After("a", "a"))

	if err := w.Do("a", "a"); err == nil {
		t.Errorf("actual nil, expected not nil")
	} else if a, e := err.Error(), "dependency cycle for works a a, a b, a c"; a != e {
		t.Errorf("actual %s, expected %s", a, e)
	}

	if err := w.Do("a", "b"); err == nil {
		t.Errorf("actual nil, expected not nil")
	} else if a, e := err.Error(), "dependency cycle for works a a, a b, a c"; a != e {
		t.Errorf("actual %s, expected %s", a, e)
	}

	if err := w.Do("a", "c"); err == nil {
		t.Errorf("actual nil, expected not nil")
	} else if a, e := err.Error(), "dependency cycle for works a a, a b, a c"; a != e {
		t.Errorf("actual %s, expected %s", a, e)
	}
}

func TestWorker_Work_workdeps(t *testing.T) {
	t.Parallel()

	var w = Worker{Quiet: true}
	var i []int

	w.Work("a", "a", func() error {
		i = append(i, 0)

		return nil
	})

	w.Work("a", "b", func() error {
		i = append(i, 1)

		return nil
	}, After("a", "a"))

	if err := w.Do("a", "a"); err != nil {
		t.Errorf("actual %v, expected nil", err)
	}

	if a, e := i, []int{0}; !reflect.DeepEqual(a, e) {
		t.Errorf("actual %v, expected %v", a, e)
	}

	i = nil

	if err := w.Do("a", "b"); err != nil {
		t.Errorf("actual %v, expected nil", err)
	}

	if a, e := i, []int{0, 1}; !reflect.DeepEqual(a, e) {
		t.Errorf("actual %v, expected %v", a, e)
	}

	w = Worker{Quiet: true}
	i = nil

	w.Work("a", "b", func() error {
		i = append(i, 1)

		return nil
	}, After("a", "a"))

	w.Work("a", "a", func() error {
		i = append(i, 0)

		return nil
	})

	if err := w.Do("a", "a"); err != nil {
		t.Errorf("actual %v, expected nil", err)
	}

	if a, e := i, []int{0}; !reflect.DeepEqual(a, e) {
		t.Errorf("actual %v, expected %v", a, e)
	}

	i = nil

	if err := w.Do("a", "b"); err != nil {
		t.Errorf("actual %v, expected nil", err)
	}

	if a, e := i, []int{0, 1}; !reflect.DeepEqual(a, e) {
		t.Errorf("actual %v, expected %v", a, e)
	}

	w = Worker{Quiet: true}
	i = nil

	w.Work("a", "a", func() error {
		i = append(i, 0)

		return nil
	})

	w.Work("a", "b", func() error {
		i = append(i, 1)

		return nil
	}, After("a", "a"))

	w.Work("a", "c", func() error {
		i = append(i, 2)

		return nil
	}, After("a", "b"))

	if err := w.Do("a", "a"); err != nil {
		t.Errorf("actual %v, expected nil", err)
	}

	if a, e := i, []int{0}; !reflect.DeepEqual(a, e) {
		t.Errorf("actual %v, expected %v", a, e)
	}

	i = nil

	if err := w.Do("a", "b"); err != nil {
		t.Errorf("actual %v, expected nil", err)
	}

	if a, e := i, []int{0, 1}; !reflect.DeepEqual(a, e) {
		t.Errorf("actual %v, expected %v", a, e)
	}

	i = nil

	if err := w.Do("a", "c"); err != nil {
		t.Errorf("actual %v, expected nil", err)
	}

	if a, e := i, []int{0, 1, 2}; !reflect.DeepEqual(a, e) {
		t.Errorf("actual %v, expected %v", a, e)
	}

	w = Worker{Quiet: true}
	i = nil

	w.Work("a", "c", func() error {
		i = append(i, 2)

		return nil
	}, After("a", "b"))

	w.Work("a", "b", func() error {
		i = append(i, 1)

		return nil
	}, After("a", "a"))

	w.Work("a", "a", func() error {
		i = append(i, 0)

		return nil
	})

	if err := w.Do("a", "a"); err != nil {
		t.Errorf("actual %v, expected nil", err)
	}

	if a, e := i, []int{0}; !reflect.DeepEqual(a, e) {
		t.Errorf("actual %v, expected %v", a, e)
	}

	i = nil

	if err := w.Do("a", "b"); err != nil {
		t.Errorf("actual %v, expected nil", err)
	}

	if a, e := i, []int{0, 1}; !reflect.DeepEqual(a, e) {
		t.Errorf("actual %v, expected %v", a, e)
	}

	i = nil

	if err := w.Do("a", "c"); err != nil {
		t.Errorf("actual %v, expected nil", err)
	}

	if a, e := i, []int{0, 1, 2}; !reflect.DeepEqual(a, e) {
		t.Errorf("actual %v, expected %v", a, e)
	}
}

func TestWorker_Work_sysdeps(t *testing.T) {
	t.Parallel()

	var w = Worker{Quiet: true}

	systems = map[System]struct{}{Darwin: {}, Amd64: {}}

	w.Work("a", "a", func() error { return nil })

	if err := w.Do("a", "a"); err != nil {
		t.Errorf("actual %v, expected nil", err)
	}

	w.Work("a", "b", func() error { return nil }, Require(Darwin))

	if err := w.Do("a", "b"); err != nil {
		t.Errorf("actual %v, expected nil", err)
	}

	w.Work("a", "c", func() error { return nil }, Require(Amd64))

	if err := w.Do("a", "c"); err != nil {
		t.Errorf("actual %v, expected nil", err)
	}

	w.Work("a", "d", func() error { return nil }, Require(Amd64), Require(Darwin))

	if err := w.Do("a", "d"); err != nil {
		t.Errorf("actual %v, expected nil", err)
	}

	w.Work("a", "e", func() error { return nil }, Require(Amd64, Darwin))

	if err := w.Do("a", "e"); err != nil {
		t.Errorf("actual %v, expected nil", err)
	}

	w.Work("a", "f", func() error { return nil }, Require(Windows))

	if err := w.Do("a", "f"); err == nil {
		t.Errorf("actual nil, expected not nil")
	} else if a, e := err.Error(), "work a f requires system windows"; a != e {
		t.Errorf("actual %s, expected %s", a, e)
	}

	w.Work("a", "g", func() error { return nil }, Require(Mips))

	if err := w.Do("a", "g"); err == nil {
		t.Errorf("actual nil, expected not nil")
	} else if a, e := err.Error(), "work a g requires system mips"; a != e {
		t.Errorf("actual %s, expected %s", a, e)
	}

	w.Work("a", "h", func() error { return nil }, Require(Mips), Require(Windows))

	if err := w.Do("a", "h"); err == nil {
		t.Errorf("actual nil, expected not nil")
	} else if a, e := err.Error(), "work a h requires system mips"; a != e {
		t.Errorf("actual %s, expected %s", a, e)
	}

	w.Work("a", "i", func() error { return nil }, Require(Mips, Windows))

	if err := w.Do("a", "i"); err == nil {
		t.Errorf("actual nil, expected not nil")
	} else if a, e := err.Error(), "work a i requires system mips"; a != e {
		t.Errorf("actual %s, expected %s", a, e)
	}

	w.Work("a", "j", func() error { return nil }, Require(Amd64), Require(Darwin), Require(Windows))

	if err := w.Do("a", "j"); err == nil {
		t.Errorf("actual nil, expected not nil")
	} else if a, e := err.Error(), "work a j requires system windows"; a != e {
		t.Errorf("actual %s, expected %s", a, e)
	}

	w.Work("a", "k", func() error { return nil }, Require(Amd64), Require(Mips), Require(Windows))

	if err := w.Do("a", "k"); err == nil {
		t.Errorf("actual nil, expected not nil")
	} else if a, e := err.Error(), "work a k requires system mips"; a != e {
		t.Errorf("actual %s, expected %s", a, e)
	}

	w.Work("a", "l", func() error { return nil }, Require(Windows), Require(Mips), Require(Amd64))

	if err := w.Do("a", "l"); err == nil {
		t.Errorf("actual nil, expected not nil")
	} else if a, e := err.Error(), "work a l requires system mips"; a != e {
		t.Errorf("actual %s, expected %s", a, e)
	}
}

func TestWorker_Backup(t *testing.T) {
	t.Parallel()

	var w = Worker{Quiet: true}
	var f, err = ioutil.TempFile(Home, ".")

	if err != nil {
		t.Fatal(err)
	}

	var base = filepath.Base(f.Name())

	const e = "test\n"

	if _, err := fmt.Fprint(f, e); err != nil {
		t.Fatal(err)
	}

	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

	var existed = true
	var backupdir = filepath.Join(Files, "backup")

	if _, err = os.Lstat(backupdir); err != nil {
		if os.IsNotExist(err) {
			existed = false
		} else {
			t.Fatal(err)
		}
	}

	if err := w.Backup(base); err != nil {
		t.Fatal(err)
	}

	if err := os.Remove(f.Name()); err != nil {
		t.Fatal(err)
	}

	var backupfile = filepath.Join(backupdir, base)
	bs, err := ioutil.ReadFile(backupfile)

	if err != nil {
		t.Fatal(err)
	}

	if a, e := string(bs), e; a != e {
		t.Errorf("actual %s, expected %s", a, e)
	}

	if err := os.Remove(backupfile); err != nil {
		t.Fatal(err)
	}

	if !existed {
		if err := os.RemoveAll(backupdir); err != nil {
			t.Fatal(err)
		}
	}
}

/*func TestWorker_Pack(t *testing.T) {
	t.Parallel()

	// Preserve any existing backup directory.
	var backupexists = true
	var backupdir = filepath.Join(Files, "backup")
	var err error

	if _, err = os.Lstat(backupdir); err != nil {
		if os.IsNotExist(err) {
			backupexists = false
		} else {
			t.Fatal(err)
		}
	}

	var tempbackupdir, hiddenbackupdir string

	if backupexists {
		if tempbackupdir, err = ioutil.TempDir(Files, "."); err != nil {
			t.Fatal(err)
		}

		hiddenbackupdir = filepath.Join(tempbackupdir, "backup")

		if err := os.Rename(backupdir, hiddenbackupdir); err != nil {
			t.Fatal(err)
		}
	}

	// No backup directory or file.
	var w = Worker{Quiet: true}
	f, err := ioutil.TempFile(Home, ".")

	if err != nil {
		t.Fatal(err)
	}

	var base = filepath.Base(f.Name())

	if _, err := fmt.Fprint(f, "new\n"); err != nil {
		t.Fatal(err)
	}

	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

	if err := w.Pack(base); err != nil {
		t.Fatal(err)
	}

	// Backup directory.
	if err := os.MkdirAll(backupdir, 0600); err != nil {
		t.Fatal(err)
	}

	if err := w.Pack(base); err != nil {
		t.Fatal(err)
	}

	// Backup file.
	var backupfile = filepath.Join(backupdir, base)

	if err := ioutil.WriteFile(backupfile, []byte("old\n"), 0600); err != nil {
		t.Fatal(err)
	}

	if err := w.Pack(base); err != nil {
		t.Fatal(err)
	}

	if bs, err := ioutil.ReadFile(f.Name()); err != nil {
		t.Fatal(err)
	} else if a, e := string(bs), "old\n"; a != e {
		t.Errorf("actual %s, expected %s", a, e)
	}

	if err := os.Remove(f.Name()); err != nil {
		t.Fatal(err)
	}

	if err := os.Remove(backupfile); err != nil {
		t.Fatal(err)
	}

	// No backup directory.
	if err := os.RemoveAll(backupdir); err != nil {
		t.Fatal(err)
	}
}*/
