package main

import (
	"person/internal/handlers"
	logMsg "person/internal/log"
)

// This application
func main() {
	logMsg.WriteInfoLog("starting application.")
	handlers.HandleRequests()
	logMsg.WriteInfoLog("ending application.")
}
