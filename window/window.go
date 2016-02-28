package window

import (
	"github.com/flimzy/js/performance"
	"github.com/gopherjs/gopherjs/js"
)

type JSWindow struct {
	js.Object
}

func Window() *JSWindow {
	w := js.Global.Get("window")
	return &JSWindow{*w}
}

func (w *JSWindow) Window() *JSWindow {
	return w
}

func (w *JSWindow) Performance() *performance.Performance {
	p := w.Get("performance")
	return &performance.Performance{*p}
}
