package solver

import (
	"maps"
)

func solve(d entryFormat) ([]int, map[int]Room, bool) {
	roomEntered := map[int]bool{}
	rooms := map[int]Room{}
	objectFoundCount := len(d.Objects)
	startingId := d.StartingMapId
	roomPath := []int{}

	// store rooms in a map for faster access
	for _, room := range d.Rooms {
		rooms[room.Id] = room
	}

	currentId := startingId
	for objectFoundCount > 0 {
		roomEntered[currentId] = true
		roomPath = append(roomPath, currentId)

		if checkObjects(rooms[currentId].Objects, d.Objects) {
			objectFoundCount--
		}

		if objectFoundCount == 0 {
			break
		}

		nextRoomId, err := getNextRoomID(roomEntered, rooms[currentId])
		if err != nil {
			nextRoomId = getRandomNextRoom(roomEntered, maps.Keys(rooms))
		}
		currentId = nextRoomId
	}
	return roomPath, rooms, objectFoundCount == 0
}
