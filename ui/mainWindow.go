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

	if pwConfirmDialog != nil {
		pwConfirmDialog.Hide()
	}

	window := utils.GetWindow(builder, "MainWindow")

	foldersListBox := utils.GetListBox(builder, "FoldersListBox")
	foldersListBox.SetActivateOnSingleClick(true)
	foldersListBox.Connect("row-selected", func() {
		// TODO: Make selecting a folder a reality
	})

	updateFolders()
	updatePasswords()

	setupMainWindowButtons()

	window.SetTitle("Pineapple Pass")
	window.SetDefaultSize(800, 800)
	window.Connect("destroy", func() {
		manager.SaveDatabase()
		os.Exit(0)
	})

	window.ShowAll()
	window.Show()
}

func setupMainWindowButtons() {
	utils.ConnectButton(builder, "NewFolderButton", "clicked", func() {
		newFolderDialog := utils.GetDialog(builder, "NewFolderDialog")
		utils.GetEntry(builder, "NewFolderNameEntry").SetText("")
		newFolderDialog.SetTitle("Create New Folder")
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

		manager.AddFolder(newFolderName)
		updateFolders()

		manager.SaveDatabase()
	})

	utils.ConnectButton(builder, "NewPasswordButton", "clicked", func() {
		newPasswordDialog := utils.GetDialog(builder, "NewPasswordDialog")
		for _, entry := range []string{"NewPasswordNameEntry", "NewPasswordEmailEntry", "NewPasswordPasswordEntry", "NewPasswordNotesEntry"} {
			utils.GetEntry(builder, entry).SetText("")
		}

		newPasswordDialog.SetTitle("Create New Password")
		newPasswordDialog.ShowAll()
		newPasswordDialog.Show()
	})

	utils.ConnectButton(builder, "NewPasswordCancel", "clicked", func() {
		utils.GetDialog(builder, "NewPasswordDialog").Hide()
	})

	utils.ConnectButton(builder, "NewPasswordOK", "clicked", func() {
		name, _ := utils.GetEntry(builder, "NewPasswordNameEntry").GetText()
		email, _ := utils.GetEntry(builder, "NewPasswordEmailEntry").GetText()
		password, _ := utils.GetEntry(builder, "NewPasswordPasswordEntry").GetText()
		notes, _ := utils.GetEntry(builder, "NewPasswordNotesEntry").GetText()

		newPassword := manager.Password{
			Name:     name,
			Email:    email,
			Password: password,
			Notes:    notes,
		}

		manager.AddPassword(newPassword)
		updatePasswords()

		utils.GetDialog(builder, "NewPasswordDialog").Hide()
	})
}

func updateFolders() {
	foldersListBox := utils.GetListBox(builder, "FoldersListBox")
	foldersListBox.GetChildren().Foreach(func(item interface{}) {
		foldersListBox.Remove(item.(*gtk.Widget))
	})

	for i, folder := range manager.Current.MasterFolder.ContainedFolders {
		label, err := gtk.LabelNew(folder.Name)
		if err == nil {
			foldersListBox.Insert(label, i)
		} else {
			log.Fatalln("Failed to create label for folder, name:", folder.Name, "err:", err)
		}
	}

	mainFolder, err := gtk.LabelNew("MasterFolder")
	if err != nil {
		log.Fatalln("Failed to create label for main folder, err:", err)
	}
	foldersListBox.Prepend(mainFolder)

	foldersListBox.ShowAll()
}

func updatePasswords() {
	passwordsListBox := utils.GetListBox(builder, "PasswordsListBox")
	passwordsListBox.GetChildren().Foreach(func(item interface{}) {
		passwordsListBox.Remove(item.(*gtk.Widget))
	})

	var allPasswords []manager.Password

	if manager.Current.CurrentFolder == -1 {
		allPasswords = manager.Current.MasterFolder.ContainedPasswords
	} else {
		allPasswords = manager.Current.MasterFolder.ContainedFolders[manager.Current.CurrentFolder].ContainedPasswords
	}

	for i, password := range allPasswords {
		label, err := gtk.LabelNew(password.Name)
		if err == nil {
			passwordsListBox.Insert(label, i)
		}
	}

	passwordsListBox.ShowAll()
}
