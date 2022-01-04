package ui

import (
	"log"
	"pineapplePass/manager"
	"pineapplePass/utils"
)

func setupNewPasswordDialog() {
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
		manager.SaveDatabase()

		updatePasswords()

		utils.GetDialog(builder, "NewPasswordDialog").Hide()
	})

	utils.ConnectCheckButton(builder, "NewPasswordShowPasswordCheckButton", "toggled", func() {
		label := utils.GetEntry(builder, "NewPasswordPasswordEntry")
		label.SetVisibility(!label.GetVisibility())
	})
}

func setupNewFolderDialog() {
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
}
