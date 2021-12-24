package manager

import "log"

type Manager struct {
	databasePath   string
	masterPassword string
	folders        []Folder
	passwords      []Password
}

type Folder struct {
	containedPasswords []Password
}

type Password struct {
	name     string
	email    string
	password string
	notes    string
}

func New() Manager {
	return Manager{}
}

func (m Manager) OpenDatabase(databasePath string, masterPassword string) {
	m.databasePath = databasePath
	m.masterPassword = masterPassword
}

func (m Manager) CreateDatabase(databasePath string, masterPassword string) {
	m.masterPassword = masterPassword
	m.databasePath = databasePath
	log.Println("Creating database at:", databasePath, "with password:", masterPassword)
}
