package main

func init() {
	w.Work("install", "readline", func() error {
		return w.Stow(".inputrc")
	})

	w.Work("uninstall", "readline", func() error {
		return w.Pack(".inputrc")
	})
}
