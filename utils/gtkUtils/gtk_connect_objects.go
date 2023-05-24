package gtkUtils

import (
	"github.com/gotk3/gotk3/gtk"
)

func ConnectButton(builder *gtk.Builder, objID string, detailedSignal string, f interface{}) {
	button := GetButton(builder, objID)
	if button != nil {
		button.Connect(detailedSignal, f)
	}
}

func ConnectCheckButton(builder *gtk.Builder, objID string, detailedSignal string, f interface{}) {
	checkButton := GetCheckButton(builder, objID)
	if checkButton != nil {
		checkButton.Connect(detailedSignal, f)
	}
}

func ConnectMenuItem(builder *gtk.Builder, objID string, detailedSignal string, f interface{}) {
	menuItem := GetMenuItem(builder, objID)
	if menuItem != nil {
		menuItem.Connect(detailedSignal, f)
	}
}
