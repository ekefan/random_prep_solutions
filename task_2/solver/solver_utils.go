package solver

import (
	"errors"
	"fmt"
	"iter"
	"strings"

	"github.com/rodaine/table"
)

func checkObjects(objects []Object, entryObjects []string) bool {
	entryObjectMap := make(map[string]bool)
	for _, obj := range entryObjects {
		entryObjectMap[strings.ToLower(obj)] = true // Normalize case for entry objects
	}
	for _, object := range objects {
		if entryObjectMap[strings.ToLower(object.Name)] { // Normalize case for room objects
			return true // Return early if a match is found
		}
	}
	return false // Return false if no matches were found
}

func getNextRoomID(roomsEntered map[int]bool, room Room) (int, error) {
	if !roomsEntered[room.East] && room.East > 0 {
		return room.East, nil
	}
	if !roomsEntered[room.West] && room.West > 0 {
		return room.West, nil
	}
	if !roomsEntered[room.South] && room.South > 0 {
		return room.South, nil
	}
	if !roomsEntered[room.North] && room.North > 0 {
		return room.North, nil
	}
	return 0, errors.New("all rooms entered")
}

func getRandomNextRoom(roomEntered map[int]bool, ids iter.Seq[int]) int {
	for id := range ids {
		if !roomEntered[id] {
			return id
		}
	}
	return 0
}

func getObjects(obj []Object) string {
	objString := strings.Builder{}
	for i, v := range obj {
		if i > 0 {
			objString.WriteString(", ")
		}
		objString.WriteString(v.Name)

	}
	return objString.String()
}

func printOutput(roomPath []int, rooms map[int]Room, mapFile string, found bool) {
	fmt.Printf("---- CALLING SOLVER FOR MAP FILE: %v -----\n", mapFile)

	fmt.Printf("-------- ALL OBJECT FOUND: %v --------\n", found)
	if len(roomPath) > 0 {
		tbl := table.New("Room Id", "Room", "Objects in the Room")
		for _, id := range roomPath {
			room := rooms[id]
			tbl.AddRow(room.Id, room.Name, getObjects(room.Objects))
		}

		tbl.Print()
		fmt.Println("--------------------------------------------")
	}
	fmt.Printf("-------- SOLVED FOR MAP FILE: %v -----------\n", mapFile)
}
