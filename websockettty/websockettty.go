package websockettty

// Boilerplate generated with
// impl 'w *WebsocketTty' tcell.Tty
type WebsocketTty struct {
	writeCB      func([]byte) (n int, err error)
	Screenbuffer [1024]byte
}

func (w *WebsocketTty) RegisterWriteCB(cb func([]byte) (n int, err error)) {
	w.writeCB = cb
}

// Start is used to activate the Tty for use.  Upon return the terminal should be
// in raw mode, non-blocking, etc.  The implementation should take care of saving
// any state that is required so that it may be restored when Stop is called.
func (w *WebsocketTty) Start() error {
	return nil
}

// Stop is used to stop using this Tty instance.  This may be a suspend, so that other
// terminal based applications can run in the foreground.  Implementations should
// restore any state collected at Start(), and return to ordinary blocking mode, etc.
// Drain is called first to drain the input.  Once this is called, no more Read
// or Write calls will be made until Start is called again.
func (w *WebsocketTty) Stop() error {
	return nil
}

// Drain is called before Stop, and ensures that the reader will wake up appropriately
// if it was blocked.  This workaround is required for /dev/tty on certain UNIX systems
// to ensure that Read() does not block forever.  This typically arranges for the tty driver
// to send data immediately (e.g. VMIN and VTIME both set zero) and sets a deadline on input.
// Implementations may reasonably make this a no-op.  There will still be control sequences
// emitted between the time this is called, and when Stop is called.
func (w *WebsocketTty) Drain() error {
	return nil
}

// NotifyResize is used register a callback when the tty thinks the dimensions have
// changed.  The standard UNIX implementation links this to a handler for SIGWINCH.
// If the supplied callback is nil, then any handler should be unregistered.
func (w *WebsocketTty) NotifyResize(cb func()) {
}

// WindowSize is called to determine the terminal dimensions.  This might be determined
// by an ioctl or other means.
func (w *WebsocketTty) WindowSize() (width int, height int, err error) {
	return 10, 10, nil
}

func (w *WebsocketTty) Read(p []byte) (n int, err error) {
	// This should receive keypress data from the websocket.
	panic("not implemented") // TODO: Implement
	// return 0, nil
}

func (w *WebsocketTty) Write(p []byte) (n int, err error) {
	for i := 0; i < len(w.Screenbuffer); i++ {
		if i < len(p) {
			w.Screenbuffer[i] = p[i]
		} else {
			w.Screenbuffer[i] = 0
		}
	}
	return len(p), nil
}

func (w *WebsocketTty) Close() error {
	return nil
}
