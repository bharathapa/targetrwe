package database

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"
)

type FakeEmptyDbConstruct struct {
	dbInstance *sql.DB
	dbError error
}
func (edc FakeEmptyDbConstruct) openDbConnection(driverName, dataSourceName string) (*sql.DB, error){
	return edc.dbInstance, edc.dbError
}
func TestConnectToDB(t *testing.T) {
	type args struct {
		dbInfr DbInfr
	}
	case1:= FakeEmptyDbConstruct{dbError: errors.New(" Unable to connect to db")}
	case2:= FakeEmptyDbConstruct{}
	arguments := []args{{case1},{case2}}
	tests := []struct {
		name    string
		args    args
		want    *sql.DB
		wantErr bool
	}{
		{"Case 1: database connection error ", arguments[0],nil,true},
		{"Case 2: successful connection", arguments[1],nil,false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConnectToDB(tt.args.dbInfr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectToDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectToDB() got = %v, want %v", got, tt.want)
			}
		})
	}
}