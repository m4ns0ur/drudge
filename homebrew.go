package main

import (
	"os"

	"github.com/willfaught/drudge/drudge"
)

func init() {
	w.Work("install", "homebrew", func() error {
		if !w.RunTry("which", "brew") {
			if err := w.Run("sh", "-c", "curl -fsLS https://raw.githubusercontent.com/Homebrew/install/master/install | ruby"); err != nil {
				return err
			}

			if !w.RunTry("grep", "/usr/local/sbin", "/etc/paths") {
				if err := w.Run("sh", "-c", "awk '/\\/usr\\/sbin/ {print \"/usr/local/sbin\"}; {print}' /etc/paths | sudo tee /etc/paths"); err != nil {
					return err
				}
			}
		}

		if err := w.Stow(".Brewfile"); err != nil {
			return err
		}

		if err := w.RunMany([][]string{
			{"brew", "bundle", "--global"},
			{"brew", "cleanup"},
			{"brew", "cask", "cleanup"},
			{"/usr/local/opt/fzf/install", "--no-update-rc"},
		}); err != nil {
			return err
		}

		if !w.RunTry("grep", "/usr/local/bin/bash", "/etc/shells") {
			if err := w.Run("sh", "-c", "echo /usr/local/bin/bash | sudo tee -a /etc/shells"); err != nil {
				return err
			}
		}

		if !w.RunTry("grep", "/usr/local/bin/zsh", "/etc/shells") {
			if err := w.Run("sh", "-c", "echo /usr/local/bin/zsh | sudo tee -a /etc/shells"); err != nil {
				return err
			}
		}

		if os.Getenv("SHELL") != "/usr/local/bin/zsh" {
			if err := w.Run("chsh", "-s", "/usr/local/bin/zsh"); err != nil {
				return err
			}
		}

		return nil
	}, drudge.After("install", "xcode"), drudge.Require(drudge.Darwin))
}

func init() {
	w.Work("uninstall", "homebrew", func() error {
		return w.Run("sh", "-c", "curl -fsLS https://raw.githubusercontent.com/Homebrew/install/master/uninstall | ruby")
	}, drudge.Before("uninstall", "xcode"), drudge.Require(drudge.Darwin))
}

func init() {
	w.Work("upgrade", "homebrew", func() error {
		if err := w.Run("brew", "update"); err != nil {
			return err
		}

		if err := w.Run("brew", "upgrade"); err != nil {
			return err
		}

		if err := w.Run("brew", "cleanup"); err != nil {
			return err
		}

		return w.Run("brew", "doctor")
	}, drudge.After("upgrade", "xcode"), drudge.Require(drudge.Darwin))
}
