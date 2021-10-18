package database

import (
	"database/sql"
	logMsg "person/internal/log"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

var OpenDbConnection = sql.Open
var Db *sql.DB
var err error

type DbInfr interface {
	openDbConnection(driverName, dataSourceName string) (*sql.DB, error)
}
type EmptyDbConstruct struct{}

func (edc EmptyDbConstruct) openDbConnection(driverName, dataSourceName string) (*sql.DB, error) {
	return OpenDbConnection(driverName, dataSourceName)
}

//ConnectToDB - establishing connection with postgres database
func ConnectToDB(dbInfr DbInfr) (*sql.DB, error) {
	logMsg.InitializeLoggers()
	dbString := "host=localhost port=5445 user=postgres password=postgres dbname=targetrwe sslmode=disable"
	Db, err = dbInfr.openDbConnection("postgres", dbString)
	if err != nil {
		logMsg.ErrorLog(err.Error())
		return nil, err
	}
	logMsg.InfoLog("DB connection successful")
	return Db, nil
}
