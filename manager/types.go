package manager

type Database struct {
	DatabasePath       string
	masterPassword     string
	masterPasswordSalt []byte
	MasterFolder       MasterFolder
	CurrentFolder      int
}

type DatabaseFile struct {
	MasterPasswordSalt    string
	EncryptedMasterFolder string
}

type MasterFolder struct {
	ContainedFolders   []Folder
	ContainedPasswords []Password
}

type Folder struct {
	Name               string
	ContainedPasswords []Password
}

type Password struct {
	Name     string
	Email    string
	Password string
	Notes    string
}

func NewDatabase() Database {
	return Database{CurrentFolder: -1}
}

func NewFolder(name string) Folder {
	return Folder{Name: name, ContainedPasswords: []Password{}}
}
