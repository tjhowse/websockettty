package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/tjhowse/websockettty/websockettty"
)

func main() {
	// This doesn't work
	withWebsocketTty(false)

	// This does
	// withWebsocketTty(true)

	// normal()
}

// This produces:

// ╔═══Hi═══╗
// ║        ║
// ║        ║
// ║        ║
// ║        ║
// ║        ║
// ║        ║
// ║        ║
// ║        ║
// ╚════════╝

// (I have clipped this to match the 10x10 terminal size hardcoded in the websockettty package)

func normal() {
	// Build a tview application
	app := tview.NewApplication()
	box := tview.NewBox().SetBorder(true).SetTitle("Hi")
	app.SetRoot(box, true)
	app.Run()
}

// This produces:

// ??????????
// ??????????
// ??????????
// ??????????
// ??????????
// ??????????
// ??????????
// ??????????
// ??????????
// ??????????

func withWebsocketTty(initScreen bool) {
	// Make a WebsocketTty object
	tty := websockettty.WebsocketTty{}
	// Look up a terminfo definition should work
	ti, err := tcell.LookupTerminfo("tmux")
	if err != nil {
		log.Fatal(err)
	}
	// Build a tcell.Screen object
	screen, err := tcell.NewTerminfoScreenFromTtyTerminfo(&tty, ti)
	if err != nil {
		log.Fatal(err)
	}
	if initScreen {
		screen.Init()
	}
	// if !screen.CanDisplay('=', false) {
	// 	log.Fatal("Your terminal cannot display the equals sign")
	// }
	// Build a tview application
	app := tview.NewApplication()
	box := tview.NewBox().SetBorder(true).SetTitle("Hi")
	app.SetRoot(box, true)

	// Start the application running in the background.
	app.SetScreen(screen)
	go app.Run()
	// Crudely give it some time to draw.
	time.Sleep(time.Millisecond * 100)

	// fmt.Println(tty.Screenbuffer)
	fmt.Printf("%s", tty.Screenbuffer)
}
