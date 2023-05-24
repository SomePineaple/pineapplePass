package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/somepineaple/pineapplePass/manager"
	"github.com/somepineaple/pineapplePass/utils"
	"github.com/somepineaple/pineapplePass/utils/gtkUtils"
	"log"
	"os"
)

var selectedPassword *manager.Password

func showMainWindow() {
	loginWindow.Hide()

	if pwConfirmDialog != nil {
		pwConfirmDialog.Hide()
	}

	window := gtkUtils.GetWindow(builder, "MainWindow")

	foldersListBox := gtkUtils.GetListBox(builder, "FoldersListBox")
	foldersListBox.SetActivateOnSingleClick(true)
	foldersListBox.Connect("row-selected", func() {
		selectedRowIndex := foldersListBox.GetSelectedRow().GetIndex()
		if selectedRowIndex <= 0 {
			manager.Current.CurrentFolder = -1
		} else {
			manager.Current.CurrentFolder = selectedRowIndex - 1
		}

		updatePasswords()
	})

	gtkUtils.GetListBox(builder, "PasswordsListBox").Connect("row-selected", updatePasswordInformationLabel)

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
	gtkUtils.ConnectButton(builder, "NewFolderButton", "clicked", func() {
		newFolderDialog := gtkUtils.GetDialog(builder, "NewFolderDialog")
		gtkUtils.GetEntry(builder, "NewFolderNameEntry").SetText("")
		newFolderDialog.SetTitle("Create New Folder")
		newFolderDialog.ShowAll()
		newFolderDialog.Show()
	})

	setupNewFolderDialog()

	gtkUtils.ConnectButton(builder, "NewPasswordButton", "clicked", func() {
		newPasswordDialog := gtkUtils.GetDialog(builder, "NewPasswordDialog")
		for _, entry := range []string{"NewPasswordNameEntry", "NewPasswordEmailEntry", "NewPasswordPasswordEntry", "NewPasswordNotesEntry"} {
			gtkUtils.GetEntry(builder, entry).SetText("")
		}

		notesTextBuffer, _ := gtkUtils.GetTextView(builder, "NewPasswordNotesTextView").GetBuffer()
		notesTextBuffer.SetText("")

		newPasswordDialog.SetTitle("Create New Password")
		newPasswordDialog.ShowAll()
		newPasswordDialog.Show()
	})

	gtkUtils.ConnectButton(builder, "EditEntryButton", "clicked", setupEditPasswordDialog)

	gtkUtils.ConnectCheckButton(builder, "MainWindowShowPassword", "clicked", updatePasswordInformationLabel)

	gtkUtils.ConnectButton(builder, "CopyEmailButton", "clicked", func() {
		utils.SetClipboardText(selectedPassword.Email)
	})

	gtkUtils.ConnectButton(builder, "CopyPasswordButton", "clicked", func() {
		utils.SetClipboardText(selectedPassword.Password)
	})

	gtkUtils.ConnectMenuItem(builder, "MenuBarExit", "activate", func() {
		manager.SaveDatabase()
		os.Exit(0)
	})

	gtkUtils.ConnectMenuItem(builder, "MenuBarOpen", "activate", func() {
		manager.SaveDatabase()
		openedDatabase := gtkUtils.OpenFileDialogue("Open Database")
		setupOpenSafePasswordDialogue(openedDatabase)
		gtkUtils.GetDialog(builder, "OpenSafePasswordDialogue").Show()

	})

	setupNewPasswordDialog()
}

func updatePasswordInformationLabel() {
	selectedPasswordIdx := gtkUtils.GetListBox(builder, "PasswordsListBox").GetSelectedRow().GetIndex()
	if selectedPasswordIdx == -1 {
		return
	}

	selectedPassword = &manager.GetFolder().ContainedPasswords[selectedPasswordIdx]

	var passwordText string
	var notesText string

	if gtkUtils.GetCheckButton(builder, "MainWindowShowPassword").GetActive() {
		passwordText = selectedPassword.Password
		notesText = selectedPassword.Notes
	} else {
		passwordText = "*******"
		notesText = "*******"
	}

	gtkUtils.GetLabel(builder, "PasswordInformationLabel").SetText(
		"Name: " + selectedPassword.Name + "\n" + "Email: " + selectedPassword.Email + "\n" + "Password: " + passwordText + "\n" + "Notes:\n" + notesText,
	)
}

func updateFolders() {
	foldersListBox := gtkUtils.GetListBox(builder, "FoldersListBox")
	foldersListBox.GetChildren().Foreach(func(item interface{}) {
		foldersListBox.Remove(item.(*gtk.Widget))
	})

	for i, folder := range manager.Current.MasterFolder.ContainedFolders {
		label, err := gtk.LabelNew(folder.Name)
		if err == nil {
			foldersListBox.Insert(label, i)
		} else {
			log.Fatalln("(updateFolders): Failed to create label for folder, name:", folder.Name, "err:", err)
		}
	}

	mainFolder, err := gtk.LabelNew("MasterFolder")
	if err != nil {
		log.Fatalln("(updateFolders): Failed to create label for main folder, err:", err)
	}
	listBoxRow, err := gtk.ListBoxRowNew()
	if err != nil {
		log.Fatalln("(updateFolders): Failed to create listBoxRow for main folder, err:", err)
	}
	listBoxRow.Add(mainFolder)

	foldersListBox.Prepend(listBoxRow)
	foldersListBox.SelectRow(listBoxRow)

	foldersListBox.ShowAll()
}

func updatePasswords() {
	passwordsListBox := gtkUtils.GetListBox(builder, "PasswordsListBox")
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
