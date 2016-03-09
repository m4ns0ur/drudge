package main

import "github.com/willfaught/drudge/drudge"

func init() {
	var packages = []string{"bower", "grunt", "gulp"}

	w.Work("install", "node", func() error {
		return w.Run("npm", append([]string{"install", "-g"}, packages...)...)
	}, drudge.After("install", "homebrew"))

	w.Work("uninstall", "node", func() error {
		return w.Run("npm", append([]string{"uninstall", "-g"}, packages...)...)
	}, drudge.Before("uninstall", "homebrew"))

	w.Work("upgrade", "node", func() error {
		return w.Run("npm", append([]string{"update", "-g"}, packages...)...)
	}, drudge.After("upgrade", "homebrew"))
}
