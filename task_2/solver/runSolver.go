package solver

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// read file directory to get the path to the json maps...
//

const (
	fileExt string = ".json"
)

type entryFormat struct {
	StartingMapId int      `json:"starting_room_id"`
	Objects       []string `json:"objects"`
	Rooms         []Room   `json:"rooms"`
}
type Room struct {
	Name    string   `json:"name"`
	Objects []Object `json:"objects"`
	Id      int      `json:"id"`
	North   int      `json:"north"`
	South   int      `json:"south"`
	East    int      `json:"east"`
	West    int      `json:"west"`
}

type Object struct {
	Name string `json:"name"`
}

// Determining found
func RunSolver(filePath string) error {

	// check if directory exists
	info, err := os.Stat(filePath)
	if err != nil {
		msg := fmt.Sprintf("path to files: %s does not exist", filePath)
		return errHandler(err, msg)
	}
	// ensuring it's a directory
	if !info.IsDir() {
		return fmt.Errorf("path %s is not a directory", filePath)
	}

	f, err := os.Open(filePath)
	if err != nil {
		msg := fmt.Sprintf("can't open file: %v", filePath)
		return errHandler(err, msg)
	}
	defer f.Close()

	files, _ := f.ReadDir(0)
	for _, v := range files {
		if strings.HasSuffix(v.Name(), fileExt) {
			jsonFile, _ := os.Open(fmt.Sprintf("%s/%s", filePath, v.Name()))
			decoder := json.NewDecoder(jsonFile)
			var data entryFormat
			if err := decoder.Decode(&data); err != nil {
				msg := err.Error()
				return errHandler(err, msg)
			}
			jsonFile.Close()
			roomPath, rooms, solved := solve(data)
			printOutput(roomPath, rooms, jsonFile.Name(), solved)
		}
	}
	return nil
}

func errHandler(err error, msg string) error {
	return fmt.Errorf("%s\n %v", msg, err)
}