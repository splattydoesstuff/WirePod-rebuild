package main

import (
	"github.com/splattydoesstuff/WirePod-rebuild/cross/podapp"
	cross_win "github.com/splattydoesstuff/WirePod-rebuild/cross/win"
)

func main() {
	podapp.StartWirePod(cross_win.NewWindows())
}
