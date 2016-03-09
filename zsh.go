package main

func init() {
	w.Work("install", "zsh", func() error {
		if err := w.Run("sh", "-c", "curl -fsLS https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh | sh"); err != nil {
			return err
		}

		return w.Stow(".zshrc")
	})

	w.Work("uninstall", "zsh", func() error {
		if err := w.Run("zsh", "-ic", "uninstall_oh_my_zsh"); err != nil {
			return err
		}

		return w.Pack(".zshrc")
	})

	w.Work("upgrade", "zsh", func() error {
		return w.Run("zsh", "-ic", "upgrade_oh_my_zsh")
	})
}
