package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/somepineaple/pineapplePass/manager"
	"github.com/somepineaple/pineapplePass/ui"
	"log"
	"os"
)

func main() {
	const appID = "dev.somepineaple.pineapplepass"
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatalln("Could not create application.", err)
	}

	application.Connect("activate", func() { ui.OnActivate(application) })

	exitCode := application.Run(os.Args)
	manager.SaveDatabase()
	os.Exit(exitCode)
}
