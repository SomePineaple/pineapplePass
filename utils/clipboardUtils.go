package utils

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

var clipboard *gtk.Clipboard

func SetClipboardText(text string) {
	checkClipboardNil()
	clipboard.SetText(text)
}

func ClearClipboard() {
	checkClipboardNil()
	clipboard.SetText("")
}

func checkClipboardNil() {
	if clipboard == nil {
		var err error
		clipboard, err = gtk.ClipboardGet(gdk.SELECTION_CLIPBOARD)

		if err != nil {
			log.Fatalln("(checkClipboardNil): Failed to get clipboard object, err:", err)
		}
	}
}
