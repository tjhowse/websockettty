# WebsocketTTY

This is a TTY implementation that satisfies tcell.Tty, such that a websocket can be used
as a transport for a terminal in a browser, using an xterm.js console linked to a websocket
connection via AttachAddon.

The goal is to provide a browser interface to a tview application running serverside.

However I'm having issues stemming from my lack of understanding of how to translate between
TTY-world and glyph-on-screen world.
