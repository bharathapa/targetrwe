package services

import (
	"database/sql"
	logMsg "person/internal/log"
	"person/internal/model"
)

type PersonService struct {
	conn *sql.DB
}

func NewPersonService(conn *sql.DB) *PersonService {
	return &PersonService{
		conn: conn,
	}
}

//GetPersons - returns all person details
func (p *PersonService) GetPersons() ([]model.Person, error) {
	ps := make([]model.Person, 0)
	rows, err := p.conn.Query("SELECT * FROM person")
	if err != nil {
		logMsg.ErrorLog(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var p model.Person
		err = rows.Scan(&p.Id, &p.UserID, &p.Firstname, &p.Lastname, &p.Address)
		if err != nil {
			logMsg.ErrorLog(err.Error())
		}
		ps = append(ps, p)
	}
	return ps, nil
}

//GetPersonByUserId - returns person detail by user id
func (p *PersonService) GetPersonByUserId(userId int) (model.Person, error) {
	sqlStatement := `SELECT * FROM person WHERE userId=$1;`
	var person model.Person
	row := p.conn.QueryRow(sqlStatement, userId)
	err := row.Scan(&person.Id, &person.UserID, &person.Firstname, &person.Lastname, &person.Address)
	switch err {
	case sql.ErrNoRows:
		logMsg.ErrorLog(err.Error())
	case nil:
		logMsg.ErrorLog("successfully fetched user data")
	default:
		logMsg.ErrorLog("panic from select statement")
	}
	return person, err
}

//GetPersonById - returns person detail by user id
func (p *PersonService) GetPersonById(id int) (model.Person, error) {
	sqlStatement := `SELECT * FROM person WHERE id=$1;`
	var person model.Person
	row := p.conn.QueryRow(sqlStatement, id)
	err := row.Scan(&person.Id, &person.UserID, &person.Firstname, &person.Lastname, &person.Address)
	switch err {
	case sql.ErrNoRows:
		logMsg.ErrorLog(err.Error())
	case nil:
		logMsg.ErrorLog("successfully fetched user data")
	default:
		logMsg.ErrorLog("panic from select statement")
	}
	return person, err
}

func (p *PersonService) Add(person model.Person) model.Person {

	sqlStatement := `
INSERT INTO person (id,userid,firstname,lastname,address)
VALUES ($1, $2, $3, $4, $5)
RETURNING id`
	id := 0
	err := p.conn.QueryRow(sqlStatement, person.Id, person.UserID, person.Firstname, person.Lastname, person.Address).Scan(&id)
	if err != nil {
		logMsg.ErrorLog(err.Error())
	}
	return person
}
func (p *PersonService) Update(pt model.Person) model.Person {

	sqlStatement := `
UPDATE person
SET userid = $2, firstname = $3, lastname= $4, address = $5
WHERE id = $1
RETURNING id,userid,firstname,lastname,address;`
	var person model.Person
	err := p.conn.QueryRow(sqlStatement, pt.Id, pt.UserID, pt.Firstname, pt.Lastname, pt.Address).Scan(&person.Id, &person.UserID, &person.Firstname, &person.Lastname, &person.Address)
	if err != nil {
		logMsg.ErrorLog(err.Error())
	}
	return person
}
func (p *PersonService) Delete(id int) (int64, error) {

	sqlStatement := `
DELETE from person 
WHERE id = $1;`
	res, err := p.conn.Exec(sqlStatement, id)
	if err != nil {
		logMsg.ErrorLog(err.Error())
	}

	affectedRows, err := res.RowsAffected()

	if err != nil {
		logMsg.ErrorLog(err.Error())
	}
	return affectedRows, err
}
func (p *PersonService) PopulateDb() {
	persons := []model.Person{
		{UserID: 1008, Id: 1008, Firstname: " Test fname", Lastname: " Test lname", Address: " Test Address"},
		{UserID: 1009, Id: 1009, Firstname: " Test fname", Lastname: " Test lname", Address: " Test Address"},
		{UserID: 1010, Id: 1010, Firstname: " Test fname", Lastname: " Test lname", Address: " Test Address"},
		{UserID: 1011, Id: 1011, Firstname: " Test fname", Lastname: " Test lname", Address: " Test Address"},
	}
	for _, person := range persons {
		sqlStatement := `
INSERT INTO person (id,userid,firstname,lastname,address)
VALUES ($1, $2, $3, $4,$5)
RETURNING id`
		id := 0
		err := p.conn.QueryRow(sqlStatement, person.Id, person.UserID, person.Firstname, person.Lastname, person.Address).Scan(&id)
		if err != nil {
			logMsg.ErrorLog(err.Error())
		}
	}
}
