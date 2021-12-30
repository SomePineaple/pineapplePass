package manager

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
