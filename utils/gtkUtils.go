package utils

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func GetLabel(builder *gtk.Builder, objID string) *gtk.Label {
	obj, err := builder.GetObject(objID)
	if err != nil {
		log.Println("(GetLabel): Failed to get object with id", objID, "err:", err)
		return nil
	}

	label, ok := obj.(*gtk.Label)
	if !ok {
		log.Println("(GetLabel): Object with id", objID, "is not a *gtk.Label type")
		return nil
	}

	return label
}

func GetEntry(builder *gtk.Builder, objID string) *gtk.Entry {
	obj, err := builder.GetObject(objID)
	if err != nil {
		log.Println("(GetEntry): Failed to get object with id", objID, "err:", err)
		return nil
	}

	entry, ok := obj.(*gtk.Entry)
	if !ok {
		log.Println("(GetEntry): Object with id", objID, "is not a *gtk.Entry type")
		return nil
	}

	return entry
}

func GetDialog(builder *gtk.Builder, objID string) *gtk.Dialog {
	obj, err := builder.GetObject(objID)
	if err != nil {
		log.Println("(GetDialog): Failed to get object with id", objID, "err:", err)
		return nil
	}

	dialog, ok := obj.(*gtk.Dialog)
	if !ok {
		log.Println("(GetDialog): Object with id", objID, "is not of type *gtk.Dialog")
		return nil
	}

	return dialog
}

func ConnectButton(builder *gtk.Builder, objID string, detailedSignal string, f interface{}) {
	obj, err := builder.GetObject(objID)
	if err != nil {
		log.Println("(connectButton): Failed to get object", objID, "err:", err)
		return
	}

	proper, ok := obj.(*gtk.Button)
	if !ok {
		log.Println("(ConnectButton):", objID, "is not a gtk button.")
		return
	}

	proper.Connect(detailedSignal, f)
}

func ConnectCheckButton(builder *gtk.Builder, objID string, detailedSignal string, f interface{}) {
	obj, err := builder.GetObject(objID)
	if err != nil {
		log.Println("(connectButton): Failed to get object", objID, "err:", err)
		return
	}

	proper, ok := obj.(*gtk.CheckButton)
	if !ok {
		log.Println("(connectButton): Failed to set object type", objID)
		return
	}

	proper.Connect(detailedSignal, f)
}
