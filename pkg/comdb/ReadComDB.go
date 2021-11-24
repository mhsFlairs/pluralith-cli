package comdb

import (
	"fmt"
	"os"
	"path"
	"pluralith/pkg/auxiliary"
)

func ReadComDB() (ComDB, error) {
	// Initialize variables
	var eventDB ComDB

	// Generate proper path
	homeDir, _ := os.UserHomeDir()
	pluralithBus := path.Join(homeDir, "Pluralith", "pluralith_bus.json")

	// Read DB file and handle non-existence
	eventDBString, readErr := os.ReadFile(pluralithBus)
	if readErr != nil {
		var newErr error

		eventDB, newErr = InitComDB() // Create empty DB file
		if newErr != nil {
			fmt.Println(newErr.Error())
			return ComDB{}, newErr
		}

		return eventDB, nil
	}

	// Parse DB string and handle parse error
	eventDBObject, parseErr := auxiliary.ParseJson(string(eventDBString))
	if parseErr != nil {
		var newErr error

		eventDB, newErr = InitComDB() // Create empty DB file
		if newErr != nil {
			fmt.Println(newErr.Error())
			return ComDB{}, newErr
		}

		return eventDB, nil
	}

	// Construct ComDB object
	eventDB = ComDB{
		Locked: eventDBObject["Locked"].(bool),
		Events: eventDBObject["Events"].([]interface{}),
		Errors: eventDBObject["Errors"].([]interface{}),
	}

	// Return parsed DB content
	return eventDB, nil
}