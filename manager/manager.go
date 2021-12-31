package manager

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"os"
	"pineapplePass/utils"
)

var Current Database

func OpenDatabase(databasePath string, masterPassword string) {
	// TODO: Open databases properly
	Current.DatabasePath = databasePath
	Current.masterPassword = masterPassword
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
	if err != nil {
		log.Fatalln("Failed to marshal dbFile, err:", err)
	}

	var outputFile *os.File
	if _, err := os.Stat(d.DatabasePath); errors.Is(err, os.ErrNotExist) {
		outputFile, err = os.Create(d.DatabasePath)
		if err != nil {
			log.Fatalln("Failed to create output file:", err)
		}
	} else {
		outputFile, err = os.OpenFile(d.DatabasePath, os.O_RDWR, os.ModePerm)
		if err != nil {
			log.Fatalln("Failed to open output file:", err)
		}
	}

	_, err = outputFile.Write(dbFileBytes)
	if err != nil {
		log.Fatalln("Failed to write database to file:", err)
	}

	err = outputFile.Close()
	if err != nil {
		log.Fatalln("Failed to close database file:", err)
	}
}

func (d Database) AddFolder(folderName string) {
	Current.MasterFolder.ContainedFolders = append(Current.MasterFolder.ContainedFolders, NewFolder(folderName))
}
