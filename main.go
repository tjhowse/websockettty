package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// func main() {
// 	web.ServeWebserver()
// }

func main() {
	tty := websockettty.WebsocketTty{}
	ti, err := tcell.LookupTerminfo("screen-256color")
	if err != nil {
		log.Fatal(err)
	}
	screen, err := tcell.NewTerminfoScreenFromTtyTerminfo(&tty, ti)
	// screen, err := tcell.NewTerminfoScreenFromTtyTerminfo(&tty, nil)
	if err != nil {
		log.Fatal(err)
	}
	app := tview.NewApplication()
	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	app.SetScreen(screen)
	app.SetRoot(box, true)
	go app.Run()
	time.Sleep(time.Millisecond * 100)

	for i := 0; i < len(tty.Screenbuffer); i++ {
		fmt.Printf("%c", tty.Screenbuffer[i])
	}
	// app.Stop()
	fmt.Println("")
}
