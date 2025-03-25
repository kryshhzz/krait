package main

import (
	"os"

	"github.com/fatih/color"

	"github.com/kryshhzz/krait/core"
	"github.com/kryshhzz/krait/utils"
)

func main() {
	color.New(color.BgHiRed, color.FgHiWhite, color.Bold).Println("\n KRAIT \n")

	utils.SRC = os.Args[2]
	utils.DEST = os.Args[3]
	utils.DATE = os.Args[1]

	if len(os.Args) > 4 {
		for _, pref_train_ID := range os.Args {
			utils.PREFFERED_TRAINS[pref_train_ID] = true
		}
	}  
	core.Krait()
}
