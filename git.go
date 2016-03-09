package main

func init() {
	w.Work("install", "git", func() error {
		return w.Stow(".gitconfig", ".gitignore")
	})

	w.Work("uninstall", "git", func() error {
		return w.Pack(".gitconfig", ".gitignore")
	})
}
