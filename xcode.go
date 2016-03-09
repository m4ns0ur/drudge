package main

import (
	"time"

	"github.com/willfaught/drudge/drudge"
)

func init() {
	w.Work("install", "xcode", func() error {
		if w.RunTry("xcode-select", "-p") {
			return nil
		}

		if err := w.Run("xcode-select", "--install"); err != nil {
			return err
		}

		time.Sleep(10 * time.Second)

		for !w.RunTry("xcode-select", "-p") {
			time.Sleep(10 * time.Second)
		}

		return w.Run("sudo", "xcodebuild", "-license", "accept")
	}, drudge.Require(drudge.Darwin))
}
