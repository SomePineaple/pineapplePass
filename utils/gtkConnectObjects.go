package utils

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

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
