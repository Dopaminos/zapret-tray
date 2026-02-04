package main

import (
	_ "embed"
	"os"
	"os/exec"
	"strings"

	"github.com/energye/systray"
)

//go:embed icons/on.png
var iconOn []byte

//go:embed icons/off.png
var iconOff []byte

var enabled bool
var mToggle *systray.MenuItem

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(iconOff)
	systray.SetTooltip("zapret")

	systray.SetOnRClick(func(menu systray.IMenu) {
		menu.ShowMenu()
	})

	systray.SetOnClick(func(menu systray.IMenu) {
		toggle()
		updateState()
		updateUI()
	})

	mToggle = systray.AddMenuItemCheckbox("Включить", "", false)
	mToggle.Click(func() {
		toggle()
		updateState()
		updateUI()
	})

	systray.AddMenuItem("Открыть меню (Терминал)", "").Click(func() {
		openZapretMenu()
	})

	systray.AddMenuItem("Выход", "").Click(func() {
		systray.Quit()
	})

	updateState()
	updateUI()
}

func onExit() {}

func toggle() {
	if enabled {
		exec.Command("systemctl", "stop", "zapret").Run()
	} else {
		exec.Command("systemctl", "start", "zapret").Run()
	}
}

func updateState() {
	out, _ := exec.Command("systemctl", "is-active", "zapret").Output()
	enabled = strings.TrimSpace(string(out)) == "active"
}

func updateUI() {
	if enabled {
		systray.SetIcon(iconOn)
		systray.SetTooltip("Включено")
		mToggle.SetTitle("Выключить")
		mToggle.Check()
	} else {
		systray.SetIcon(iconOff)
		systray.SetTooltip("Выключено")
		mToggle.SetTitle("Включить")
		mToggle.Uncheck()
	}
}

func openZapretMenu() {
	term := os.Getenv("TERMINAL")
	candidates := []string{
		term,
		"alacritty", "kitty", "foot", "wezterm", "konsole", "gnome-terminal",
		"xfce4-terminal", "urxvt", "terminator", "tilix", "xterm",
	}

	for _, t := range candidates {
		if t == "" {
			continue
		}
		cmd := exec.Command(t, "-e", "zapret")
		if cmd.Run() == nil {
			return
		}
	}

	exec.Command("xterm", "-e", "zapret").Run() // fallback
}
