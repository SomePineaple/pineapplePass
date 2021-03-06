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

func setupEditPasswordDialog() {
	if selectedPassword == nil {
		return
	}

	gtkUtils.GetEntry(builder, "EditPasswordNameEntry").SetText(selectedPassword.Name)
	gtkUtils.GetEntry(builder, "EditPasswordEmailEntry").SetText(selectedPassword.Email)
	gtkUtils.GetEntry(builder, "EditPasswordPasswordEntry").SetText(selectedPassword.Password)

	notesTextBuffer, err := gtkUtils.GetTextView(builder, "NewPasswordNotesTextView").GetBuffer()
	if err != nil {
		log.Fatalln("(setupEditPasswordDialog): Failed to get notesTextBuffer, err:", err)
	}

	notesTextBuffer.SetText(selectedPassword.Notes)

	editPasswordDialog := gtkUtils.GetDialog(builder, "EditPasswordDialog")

	gtkUtils.ConnectButton(builder, "EditPasswordCancel", "clicked", editPasswordDialog.Hide)
	gtkUtils.ConnectButton(builder, "EditPasswordSave", "clicked", func() {
		selectedPassword.Name, _ = gtkUtils.GetEntry(builder, "EditPasswordNameEntry").GetText()
		selectedPassword.Email, _ = gtkUtils.GetEntry(builder, "EditPasswordEmailEntry").GetText()
		selectedPassword.Password, _ = gtkUtils.GetEntry(builder, "EditPasswordPasswordEntry").GetText()

		notesTextBuffer, _ = gtkUtils.GetTextView(builder, "EditPasswordNotesTextView").GetBuffer()
		start, end := notesTextBuffer.GetBounds()
		notesText, _ := notesTextBuffer.GetText(start, end, false)

		selectedPassword.Notes = notesText
		editPasswordDialog.Hide()

		updatePasswordInformationLabel()
	})

	gtkUtils.ConnectCheckButton(builder, "EditPasswordShowPasswordCheckButton", "clicked", func() {
		gtkUtils.GetEntry(builder, "EditPasswordPasswordEntry").SetVisibility(gtkUtils.GetCheckButton(builder, "EditPasswordShowPasswordCheckButton").GetActive())
	})

	editPasswordDialog.ShowAll()
	editPasswordDialog.Show()
}
