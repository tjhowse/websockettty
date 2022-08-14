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