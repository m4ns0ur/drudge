package main

import (
	"os"
	"path/filepath"

	"github.com/willfaught/drudge/drudge"
)

func init() {
	w.Work("install", "vim", func() error {
		if _, err := os.Stat(filepath.Join(drudge.Home, ".vim")); os.IsNotExist(err) {
			if err := w.Run("sh", "-c", "curl -fsLS https://bit.ly/janus-bootstrap | bash"); err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		return nil
	})

	w.Work("uninstall", "vim", func() error {
		if err := os.RemoveAll(filepath.Join(drudge.Home, ".vim")); err != nil {
			return err
		}

		return nil
	})
}
