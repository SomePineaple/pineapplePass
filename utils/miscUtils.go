package utils

import (
	"os"
	"pineapplePass/manager"
)

func ExitProgram(exitCode int) {
	if manager.Current.GetDatabasePath() != "" {
		manager.Current.SaveDatabase()
	}

	os.Exit(exitCode)
}
