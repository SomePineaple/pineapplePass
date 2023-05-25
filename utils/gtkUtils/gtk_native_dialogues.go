package gtkUtils

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func OpenFileDialogue(title string) string {

	fileChooser, err := gtk.FileChooserNativeDialogNew(title, nil, gtk.FILE_CHOOSER_ACTION_OPEN, "Open", "Cancel")
	if err != nil {
		log.Println("(OpenFileDialogue): Unable to open file chooser dialogue")
		return ""
	}

	fileChooser.Run()
	filename := fileChooser.GetFilename()

	fileChooser.Destroy()

	return filename
}

func SaveFileDialogue(title string) string {
	fileChooser, err := gtk.FileChooserNativeDialogNew(title, nil, gtk.FILE_CHOOSER_ACTION_SAVE, "Export", "Cancel")
	if err != nil {
		log.Println("(SaveFileDialogue): Unable to open file chooser dialogue")
		return ""
	}
	fileChooser.Run()
	filename := fileChooser.GetFilename()

	fileChooser.Destroy()

	return filename
}
