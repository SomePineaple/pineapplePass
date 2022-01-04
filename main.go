package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
	"pineapplePass/manager"
	"pineapplePass/ui"
)

func main() {
	// Create Gtk Application, change appID to your application domain name reversed.
	const appID = "dev.somepineaple.pineapplepass"
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	// Check to make sure no errors when creating Gtk Application
	if err != nil {
		log.Fatal("Could not create application.", err)
	}
	// Sets the function to run when the application starts
	application.Connect("activate", func() { ui.OnActivate(application) })
	// Run Gtk application
	exitCode := application.Run(os.Args)
	manager.SaveDatabase()
	os.Exit(exitCode)
}
