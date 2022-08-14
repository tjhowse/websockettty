package websockettty

// This is going to be a knockoff of tcell.tScreen, but with accessors that let us see into the
// cellbuffer so we can send them out over websocket.

// We might be able to get away with using NewTerminfoScreenFromTtyTerminfo and passing it a custom Tty.

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

// Read reads up to len(p) bytes into p. It returns the number of bytes
// read (0 <= n <= len(p)) and any error encountered. Even if Read
// returns n < len(p), it may use all of p as scratch space during the call.
// If some data is available but not len(p) bytes, Read conventionally
// returns what is available instead of waiting for more.
//
// When Read encounters an error or end-of-file condition after
// successfully reading n > 0 bytes, it returns the number of
// bytes read. It may return the (non-nil) error from the same call
// or return the error (and n == 0) from a subsequent call.
// An instance of this general case is that a Reader returning
// a non-zero number of bytes at the end of the input stream may
// return either err == EOF or err == nil. The next Read should
// return 0, EOF.
//
// Callers should always process the n > 0 bytes returned before
// considering the error err. Doing so correctly handles I/O errors
// that happen after reading some bytes and also both of the
// allowed EOF behaviors.
//
// Implementations of Read are discouraged from returning a
// zero byte count with a nil error, except when len(p) == 0.
// Callers should treat a return of 0 and nil as indicating that
// nothing happened; in particular it does not indicate EOF.
func (w *WebsocketTty) Read(p []byte) (n int, err error) {
	// This should receive keypress data from the websocket.
	// panic("not implemented") // TODO: Implement
	return 0, nil
}

// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
// Write must return a non-nil error if it returns n < len(p).
// Write must not modify the slice data, even temporarily.
func (w *WebsocketTty) Write(p []byte) (n int, err error) {
	// This should write screen character shit to the websocket.
	// if w.writeCB != nil {
	// 	return w.writeCB(p)
	// }
	for i := 0; i < len(w.Screenbuffer); i++ {
		if i < len(p) {
			w.Screenbuffer[i] = p[i]
		}
	}
	return 0, nil
}

func (w *WebsocketTty) Close() error {
	return nil
}
