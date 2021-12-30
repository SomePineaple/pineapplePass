package manager

import "log"

var Current Manager

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

func (m Manager) SaveDatabase() {

}

func (m Manager) GetDatabasePath() string {
	return m.databasePath
}
