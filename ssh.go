package main

import (
	"os"
	"path/filepath"

	"github.com/willfaught/drudge/drudge"
)

func init() {
	w.Work("install", "ssh", func() error {
		var idrsa, idrsapub bool

		if _, err := os.Stat(filepath.Join(drudge.Home, ".ssh", "id_rsa")); err == nil {
			idrsa = true
		} else if !os.IsNotExist(err) {
			return err
		}

		if _, err := os.Stat(filepath.Join(drudge.Home, ".ssh", "id_rsa.pub")); err == nil {
			idrsapub = true
		} else if !os.IsNotExist(err) {
			return err
		}

		if !idrsa || !idrsapub {
			if idrsa {
				if err := w.Backup(filepath.Join(drudge.Home, ".ssh", "id_rsa")); err != nil {
					return err
				}
			}

			if idrsapub {
				if err := w.Backup(filepath.Join(drudge.Home, ".ssh", "id_rsa.pub")); err != nil {
					return err
				}
			}

			if err := w.Run("ssh-keygen", "-t", "rsa", "-N", ""); err != nil {
				return err
			}
		}

		return nil
	}, drudge.Require(drudge.Darwin, drudge.Linux))
}
