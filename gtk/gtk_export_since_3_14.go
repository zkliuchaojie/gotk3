// Same copyright and license as the rest of the files in this project
//go:build !gtk_3_6 && !gtk_3_8 && !gtk_3_10 && !gtk_3_12
// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12

package gtk

// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/zkliuchaojie/gotk3/glib"
	"github.com/zkliuchaojie/gotk3/internal/callback"
)

//export goListBoxForEachFuncs
func goListBoxForEachFuncs(box *C.GtkListBox, row *C.GtkListBoxRow, userData C.gpointer) {
	fn := callback.Get(uintptr(userData)).(ListBoxForeachFunc)
	fn(wrapListBox(glib.Take(unsafe.Pointer(box))), wrapListBoxRow(glib.Take(unsafe.Pointer(row))))
}
