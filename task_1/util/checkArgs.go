package util

import (
	"maps"
	"fmt"
)

// InspectArgs checks that args satisfy conditions to serve as keys in jsonArray entries
func inspectArgs(args []string, jsonArray []map[string]interface{}) error {
	availableKeys := make(map[string]bool)
	numberOfArgvs := len(args)


	if numberOfArgvs < 1 {
		return fmt.Errorf("need at least one key")
	}

	for _, entry := range jsonArray {
		for k := range maps.Keys(entry) {
			availableKeys[k] = true
		}
		if len(args) >= len(entry) {
			return fmt.Errorf("more keys provided than available keys in json input")
		}
	}
	for _, key := range args {
		if !availableKeys[key] {
			return fmt.Errorf("%s is a duplicate key or it's not available", key)
		}
		availableKeys[key] = false
	}

	return nil
}