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
	withWebsocketTty()
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

func normal() {
	// Build a tview application
	app := tview.NewApplication()
	box := tview.NewBox().SetBorder(true).SetTitle("Hi")
	app.SetRoot(box, true)

	// Start the application running in the backgroun.
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

func withWebsocketTty() {
	// Make a WebsocketTty object
	tty := websockettty.WebsocketTty{}
	// Look up a terminfo definition might should work
	ti, err := tcell.LookupTerminfo("screen-256color")
	if err != nil {
		log.Fatal(err)
	}
	// Build a tcell.Screen object
	screen, err := tcell.NewTerminfoScreenFromTtyTerminfo(&tty, ti)
	if err != nil {
		log.Fatal(err)
	}
	// Build a tview application
	app := tview.NewApplication()
	box := tview.NewBox().SetBorder(true).SetTitle("Hi")
	app.SetScreen(screen)
	app.SetRoot(box, true)

	// Start the application running in the backgroun.
	go app.Run()

	// Crudely give it some time to draw.
	time.Sleep(time.Millisecond * 100)

	// Print the TTY buffer to stdout
	for i := 0; i < len(tty.Screenbuffer); i++ {
		fmt.Printf("%c", tty.Screenbuffer[i])
	}
	fmt.Println("")
}
