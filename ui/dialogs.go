package ui

import (
	"log"
	"pineapplePass/manager"
	"pineapplePass/utils/gtkUtils"
)

func setupNewPasswordDialog() {
	gtkUtils.ConnectButton(builder, "NewPasswordCancel", "clicked", func() {
		gtkUtils.GetDialog(builder, "NewPasswordDialog").Hide()
	})

	gtkUtils.ConnectButton(builder, "NewPasswordOK", "clicked", func() {
		name, _ := gtkUtils.GetEntry(builder, "NewPasswordNameEntry").GetText()
		email, _ := gtkUtils.GetEntry(builder, "NewPasswordEmailEntry").GetText()
		password, _ := gtkUtils.GetEntry(builder, "NewPasswordPasswordEntry").GetText()
		notesTextBuffer, _ := gtkUtils.GetTextView(builder, "NewPasswordNotesTextView").GetBuffer()

		start, end := notesTextBuffer.GetBounds()

		notesText, _ := notesTextBuffer.GetText(start, end, false)

		newPassword := manager.Password{
			Name:     name,
			Email:    email,
			Password: password,
			Notes:    notesText,
		}

		manager.AddPassword(newPassword)
		manager.SaveDatabase()

		updatePasswords()

		gtkUtils.GetDialog(builder, "NewPasswordDialog").Hide()
	})

	gtkUtils.ConnectCheckButton(builder, "NewPasswordShowPasswordCheckButton", "toggled", func() {
		label := gtkUtils.GetEntry(builder, "NewPasswordPasswordEntry")
		label.SetVisibility(!label.GetVisibility())
	})
}

func setupNewFolderDialog() {
	gtkUtils.ConnectButton(builder, "NewFolderCancel", "clicked", func() {
		newFolderDialog := gtkUtils.GetDialog(builder, "NewFolderDialog")
		newFolderDialog.Hide()
	})

	gtkUtils.ConnectButton(builder, "NewFolderOK", "clicked", func() {
		gtkUtils.GetDialog(builder, "NewFolderDialog").Hide()
		newFolderName, err := gtkUtils.GetEntry(builder, "NewFolderNameEntry").GetText()
		if err != nil {
			log.Fatalln("Failed to get name for folder from NewFolderNameEntry entryBox, err:", err)
		}

		manager.AddFolder(newFolderName)
		updateFolders()

		manager.SaveDatabase()
	})
}
