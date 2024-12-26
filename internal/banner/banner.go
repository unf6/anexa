package banner

import (
	
	"os"
	"os/exec"
	"runtime"
)

func clear_screen() {
	var clearCmd * exec.Cmd
	if runtime.GOOS == "windows" {
		clearCmd = exec.Command("cmd", "/c", "cls")
	} else {
		clearCmd = exec.Command("clear")
	}
	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}
