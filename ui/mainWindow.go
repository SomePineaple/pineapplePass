package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"os"
	"pineapplePass/manager"
	"pineapplePass/utils"
)

func showMainWindow() {
	loginWindow.Hide()
	pwConfirmDialog.Hide()

	window := utils.GetWindow(builder, "MainWindow")

	foldersListBox := utils.GetListBox(builder, "FoldersListBox")

	foldersListBox.SetActivateOnSingleClick(true)

	tmp, _ := gtk.LabelNew("Hello there this is a test")

	foldersListBox.Prepend(tmp)
	foldersListBox.ShowAll()
	foldersListBox.SetHAlign(gtk.ALIGN_CENTER)
	foldersListBox.Show()

	window.SetTitle("Pineapple Pass")
	window.SetDefaultSize(800, 800)
	window.Connect("destroy", func() {
		manager.Current.SaveDatabase()
		os.Exit(0)
	})

	window.ShowAll()
	window.Show()
}
