package manager

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"os"
	"pineapplePass/utils"
)

var Current Database

func (d Database) OpenDatabase(databasePath string, masterPassword string) {
	d.DatabasePath = databasePath
	d.masterPassword = masterPassword
}

func CreateDatabase(databasePath string, masterPassword string) {
	Current.masterPassword = masterPassword
	Current.DatabasePath = databasePath
	Current.masterPasswordSalt = append(utils.GeneratePasswordSalt())
	log.Println("Master Password Salt:", Current.masterPasswordSalt)
}

func (d Database) SaveDatabase() {
	aesKey := utils.GenerateAesKey(d.masterPassword, d.masterPasswordSalt)
	bytesToEncrypt, err := json.Marshal(d.MasterFolder)
	if err != nil {
		log.Fatalln("(SaveDatabase): Failed to marshal master folder:", err)
	}

	encryptedBytes := utils.AesEncryptBytes(bytesToEncrypt, aesKey)

	dbFile := DatabaseFile{
		MasterPasswordSalt:    base64.StdEncoding.EncodeToString(d.masterPasswordSalt),
		EncryptedMasterFolder: base64.StdEncoding.EncodeToString(encryptedBytes),
	}

	dbFileBytes, err := json.Marshal(dbFile)

	_, err = os.Stdout.Write(dbFileBytes)
	if err != nil {
		return
	}

	log.Println()
}

func (d Database) AddFolder(folderName string) {
	Current.MasterFolder.ContainedFolders = append(Current.MasterFolder.ContainedFolders, NewFolder(folderName))
}
