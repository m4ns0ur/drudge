package main

import (
	"io/ioutil"
	"path/filepath"

	"github.com/willfaught/drudge/drudge"
)

func init() {
	w.Work("install", "cron", func() error {
		var crontab = filepath.Join(drudge.Files, ".crontab")
		var bs, err = ioutil.ReadFile(crontab)

		if err != nil {
			return err
		}

		var expected = string(bs)
		actual, err := w.RunOutput("crontab", "-l")

		if err != nil {
			return err
		}

		if actual != expected {
			if err := w.Run("crontab", "-r"); err != nil {
				return err
			}

			if err := w.Run("crontab", crontab); err != nil {
				return err
			}
		}

		return w.Stow(".crontab")
	})

	w.Work("uninstall", "cron", func() error {
		if err := w.Run("crontab", "-r"); err != nil {
			return err
		}

		return w.Pack(".crontab")
	})
}
