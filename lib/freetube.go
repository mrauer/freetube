package lib

import (
	"encoding/gob"
	"fmt"
	"os"
)

const (
	HISTORY_FILE_PATH = "videos/.history.gob"
)

func GetHistory() map[string]bool {
	data := make(map[string]bool)

	dataFile, err := os.Open(HISTORY_FILE_PATH)

	if err != nil {
		return data
	}

	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&data)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataFile.Close()

	return data
}

func StoreHistory(data map[string]bool) {
	dataFile, err := os.Create(HISTORY_FILE_PATH)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataEncoder := gob.NewEncoder(dataFile)
	dataEncoder.Encode(data)

	dataFile.Close()
}
