package main

func init() {
	w.Work("install", "ack", func() error {
		return w.Stow(".ackrc")
	})

	w.Work("uninstall", "ack", func() error {
		return w.Pack(".ackrc")
	})
}
