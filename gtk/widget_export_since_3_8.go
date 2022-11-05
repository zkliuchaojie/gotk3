//go:build !gtk_3_6
// +build !gtk_3_6

package gtk

// #include <gdk/gdk.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/zkliuchaojie/gotk3/gdk"
	"github.com/zkliuchaojie/gotk3/glib"
	"github.com/zkliuchaojie/gotk3/internal/callback"
)

//export goTickCallbacks
func goTickCallbacks(widget *C.GtkWidget, frameClock *C.GdkFrameClock, userData C.gpointer) C.gboolean {
	fn := callback.Get(uintptr(userData)).(TickCallback)
	return gbool(fn(
		wrapWidget(glib.Take(unsafe.Pointer(widget))),
		gdk.WrapFrameClock(unsafe.Pointer(frameClock)),
	))
}
