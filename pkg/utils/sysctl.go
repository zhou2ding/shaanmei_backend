package utils

import (
	"github.com/eiannone/keyboard"
	"log"
	"os"
)

func PauseExit() {
	log.Println("按任意键退出")
	_, _, _ = keyboard.GetSingleKey()
	os.Exit(0)
}
