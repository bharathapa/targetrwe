package constants

const (
	Host            = "localhost"
	PORT            = "8080"
	GET             = "GET"
	POST            = "POST"
	DELETE          = "DELETE"
	PUT             = "PUT"
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"

	//api endpoints
	GetAllPerson      = "/"
	GetAllPersons     = "/persons"
	GetPersonById     = "/persons/{id}"
	GetPersonByUserId = "/persons/user/{id}"
	CreatePerson      = "/person"
	UpdatePerson      = "/persons"
	DeletePerson      = "/persons/{id}"
)
