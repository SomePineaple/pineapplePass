package manager

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/somepineaple/pineapplePass/utils"
	"log"
	"os"
)

var Current Database

func OpenDatabase(databasePath string, masterPassword string) {
	Current.DatabasePath = databasePath
	Current.masterPassword = masterPassword

	fileBytes, err := os.ReadFile(Current.DatabasePath)
	if err != nil {
		log.Fatalln("(OpenDatabase): Failed to get bytes from databaseFile, err:", err)
	}

	var dbFile DatabaseFile
	err = json.Unmarshal(fileBytes, &dbFile)
	if err != nil {
		log.Fatalln("(OpenDatabase): Failed to unmarshal database, database likely corrupted. err:", err, "fileBytes:", fileBytes)
	}

	Current.masterPasswordSalt, err = base64.StdEncoding.DecodeString(dbFile.MasterPasswordSalt)
	if err != nil {
		log.Fatalln("(OpenDatabase): Failed to get masterPasswordSalt from base64 database, database likely corrupted. err:", err)
	}

	aesKey := utils.GenerateAesKey(Current.masterPassword, Current.masterPasswordSalt)

	encryptedMasterFolder, err := base64.StdEncoding.DecodeString(dbFile.EncryptedMasterFolder)
	if err != nil {
		log.Fatalln("(OpenDatabase): Failed to get encryptedMasterFolder from base64 database, database likely corrupted. err:", err)
	}

	decryptedMasterFolder := utils.AesDecryptBytes(encryptedMasterFolder, aesKey)

	err = json.Unmarshal(decryptedMasterFolder, &Current.MasterFolder)
	if err != nil {
		log.Fatalln("(OpenDatabase): Failed to decrypt master folder, likely due to an incorrect password, err:", err)
	}
}

func CreateDatabase(databasePath string, masterPassword string) {
	Current.masterPassword = masterPassword
	Current.DatabasePath = databasePath
	Current.masterPasswordSalt = append(utils.GeneratePasswordSalt())
}

func SaveDatabase() {
	if Current.DatabasePath == "" {
		return
	}

	ExportDatabase(Current.DatabasePath)
}

func ExportDatabase(savePath string) {
	aesKey := utils.GenerateAesKey(Current.masterPassword, Current.masterPasswordSalt)
	bytesToEncrypt, err := json.Marshal(Current.MasterFolder)
	if err != nil {
		log.Fatalln("(ExportDatabase): Failed to marshal master folder:", err)
	}

	encryptedBytes := utils.AesEncryptBytes(bytesToEncrypt, aesKey)

	dbFile := DatabaseFile{
		MasterPasswordSalt:    base64.StdEncoding.EncodeToString(Current.masterPasswordSalt),
		EncryptedMasterFolder: base64.StdEncoding.EncodeToString(encryptedBytes),
	}

	dbFileBytes, err := json.Marshal(dbFile)
	if err != nil {
		log.Fatalln("(ExportDatabase): Failed to marshal dbFile, err:", err)
	}

	var outputFile *os.File
	if _, err = os.Stat(savePath); errors.Is(err, os.ErrNotExist) {
		outputFile, err = os.Create(savePath)
		if err != nil {
			log.Fatalln("(ExportDatabase): Failed to create output file:", err)
		}
	} else {
		outputFile, err = os.OpenFile(savePath, os.O_WRONLY, os.ModePerm)
		if err != nil {
			log.Fatalln("(ExportDatabase): Failed to open output file:", err)
		}
	}

	_, err = outputFile.Write(dbFileBytes)
	if err != nil {
		log.Fatalln("(ExportDatabase): Failed to write database to file:", err)
	}

	err = outputFile.Close()
	if err != nil {
		log.Fatalln("(ExportDatabase): Failed to close database file:", err)
	}
}

func GetFolder() Folder {
	if Current.CurrentFolder == -1 {
		return Folder{
			Name:               "MasterFolder",
			ContainedPasswords: Current.MasterFolder.ContainedPasswords,
		}
	} else {
		return Current.MasterFolder.ContainedFolders[Current.CurrentFolder]
	}
}

func AddFolder(folderName string) {
	Current.MasterFolder.ContainedFolders = append(Current.MasterFolder.ContainedFolders, NewFolder(folderName))
}

func AddPassword(password Password) {
	if Current.CurrentFolder == -1 {
		Current.MasterFolder.ContainedPasswords = append(Current.MasterFolder.ContainedPasswords, password)
	} else {
		Current.MasterFolder.ContainedFolders[Current.CurrentFolder].ContainedPasswords = append(Current.MasterFolder.ContainedFolders[Current.CurrentFolder].ContainedPasswords, password)
	}
}
