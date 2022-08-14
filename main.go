package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/tjhowse/websockettty/websockettty"
)

func main() {
	withWebsocketTty()
	os.Exit(0)
	// normal()

	// reply := []byte("\x1B[1;3;31mRed background text\x1B[0m")
	// xterm, tmux:
	// hide_cursor := []byte{27, 91, 63, 50, 53, 108} //"\x1b[?25l"
	// exit_acs := []byte{27, 40, 66} // "\x1b(B"
	// attr_off := []byte{27, 91, 109} // "\x1b[m"

	//                 0x1b   [   ?   2   5    l 0x1b  (    B ESC   [    m ESC   ]   8   ;   ; ESC   \ ESC   [   H ESC
	// reply := []byte{  27, 91, 63, 50, 53, 108,  27, 40, 66, 27, 91, 109, 27, 93, 56, 59, 59, 27, 92, 27, 91, 72, 27, 91, 50, 74, 27, 91, 49, 59, 49, 72, 27, 40, 66, 27, 91, 109, 27, 93, 56, 59, 59, 27, 92, 63, 63, 27, 91, 50, 59, 49, 72, 63, 63, 27, 91, 63, 50, 53, 108}

	// This one I modified to put @ instead of ?
	// reply := []byte{27, 91, 63, 50, 53, 108, 27, 40, 66, 27, 91, 109, 27, 93, 56, 59, 59, 27, 92, 27, 91, 72, 27, 91, 50, 74, 27, 91, 49, 59, 49, 72, 27, 40, 66, 27, 91, 109, 27, 93, 56, 59, 59, 27, 92, 64, 64, 27, 91, 50, 59, 49, 72, 64, 64, 27, 91, 63, 50, 53, 108}

	// \x1b[?25l		hide cursor
	// \x1b(B			exit acs
	// \x1b[m			attr off
	// \x1b]8;;\x1b\	exit URL
	// \x1b[H			Home
	// \x1b[2J			Clear, combined with previous
	// \x1b[1;1H		Mod key? Something
	// \x1b(B			exit acs
	// \x1b[m			attr off
	// \x1b]8;;\x1b\	exit URL
	// ??				Actual text!
	// \x1b[2;1H		Move to 2,1
	// ??				Actual text
	// \x1b[?25l		Hide cursor

	// Red background message in ints.
	// reply := []byte{27, 91, 49, 59, 51, 59, 51, 49, 109, 82, 101, 100, 32, 98, 97, 99, 107, 103, 114, 111, 117, 110, 100, 32, 116, 101, 120, 116, 27, 91, 48, 109}
	// reply := []byte{27, 91, 49, 59, 51, 59, 51, 49, 109, 82, 101, 100, 32, 98, 97, 99, 107, 103, 114, 111, 117, 110, 100, 32, 116, 101, 120, 116, 27, 91, 48, 109}
	// for _, r := range reply {
	// 	if r != 27 {
	// 		fmt.Printf("%c", r)
	// 	} else {
	// 		fmt.Printf("\n\\x1b")
	// 	}
	// }

	// fmt.Println(reply)
	// fmt.Printf("%s", reply)
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

func withWebsocketTty() {
	// Make a WebsocketTty object
	tty := websockettty.WebsocketTty{}
	// Look up a terminfo definition might should work
	// I have tried a few others, including xterm-256color.
	// ti, err := tcell.LookupTerminfo("screen-256color")
	ti, _ := tcell.LookupTerminfo("tmux")
	// ti, err := tcell.LookupTerminfo("xterm")
	// ti, err := tcell.LookupTerminfo("xterm-debian")
	// ti, err := tcell.LookupTerminfo("st")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Build a tcell.Screen object
	screen, err := tcell.NewTerminfoScreenFromTtyTerminfo(&tty, ti)
	// screen, err := tcell.NewTerminfoScreenFromTty(nil)
	// screen, err := tcell.NewTerminfoScreenFromTtyTerminfo(&tty, nil)
	if err != nil {
		log.Fatal(err)
	}
	screen.Init()
	// if !screen.CanDisplay('=', false) {
	// 	log.Fatal("Your terminal cannot display the equals sign")
	// }
	// Build a tview application
	app := tview.NewApplication()
	box := tview.NewBox().SetBorder(true).SetTitle("Hi")
	app.SetRoot(box, true)

	// Start the application running in the background.
	app.SetScreen(screen)
	// app.Sync()
	go app.Run()

	// app.Draw()
	// app.ForceDraw()

	// Crudely give it some time to draw.
	time.Sleep(time.Millisecond * 100)

	// fmt.Println(tty.Screenbuffer)
	fmt.Printf("%s", tty.Screenbuffer)

	// Print the TTY buffer to stdout
	// for i := 0; i < len(tty.Screenbuffer); i++ {
	// 	fmt.Printf("%c", tty.Screenbuffer[i])
	// }
	// fmt.Println("")
	// fmt.Println(tty.Screenbuffer[0])
	// fmt.Println("")
}
