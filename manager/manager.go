package manager

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"pineapplePass/utils"
)

var Current Database

func (d Database) OpenDatabase(databasePath string, masterPassword string) {
	d.DatabasePath = databasePath
	d.masterPassword = masterPassword
}

func (d Database) CreateDatabase(databasePath string, masterPassword string) {
	d.masterPassword = masterPassword
	d.DatabasePath = databasePath
	d.masterPasswordSalt = utils.GeneratePasswordSalt()
}

func (d Database) SaveDatabase() {
	aesKey := utils.GenerateAesKey(d.masterPassword, d.masterPasswordSalt)
	bytesToEncrypt, err := json.Marshal(d.MasterFolder)
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
