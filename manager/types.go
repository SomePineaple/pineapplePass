package manager

type Database struct {
	databasePath       string
	masterPassword     string
	masterPasswordSalt []byte
	masterFolder       MasterFolder
}

type DatabaseFile struct {
	masterPasswordSalt    string
	encryptedMasterFolder string
}

type MasterFolder struct {
	containedFolders   []Folder
	containedPasswords []Password
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
