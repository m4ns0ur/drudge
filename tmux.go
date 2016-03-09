package main

func init() {
	w.Work("install", "tmux", func() error {
		return w.Stow(".tmux.conf")
	})

	w.Work("uninstall", "tmux", func() error {
		return w.Pack(".tmux.conf")
	})
}
