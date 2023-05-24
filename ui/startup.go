package ui

import (
	"errors"
	"github.com/gotk3/gotk3/gtk"
	"github.com/somepineaple/pineapplePass/manager"
	"github.com/somepineaple/pineapplePass/utils/gtkUtils"
	"log"
	"os"
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

	err = builder.AddFromFile("./pineapple_pass.xml")
	if err != nil {
		log.Fatal("Failed to load ./pineapple_pass.xml:", err)
	}

	gtkUtils.ConnectCheckButton(builder, "ShowPwCheckBox", "toggled", func() {
		entryBox := gtkUtils.GetEntry(builder, "PasswordEntry")
		if entryBox != nil {
			entryBox.SetVisibility(!entryBox.GetVisibility())
		}

		pwConfirmEntry := gtkUtils.GetEntry(builder, "PasswordConfirmEntry")
		if pwConfirmEntry != nil {
			pwConfirmEntry.SetVisibility(!pwConfirmEntry.GetVisibility())
		}
	})

	loginWindow = gtkUtils.GetWindow(builder, "LoginWindow")

	application.AddWindow(loginWindow)

	gtkUtils.ConnectButton(builder, "LoginButton", "clicked", func() {
		if _, err := os.Stat("./defaultSafe.ppass"); errors.Is(err, os.ErrNotExist) {
			showPasswordConfirmDialogue()
		} else {
			passwordEntry := gtkUtils.GetEntry(builder, "PasswordEntry")

			password, err := passwordEntry.GetText()
			if err != nil {
				log.Println("Failed to get text from original PW entry:", err)
			}

			manager.Current = manager.NewDatabase()
			manager.OpenDatabase("./defaultSafe.ppass", password)
			showMainWindow()
		}
	})

	loginWindow.SetTitle("Pineapple Pass")
	loginWindow.SetDefaultSize(400, 300)
	loginWindow.ShowAll()
	loginWindow.Show()
}

func showPasswordConfirmDialogue() {
	pwConfirmDialog = gtkUtils.GetDialog(builder, "PasswordConfirmDialogue")
	pwConfirmDialog.SetTitle("Please Confirm Your Password")

	gtkUtils.ConnectButton(builder, "PasswordConfirmCancel", "clicked", func() {
		pwConfirmDialog.Hide()
	})

	gtkUtils.ConnectButton(builder, "PasswordConfirmOK", "clicked", func() {
		originalPwEntry := gtkUtils.GetEntry(builder, "PasswordEntry")
		pwConfirmEntry := gtkUtils.GetEntry(builder, "PasswordConfirmEntry")

		originalText, err := originalPwEntry.GetText()
		if err != nil {
			log.Println("Failed to get text from original PW entry:", err)
		}

		confirmText, err := pwConfirmEntry.GetText()
		if err != nil {
			log.Println("Failed to get text from pw confirm entry:", err)
		}

		if originalText != confirmText {
			failedLabel := gtkUtils.GetLabel(builder, "PwFailedLabel")
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
