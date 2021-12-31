package ui

import (
	"errors"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
	"pineapplePass/manager"
	"pineapplePass/utils"
)

var builder *gtk.Builder
var loginWindow *gtk.Window
var pwConfirmDialog *gtk.Dialog

func OnActivate(application *gtk.Application) {
	var err error
	builder, err = gtk.BuilderNew()
	if err != nil {
		log.Fatal("Failed to create a new builder object:", err)
	}

	err = builder.AddFromFile("./pineapplePass.glade")
	if err != nil {
		log.Fatal("Failed to load ./pineapplePass.glade:", err)
	}

	utils.ConnectCheckButton(builder, "ShowPwCheckBox", "toggled", func() {
		entryBox := utils.GetEntry(builder, "PasswordEntry")
		if entryBox != nil {
			entryBox.SetVisibility(!entryBox.GetVisibility())
		}

		pwConfirmEntry := utils.GetEntry(builder, "PasswordConfirmEntry")
		if pwConfirmEntry != nil {
			pwConfirmEntry.SetVisibility(!pwConfirmEntry.GetVisibility())
		}
	})

	loginWindow = utils.GetWindow(builder, "LoginWindow")

	application.AddWindow(loginWindow)

	utils.ConnectButton(builder, "LoginButton", "pressed", func() {
		if _, err := os.Stat("./defaultSafe.ppass"); errors.Is(err, os.ErrNotExist) {
			showPasswordConfirmDialogue()
		} else {
			passwordEntry := utils.GetEntry(builder, "PasswordEntry")

			password, err := passwordEntry.GetText()
			if err != nil {
				log.Println("Failed to get text from original PW entry:", err)
			}

			manager.Current = manager.NewDatabase()
			manager.OpenDatabase("./defaultSafe.ppass", password)
		}
	})

	loginWindow.SetTitle("Pineapple Pass")
	loginWindow.SetDefaultSize(400, 300)
	loginWindow.ShowAll()
	loginWindow.Show()
}

func showPasswordConfirmDialogue() {
	pwConfirmDialog = utils.GetDialog(builder, "PasswordConfirmDialogue")
	pwConfirmDialog.SetTitle("Please Confirm Your Password")

	utils.ConnectButton(builder, "PasswordConfirmCancel", "clicked", func() {
		pwConfirmDialog.Hide()
	})

	utils.ConnectButton(builder, "PasswordConfirmOK", "clicked", func() {
		originalPwEntry := utils.GetEntry(builder, "PasswordEntry")
		pwConfirmEntry := utils.GetEntry(builder, "PasswordConfirmEntry")

		originalText, err := originalPwEntry.GetText()
		if err != nil {
			log.Println("Failed to get text from original PW entry:", err)
		}

		confirmText, err := pwConfirmEntry.GetText()
		if err != nil {
			log.Println("Failed to get text from pw confirm entry:", err)
		}

		if originalText != confirmText {
			failedLabel := utils.GetLabel(builder, "PwFailedLabel")
			failedLabel.SetText("Passwords do not match")
		} else {
			manager.Current = manager.NewDatabase()
			manager.CreateDatabase("./defaultSafe.ppass", confirmText)

			showMainWindow()
		}
	})

	pwConfirmDialog.SetDefaultSize(200, 200)
	pwConfirmDialog.ShowAll()
	pwConfirmDialog.Show()
}
