package utils

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func GetTextView(builder *gtk.Builder, objID string) *gtk.TextView {
	obj, err := builder.GetObject(objID)
	if err != nil {
		log.Println("(GetTextView): Failed to get object with id", objID, "err:", err)
		return nil
	}

	textView, ok := obj.(*gtk.TextView)
	if !ok {
		log.Println("(GetTextView): Object with id", objID, "is not of type *gtk.ListBox")
		return nil
	}

	return textView
}

func GetListBox(builder *gtk.Builder, objID string) *gtk.ListBox {
	obj, err := builder.GetObject(objID)
	if err != nil {
		log.Println("(GetListBox): Failed to get object with id", objID, "err:", err)
		return nil
	}

	listBox, ok := obj.(*gtk.ListBox)
	if !ok {
		log.Println("(GetListBox): Object with id", objID, "is not of type *gtk.ListBox")
		return nil
	}

	return listBox
}

func GetWindow(builder *gtk.Builder, objID string) *gtk.Window {
	obj, err := builder.GetObject(objID)
	if err != nil {
		log.Println("(GetWindow): Failed to get object with id", objID, "err:", err)
		return nil
	}

	window, ok := obj.(*gtk.Window)
	if !ok {
		log.Println("(GetWindow): Object with id", objID, "is not of type *gtk.Window")
		return nil
	}

	return window
}

func GetCheckButton(builder *gtk.Builder, objID string) *gtk.CheckButton {
	obj, err := builder.GetObject(objID)
	if err != nil {
		log.Println("(GetCheckButton): Failed to get object with id", objID, "err:", err)
		return nil
	}

	checkButton, ok := obj.(*gtk.CheckButton)
	if !ok {
		log.Println("(GetCheckButton): Object with id", objID, "is not of type *gtk.CheckButton")
		return nil
	}

	return checkButton
}

func GetButton(builder *gtk.Builder, objID string) *gtk.Button {
	obj, err := builder.GetObject(objID)
	if err != nil {
		log.Println("(GetButton): Failed to get object with id", objID, "err:", err)
		return nil
	}

	button, ok := obj.(*gtk.Button)
	if !ok {
		log.Println("(GetButton): Object with id", objID, "is not of type *gtk.Button")
		return nil
	}

	return button
}

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
