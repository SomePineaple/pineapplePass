package manager

type Database struct {
	DatabasePath       string
	masterPassword     string
	masterPasswordSalt []byte
	MasterFolder       MasterFolder
}

type DatabaseFile struct {
	masterPasswordSalt    string
	encryptedMasterFolder string
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
	return Database{}
}

func NewFolder(name string) Folder {
	return Folder{Name: name, ContainedPasswords: []Password{}}
}
