package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/getlantern/systray"
)

//go:embed on.png
var iconOn []byte

//go:embed off.png
var iconOff []byte
var mToggle *systray.MenuItem
var mAuto *systray.MenuItem

const autostartFile = ".config/autostart/zapret-tray.desktop"

func main() {
	home, _ := os.UserHomeDir()
	autoPath := filepath.Join(home, autostartFile)
	execPath, _ := os.Executable()

	systray.Run(func() { onReady(execPath, autoPath) }, nil)
}

func onReady(execPath, autoPath string) {
	systray.SetTitle("zapret")
	systray.SetTooltip("zapret")

	mToggle = systray.AddMenuItemCheckbox("Выключено", "", false)
	mAuto = systray.AddMenuItemCheckbox("Автозагрузка", "", false)
	open := systray.AddMenuItem("Открыть меню snowy-fluffy", "")
	quit := systray.AddMenuItem("Выход", "")

	updateAuto(mAuto, autoPath)
	updateToggle()

	go func() {
		for {
			select {
			case <-mToggle.ClickedCh:
				toggleZapret()
				updateToggle()
			case <-mAuto.ClickedCh:
				toggleAuto(execPath, autoPath, mAuto)
			case <-open.ClickedCh:
				openZapretMenu()
			case <-quit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func toggleZapret() {
	action := "start"
	if isZapretActive() {
		action = "stop"
	}
	exec.Command("sudo", "systemctl", action, "zapret").Run()
}

func isZapretActive() bool {
	out, _ := exec.Command("systemctl", "is-active", "zapret").Output()
	return strings.TrimSpace(string(out)) == "active"
}

func updateToggle() {
	if isZapretActive() {
		mToggle.Check()
		mToggle.SetTitle("Выключить")
	} else {
		mToggle.Uncheck()
		mToggle.SetTitle("Включить")
	}
}

func toggleAuto(execPath, path string, m *systray.MenuItem) {
	if m.Checked() {
		os.Remove(path)
		m.Uncheck()
	} else {
		os.MkdirAll(filepath.Dir(path), 0755)
		content := fmt.Sprintf(`[Desktop Entry]
Type=Application
Exec=%s
Terminal=false
Hidden=false
Name=Zapret Tray
`, execPath)
		os.WriteFile(path, []byte(content), 0644)
		m.Check()
	}
}

func updateAuto(m *systray.MenuItem, path string) {
	if _, err := os.Stat(path); err == nil {
		m.Check()
	}
}

func openZapretMenu() {
	cmd := "zapret"
	if _, err := exec.LookPath("zapret"); err != nil {
		cmd = `echo "zapret не установлен!" && echo "" && echo "Установить: sh -c \"\$(curl -fsSL https://raw.githubusercontent.com/Snowy-Fluffy/zapret.installer/refs/heads/main/installer.sh)\"" && echo "" && echo "Нажмите Enter..." && read`
	}
	term := os.Getenv("TERMINAL")
	if term == "" {
		term = "xterm"
	}
	exec.Command(term, "-e", "sh", "-c", cmd).Start()
}
