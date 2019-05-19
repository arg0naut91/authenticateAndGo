package utils

import (
	"fmt"
	"os"
)

func GetPort() string {

	var port = os.Getenv("PORT")

	if port == "" {
		port = "8080"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port

}

func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	} else {
		return false
	}
}
