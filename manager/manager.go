package manager

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"pineapplePass/utils"
)

var Current Database

func New() Database {
	return Database{}
}

func (d Database) OpenDatabase(databasePath string, masterPassword string) {
	d.databasePath = databasePath
	d.masterPassword = masterPassword
}

func (d Database) CreateDatabase(databasePath string, masterPassword string) {
	d.masterPassword = masterPassword
	d.databasePath = databasePath
	d.masterPasswordSalt = utils.GeneratePasswordSalt()
}

func (d Database) SaveDatabase() {
	aesKey := utils.GenerateAesKey(d.masterPassword, d.masterPasswordSalt)
	bytesToEncrypt, err := json.Marshal(d.masterFolder)
	if err != nil {
		log.Fatalln("(SaveDatabase): Failed to marshal master folder:", err)
	}

	encryptedBytes := utils.AesEncryptBytes(bytesToEncrypt, aesKey)

	dbFile := DatabaseFile{
		masterPasswordSalt:    base64.StdEncoding.EncodeToString(d.masterPasswordSalt),
		encryptedMasterFolder: base64.StdEncoding.EncodeToString(encryptedBytes),
	}

	log.Println(json.Marshal(dbFile))
}

func (d Database) GetDatabasePath() string {
	return d.databasePath
}
