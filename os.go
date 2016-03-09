package main

import (
	"path/filepath"

	"github.com/willfaught/drudge/drudge"
)

func init() {
	var defs = []defaultsCmd{
		{Name: "com.apple.messageshelper.MessageController", Key: "SOInputLineSettings", Add: true, Dict: map[string]string{"automaticEmojiSubstitutionEnablediMessage": "false"}},
		{Name: "com.apple.messageshelper.MessageController", Key: "SOInputLineSettings", Add: true, Dict: map[string]string{"automaticQuoteSubstitutionEnabled": "false"}},
		{Name: "com.apple.messageshelper.MessageController", Key: "SOInputLineSettings", Add: true, Dict: map[string]string{"continuousSpellCheckingEnabled": "false"}},
		{Name: "com.apple.finder", Key: "FXInfoPanesExpanded", Dict: map[string]string{"General": "true", "OpenWith": "true"}},
		{Key: "_HIHideMenuBar", Kind: "bool", Value: "true", Global: true},
		{Key: "AppleFontSmoothing", Kind: "int", Value: "2", Global: true},
		{Key: "AppleKeyboardUIMode", Kind: "int", Value: "3", Global: true},
		{Key: "AppleLanguages", Array: []string{"en"}, Global: true},
		{Key: "AppleLocale", Kind: "string", Value: "en_US@currency=USD", Global: true},
		{Key: "AppleMeasurementUnits", Kind: "string", Value: "Inches", Global: true},
		{Key: "AppleMetricUnits", Kind: "bool", Value: "false", Global: true},
		{Key: "ApplePressAndHoldEnabled", Kind: "bool", Value: "false", Global: true},
		{Key: "AppleShowAllExtensions", Kind: "bool", Value: "true", Global: true},
		{Key: "com.apple.mouse.tapBehavior", Kind: "int", Value: "1", Current: true, Global: true},
		{Key: "com.apple.mouse.tapBehavior", Kind: "int", Value: "1", Global: true},
		{Key: "com.apple.sound.beep.feedback", Kind: "int", Value: "0"},
		{Key: "com.apple.springing.delay", Kind: "float", Value: "0", Global: true},
		{Key: "com.apple.trackpad.enableSecondaryClick", Kind: "bool", Value: "true", Current: true, Global: true},
		{Key: "KeyRepeat", Kind: "int", Value: "1", Global: true},
		{Key: "NSAutomaticDashSubstitutionEnabled", Kind: "bool", Value: "false", Global: true},
		{Key: "NSAutomaticQuoteSubstitutionEnabled", Kind: "bool", Value: "false", Global: true},
		{Key: "NSAutomaticSpellingCorrectionEnabled", Kind: "bool", Value: "false", Global: true},
		{Key: "NSDocumentSaveNewDocumentsToCloud", Kind: "bool", Value: "false", Global: true},
		{Key: "NSDocumentSaveNewDocumentsToCloud", Kind: "bool", Value: "true", Global: true},
		{Key: "NSNavPanelExpandedStateForSaveMode", Kind: "bool", Value: "true", Global: true},
		{Key: "NSTableViewDefaultSizeMode", Kind: "int", Value: "2", Global: true},
		{Key: "NSTextShowsControlCharacters", Kind: "bool", Value: "true", Global: true},
		{Key: "NSWindowResizeTime", Kind: "float", Value: "0.001", Global: true},
		{Key: "PMPrintingExpandedStateForPrint", Kind: "bool", Value: "true", Global: true},
		{Name: "/.Spotlight-V100/VolumeConfiguration", Key: "Exclusions", Array: []string{"/Volumes"}, Sudo: true},
		{Name: "com.apple.windowserver", Key: "DisplayResolutionEnabled", Kind: "bool", Value: "true"},
		{Name: "com.apple.ActivityMonitor", Key: "IconType", Kind: "int", Value: "5"},
		{Name: "com.apple.ActivityMonitor", Key: "OpenMainWindow", Kind: "bool", Value: "true"},
		{Name: "com.apple.ActivityMonitor", Key: "ShowCategory", Kind: "int", Value: "0"},
		{Name: "com.apple.ActivityMonitor", Key: "SortColumn", Kind: "string", Value: "CPUUsage"},
		{Name: "com.apple.ActivityMonitor", Key: "SortDirection", Kind: "int", Value: "0"},
		{Name: "com.apple.AppleMultitouchTrackpad", Key: "Clicking", Kind: "int", Value: "1"},
		{Name: "com.apple.AppleMultitouchTrackpad", Key: "TrackpadCornerSecondaryClick", Kind: "int", Value: "0"},
		{Name: "com.apple.AppleMultitouchTrackpad", Key: "TrackpadRightClick", Kind: "int", Value: "1"},
		{Name: "com.apple.appstore", Key: "ShowDebugMenu", Kind: "bool", Value: "true"},
		{Name: "com.apple.BluetoothAudioAgent", Key: "Apple Bitpool Min (editable)", Kind: "int", Value: "40"},
		{Name: "com.apple.commerce", Key: "AutoUpdate", Kind: "bool", Value: "true"},
		{Name: "com.apple.dashboard", Key: "mcx-disabled", Kind: "bool", Value: "true"},
		{Name: "com.apple.desktopservices", Key: "DSDontWriteNetworkStores", Kind: "bool", Value: "false"},
		{Name: "com.apple.dock", Key: "autohide-delay", Kind: "float", Value: "0"},
		{Name: "com.apple.dock", Key: "autohide-time-modifier", Kind: "float", Value: "0"},
		{Name: "com.apple.dock", Key: "autohide", Kind: "bool", Value: "true"},
		{Name: "com.apple.dock", Key: "dashboard-in-overlay", Kind: "bool", Value: "false"},
		{Name: "com.apple.dock", Key: "enable-spring-load-actions-on-all-items", Kind: "bool", Value: "true"},
		{Name: "com.apple.dock", Key: "expose-animation-duration", Kind: "int", Value: "0"},
		{Name: "com.apple.dock", Key: "expose-group-by-app", Kind: "bool", Value: "true"},
		{Name: "com.apple.dock", Key: "launchanim", Kind: "bool", Value: "false"},
		{Name: "com.apple.dock", Key: "magnification", Kind: "bool", Value: "false"},
		{Name: "com.apple.dock", Key: "mineffect", Kind: "string", Value: "scale"},
		{Name: "com.apple.dock", Key: "minimize-to-application", Kind: "bool", Value: "true"},
		{Name: "com.apple.dock", Key: "mru-spaces", Kind: "bool", Value: "false"},
		{Name: "com.apple.dock", Key: "persistent-apps", Array: []string{""}},
		{Name: "com.apple.dock", Key: "show-process-indicators", Kind: "bool", Value: "true"},
		{Name: "com.apple.dock", Key: "showhidden", Kind: "bool", Value: "true"},
		{Name: "com.apple.dock", Key: "static-only", Kind: "bool", Value: "true"},
		{Name: "com.apple.dock", Key: "tilesize", Kind: "int", Value: "48"},
		{Name: "com.apple.dock", Key: "wvous-bl-corner", Kind: "int", Value: "5"},
		{Name: "com.apple.dock", Key: "wvous-bl-modifier", Kind: "int", Value: "0"},
		{Name: "com.apple.dock", Key: "wvous-br-corner", Kind: "int", Value: "10"},
		{Name: "com.apple.dock", Key: "wvous-br-modifier", Kind: "int", Value: "0"},
		{Name: "com.apple.dock", Key: "wvous-tl-corner", Kind: "int", Value: "0"},
		{Name: "com.apple.dock", Key: "wvous-tl-modifier", Kind: "int", Value: "0"},
		{Name: "com.apple.dock", Key: "wvous-tr-corner", Kind: "int", Value: "0"},
		{Name: "com.apple.dock", Key: "wvous-tr-modifier", Kind: "int", Value: "0"},
		{Name: "com.apple.driver.AppleBluetoothMultitouch.trackpad", Key: "Clicking", Kind: "bool", Value: "true"},
		{Name: "com.apple.driver.AppleBluetoothMultitouch.trackpad", Key: "TrackpadCornerSecondaryClick", Kind: "int", Value: "0"},
		{Name: "com.apple.driver.AppleBluetoothMultitouch.trackpad", Key: "TrackpadRightClick", Kind: "bool", Value: "true"},
		{Name: "com.apple.finder", Key: "_FXShowPosixPathInTitle", Kind: "bool", Value: "true"},
		{Name: "com.apple.finder", Key: "AppleShowAllFiles", Kind: "bool", Value: "true"},
		{Name: "com.apple.finder", Key: "DisableAllAnimations", Kind: "bool", Value: "true"},
		{Name: "com.apple.finder", Key: "FXDefaultSearchScope", Kind: "string", Value: "SCcf"},
		{Name: "com.apple.finder", Key: "FXEnableExtensionChangeWarning", Kind: "bool", Value: "false"},
		{Name: "com.apple.finder", Key: "FXPreferredViewStyle", Kind: "string", Value: "clmv"},
		{Name: "com.apple.finder", Key: "NewWindowTarget", Kind: "string", Value: "PfDe"},
		{Name: "com.apple.finder", Key: "NewWindowTargetPath", Kind: "string", Value: "file://$HOME/"},
		{Name: "com.apple.finder", Key: "QLEnableTextSelection", Kind: "bool", Value: "true"},
		{Name: "com.apple.finder", Key: "ShowExternalHardDrivesOnDesktop", Kind: "bool", Value: "false"},
		{Name: "com.apple.finder", Key: "ShowHardDrivesOnDesktop", Kind: "bool", Value: "false"},
		{Name: "com.apple.finder", Key: "ShowMountedServersOnDesktop", Kind: "bool", Value: "false"},
		{Name: "com.apple.finder", Key: "ShowPathbar", Kind: "bool", Value: "false"},
		{Name: "com.apple.finder", Key: "ShowRecentTags", Kind: "bool", Value: "false"},
		{Name: "com.apple.finder", Key: "ShowRemovableMediaOnDesktop", Kind: "bool", Value: "false"},
		{Name: "com.apple.finder", Key: "ShowStatusBar", Kind: "bool", Value: "true"},
		{Name: "com.apple.finder", Key: "WarnOnEmptyTrash", Kind: "bool", Value: "false"},
		{Name: "com.apple.ical", Key: "first day of the week", Kind: "int", Value: "1"},
		{Name: "com.apple.ical", Key: "first minute of work hours", Kind: "int", Value: "540"},
		{Name: "com.apple.ImageCapture", Key: "disableHotPlug", Kind: "bool", Value: "true", Current: true},
		{Name: "com.apple.LaunchServices", Key: "LSQuarantine", Kind: "bool", Value: "false"},
		{Name: "com.apple.mail", Key: "AddressesIncludeNameOnPasteboard", Kind: "bool", Value: "false"},
		{Name: "com.apple.mail", Key: "DisableInlineAttachmentViewing", Kind: "bool", Value: "true"},
		{Name: "com.apple.mail", Key: "DisableReplyAnimations", Kind: "bool", Value: "true"},
		{Name: "com.apple.Mail", Key: "DisableReplyAnimations", Kind: "bool", Value: "true"},
		{Name: "com.apple.Mail", Key: "DisableSendAnimations", Kind: "bool", Value: "true"},
		{Name: "com.apple.mail", Key: "DisableSendAnimations", Kind: "bool", Value: "true"},
		{Name: "com.apple.mail", Key: "SpellCheckingBehavior", Kind: "string", Value: "NoSpellCheckingEnabled"},
		{Name: "com.apple.NetworkBrowser", Key: "BrowseAllInterfaces", Kind: "bool", Value: "true"},
		{Name: "com.apple.print.PrintingPrefs", Key: "Quit When Finished", Kind: "bool", Value: "true"},
		{Name: "com.apple.Safari", Key: "AutoOpenSafeDownloads", Kind: "bool", Value: "false"},
		{Name: "com.apple.Safari", Key: "com.apple.Safari.ContentPageGroupIdentifier.WebKit2DeveloperExtrasEnabled", Kind: "bool", Value: "true"},
		{Name: "com.apple.Safari", Key: "com.apple.Safari.ContentPageGroupIdentifier.WebKit2TabsToLinks", Kind: "bool", Value: "true"},
		{Name: "com.apple.Safari", Key: "FindOnPageMatchesWordStartsOnly", Kind: "bool", Value: "false"},
		{Name: "com.apple.Safari", Key: "HomePage", Kind: "string", Value: "about:blank"},
		{Name: "com.apple.Safari", Key: "IncludeDevelopMenu", Kind: "bool", Value: "true"},
		{Name: "com.apple.Safari", Key: "IncludeInternalDebugMenu", Kind: "bool", Value: "true"},
		{Name: "com.apple.Safari", Key: "ShowFavoritesBar", Kind: "bool", Value: "false"},
		{Name: "com.apple.Safari", Key: "ShowFullURLInSmartSearchField", Kind: "bool", Value: "true"},
		{Name: "com.apple.Safari", Key: "WebKitDeveloperExtrasEnabledPreferenceKey", Kind: "bool", Value: "true"},
		{Name: "com.apple.Safari", Key: "WebKitTabToLinksPreferenceKey", Kind: "bool", Value: "true"},
		{Name: "com.apple.screencapture", Key: "disable-shadow", Kind: "bool", Value: "true"},
		{Name: "com.apple.screencapture", Key: "location", Kind: "string", Value: "$HOME/Desktop"},
		{Name: "com.apple.screensaver", Key: "askForPassword", Kind: "int", Value: "1"},
		{Name: "com.apple.screensaver", Key: "askForPasswordDelay", Kind: "int", Value: "0"},
		{Name: "com.apple.SoftwareUpdate", Key: "AutomaticCheckEnabled", Kind: "bool", Value: "true"},
		{Name: "com.apple.SoftwareUpdate", Key: "AutomaticDownload", Kind: "int", Value: "1"},
		{Name: "com.apple.SoftwareUpdate", Key: "CriticalUpdateInstall", Kind: "int", Value: "1"},
		{Name: "com.apple.SoftwareUpdate", Key: "ScheduleFrequency", Kind: "int", Value: "1"},
		{Name: "com.apple.systemsound", Key: "com.apple.sound.beep.volume", Kind: "float", Value: "0"},
		{Name: "com.apple.systemsound", Key: "com.apple.sound.uiaudio.enabled", Kind: "int", Value: "0"},
		{Name: "com.apple.terminal", Key: "Default Window Settings", Kind: "string", Value: "Solarized Dark"},
		{Name: "com.apple.terminal", Key: "SecureKeyboardEntry", Kind: "bool", Value: "true"},
		{Name: "com.apple.Terminal", Key: "ShowLineMarks", Kind: "int", Value: "0"},
		{Name: "com.apple.terminal", Key: "Startup Window Settings", Kind: "string", Value: "Solarized Dark"},
		{Name: "com.apple.terminal", Key: "StringEncodings", Kind: "array", Value: "4"},
		{Name: "com.apple.TextEdit", Key: "PlainTextEncoding", Kind: "int", Value: "4"},
		{Name: "com.apple.TextEdit", Key: "PlainTextEncodingForWrite", Kind: "int", Value: "4"},
		{Name: "com.apple.TextEdit", Key: "RichText", Kind: "int", Value: "0"},
		{Name: "com.apple.TimeMachine", Key: "DoNotOfferNewDisksForBackup", Kind: "bool", Value: "true"},
		{Name: filepath.Join(drudge.Home, "Library/Preferences/.GlobalPreferences.plist"), Key: "AppleInterfaceTheme", Kind: "string", Value: "Dark"},
	}

	var cmds = [][]string{
		{"/System/Library/Frameworks/CoreServices.framework/Frameworks/LaunchServices.framework/Support/lsregister", "-kill", "-r", "-domain", "local", "-domain", "system", "-domain", "user"},
		{"/usr/libexec/PlistBuddy", "-c", "Set :DesktopViewSettings:IconViewSettings:arrangeBy grid", filepath.Join(drudge.Home, "Library/Preferences/com.apple.finder.plist")},
		{"/usr/libexec/PlistBuddy", "-c", "Set :FK_StandardViewSettings:IconViewSettings:arrangeBy grid", filepath.Join(drudge.Home, "Library/Preferences/com.apple.finder.plist")},
		{"/usr/libexec/PlistBuddy", "-c", "Set :StandardViewSettings:IconViewSettings:arrangeBy grid", filepath.Join(drudge.Home, "Library/Preferences/com.apple.finder.plist")},

		{"chflags", "nohidden", filepath.Join(drudge.Home, "Library")},
		{"find", filepath.Join(drudge.Home, "Library/Application Support/Dock"), "-name", "*.db", "-maxdepth", "1", "-delete"},

		{"sudo", "nvram", "SystemAudioVolume=' '"},
		{"sudo", "systemsetup", "-settimezone", "America/Los_Angeles"},
		{"sudo", "systemsetup", "-setcomputersleep", "20"},
		{"sudo", "systemsetup", "-setrestartfreeze", "on"},
	}

	w.Work("install", "os", func() error {
		for _, def := range defs {
			var cmd = def.command()

			if err := w.Run(cmd[0], cmd[1:]...); err != nil {
				return err
			}
		}

		for _, cmd := range cmds {
			if err := w.Run(cmd[0], cmd[1:]...); err != nil {
				return err
			}
		}

		return nil
	}, drudge.Require(drudge.Darwin))
}

type defaultsCmd struct {
	Array                      []string
	Add, Current, Global, Sudo bool
	Dict                       map[string]string
	Key, Kind, Name, Value     string
}

func (d *defaultsCmd) command() []string {
	var c []string

	if d.Sudo {
		c = append(c, "sudo")
	}

	c = append(c, "defaults")

	if d.Current {
		c = append(c, "-currentHost")
	}

	c = append(c, "write")

	if d.Global {
		c = append(c, "-g")
	}

	if d.Name != "" {
		c = append(c, d.Name)
	}

	c = append(c, d.Key)

	if len(d.Array) > 0 {
		if d.Add {
			c = append(c, "-array-add")
		} else {
			c = append(c, "-array")
		}

		c = append(c, d.Array...)
	} else if len(d.Dict) > 0 {
		if d.Add {
			c = append(c, "-dict-add")
		} else {
			c = append(c, "-dict")
		}

		for k, v := range d.Dict {
			c = append(c, k, v)
		}
	} else {
		if d.Kind != "" {
			c = append(c, "-"+d.Kind)
		}

		c = append(c, d.Value)
	}

	return c
}
