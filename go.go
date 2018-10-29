package main

import (
	"os"

	"github.com/willfaught/drudge/drudge"
)

func init() {
	var packages = []string{
		"github.com/alecthomas/gometalinter",
		"golang.org/x/lint/golint",
		"github.com/govend/govend",
		"github.com/mailgun/godebug",
		"github.com/nsf/gocode",
		"github.com/rogpeppe/godef",
		"github.com/zmb3/gogetdoc",
		"golang.org/x/tools/cmd/eg",
		"golang.org/x/tools/cmd/goimports",
		"golang.org/x/tools/cmd/gorename",
		"golang.org/x/tools/cmd/guru",
		"golang.org/x/tools/cmd/present",
	}

	w.Work("install", "go", func() error {
		return w.Run("go", append([]string{"get"}, packages...)...)
	}, drudge.After("install", "homebrew"))

	w.Work("uninstall", "go", func() error {
		if p := os.Getenv("GOPATH"); p != "" {
			return os.RemoveAll(p)
		}

		return nil
	}, drudge.Before("uninstall", "homebrew"))

	w.Work("upgrade", "go", func() error {
		return w.Run("go", append([]string{"get", "-u"}, packages...)...)
	}, drudge.After("upgrade", "homebrew"))
}
