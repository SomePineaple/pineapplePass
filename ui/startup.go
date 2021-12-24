package ui

import (
	"errors"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
	"pineapplePass/manager"
	"pineapplePass/utils"
)

func OnActivate(application *gtk.Application) {
	builder, err := gtk.BuilderNew()
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

	obj, err := builder.GetObject("LoginWindow")
	if err != nil {
		log.Fatal("Failed to load get window from builder:", err)
	}

	window, ok := obj.(*gtk.Window)
	if !ok {
		log.Fatal("Failed to grab window from glade file")
	}

	application.AddWindow(window)

	utils.ConnectButton(builder, "LoginButton", "pressed", func() {
		if _, err := os.Stat("./defaultSafe.ppass"); errors.Is(err, os.ErrNotExist) {
			showPasswordConfirmDialogue(builder)
		}
	})

	window.SetTitle("Pineapple Pass")
	window.SetDefaultSize(800, 300)
	window.ShowAll()
	window.Show()
}

func showPasswordConfirmDialogue(builder *gtk.Builder) {
	pwConfirmDialog := utils.GetDialog(builder, "PasswordConfirmDialogue")
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
			mgr := manager.New()

			mgr.CreateDatabase("./defaultSafe.ppass", confirmText)
		}
	})

	pwConfirmDialog.SetDefaultSize(200, 200)
	pwConfirmDialog.ShowAll()
	pwConfirmDialog.Show()
}
