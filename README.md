# WebsocketTTY

This is a TTY implementation that satisfies tcell.Tty, such that a websocket can be used
as a transport for a terminal in a browser, using an xterm.js console linked to a websocket
connection via AttachAddon.

The goal is to provide a browser interface to a tview application running serverside.

# The Problem

I'm having issues stemming from my lack of understanding of how to translate between
TTY-world and glyph-on-screen world. I'm getting *something* printed to console when
I scan over the display buffer and print each character, but each character comes out
as a question mark (?):

Expectation:
```
╔═══Hi═══╗
║        ║
║        ║
║        ║
║        ║
║        ║
║        ║
║        ║
║        ║
╚════════╝
```

Reality:
```
??????????
??????????
??????????
??????????
??????????
??????????
??????????
??????????
??????????
??????????
```

I initially hit this issue after sending the screen buffer over websocket to xterm.js
in the browser and I assumed it was an xterm rendering problem. I put together a minimal
example printing to a local console and got the same problem.

# The Solution?

I made it even simpler and printed bytes to console, then decode the terminal escape sequences.
```
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
```
So it looks like tview/tcell is producing the
```
??
??
```
output rather than it being malformed somewhere after the screen buffer.

My best guess is that the tcell.Screen I'm generating with NewTerminfoScreenFromTtyTerminfo
has no supported characters in its list, so everything gets replaced with the fallback "?"
character.

```go
	tty := websockettty.WebsocketTty{}
	ti, _ := tcell.LookupTerminfo("tmux")
	screen, err := tcell.NewTerminfoScreenFromTtyTerminfo(&tty, ti)
	if err != nil {
		log.Fatal(err)
	}
	screen.Init()
	if !screen.CanDisplay('=', false) {
		log.Fatal("Your terminal cannot display the equals sign")
	} else {
		fmt.Println("Can display equals sign")
	}
```

This passes the CanDisplay check though...