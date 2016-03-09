package main

func init() {
	w.Work("install", "terminal", func() error {
		return w.Stow("Library/Preferences/com.apple.Terminal.plist")
	})

	w.Work("uninstall", "terminal", func() error {
		return w.Pack("Library/Preferences/com.apple.Terminal.plist")
	})
}
