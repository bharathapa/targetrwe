package main

import (
	"person/internal/handlers"
	logMsg "person/internal/log"
)

// This api basically performs CRUD operation
func main() {
	logMsg.InfoLog("starting application.")
	handlers.HandleRequests()
	logMsg.InfoLog("ending application.")
}
