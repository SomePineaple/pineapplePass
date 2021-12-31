package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"pineapplePass/utils"
)

func showMainWindow() {
	window := utils.GetWindow(builder, "MainWindow")

	foldersListBox := utils.GetListBox(builder, "FoldersListBox")

	foldersListBox.SetActivateOnSingleClick(true)

	tmp, _ := gtk.LabelNew("Hello there this is a test")

	foldersListBox.Prepend(tmp)
	foldersListBox.ShowAll()

	window.SetTitle("Pineapple Pass")
	window.ShowAll()
	window.Show()
}
