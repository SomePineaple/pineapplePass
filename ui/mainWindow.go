package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
	"pineapplePass/manager"
	"pineapplePass/utils"
)

func showMainWindow() {
	loginWindow.Hide()
	pwConfirmDialog.Hide()

	window := utils.GetWindow(builder, "MainWindow")

	// TODO: Sort list box alphabetically
	foldersListBox := utils.GetListBox(builder, "FoldersListBox")
	foldersListBox.SetActivateOnSingleClick(true)

	setupMainWindowButtons()

	window.SetTitle("Pineapple Pass")
	window.SetDefaultSize(800, 800)
	window.Connect("destroy", func() {
		manager.Current.SaveDatabase()
		os.Exit(0)
	})

	window.ShowAll()
	window.Show()
}

func setupMainWindowButtons() {
	utils.ConnectButton(builder, "NewFolderButton", "clicked", func() {
		newFolderDialog := utils.GetDialog(builder, "NewFolderDialog")
		utils.GetEntry(builder, "NewFolderNameEntry").SetText("")
		newFolderDialog.ShowAll()
		newFolderDialog.Show()
	})

	utils.ConnectButton(builder, "NewFolderCancel", "clicked", func() {
		newFolderDialog := utils.GetDialog(builder, "NewFolderDialog")
		newFolderDialog.Hide()
	})

	utils.ConnectButton(builder, "NewFolderOK", "clicked", func() {
		utils.GetDialog(builder, "NewFolderDialog").Hide()
		newFolderName, err := utils.GetEntry(builder, "NewFolderNameEntry").GetText()
		if err != nil {
			log.Fatalln("Failed to get name for folder from NewFolderNameEntry entryBox, err:", err)
		}

		manager.Current.MasterFolder.ContainedFolders = append(manager.Current.MasterFolder.ContainedFolders, manager.NewFolder(newFolderName))
		updateFolders()
	})
}

func updateFolders() {
	foldersListBox := utils.GetListBox(builder, "FoldersListBox")
	foldersListBox.GetChildren().Foreach(func(item interface{}) {
		foldersListBox.Remove(item.(*gtk.Widget))
	})

	for i := 0; i < len(manager.Current.MasterFolder.ContainedFolders); i++ {
		label, err := gtk.LabelNew(manager.Current.MasterFolder.ContainedFolders[i].Name)
		if err == nil {
			foldersListBox.Prepend(label)
		} else {
			log.Fatalln("Failed to create label for folder, name:", manager.Current.MasterFolder.ContainedFolders[i].Name, "err:", err)
		}
	}

	foldersListBox.ShowAll()
}

func updatePasswords() {
	passwordsListBox := utils.GetListBox(builder, "PasswordsListBox")
	passwordsListBox.GetChildren().Free()
}
