
# Person

Person is a simple CRUD application written in go which exposes API endpoint to create the person.

## Installation

Install docker in your local system

Run the following command to setup the local postgress db

```bash
#docker run -p 5445:5445 -d \
>       -e POSTGRES_PASSWORD=postgres \
>       -e POSTGRES_USER=postgres \
>       -e POSTGRES_DB=targetrwe \
>       -v pgdata:/var/lib/postgresql/data \
>       postgres
>
#docker ps
#docker exec -it 1d096bc3c434 psql -U postgres targetrwe
targetrwe=# CREATE TABLE person (
targetrwe(#  id integer NOT NULL,
targetrwe(#  userid serial NOT NULL,
targetrwe(#  fistname character varying(255),
targetrwe(#  lastname  character varying(255),
targetrwe(#  address  character varying(255)
targetrwe(# );

targetrwe=# INSERT INTO person (id,userid,fistname,lastname,address)
targetrwe-# VALUES(1,100,'firstname','lastname','address');
INSERT 0 1
```
No need to insert the record but its your choice. Application will automatically insert few records on startup.

## Running the application

```
make run 
```

The above command will run the application
## Endpoints

Use postman or similar tool to call the API

1. http:\\localhost:8080/ , GET
2. http:\\localhost:8080/persons , GET 

Endpoint 1,2 serve the same purpose, i.e. will list all the persons

http:\\localhost:8080/person, method POST to create new persons

```
{
        "id": 1,
        "userid": 10000,
        "firstname": "John",
        "lennon": "Lenon",
        "address": "111 ABC Way, Dallas, TX"
    }
```
 http:\\localhost:8080/person, method POST to create new person

```
{
        "id": 1,
        "userid": 10000,
        "firstname": "John",
        "lennon": "Lenon",
        "address": "111 ABC Way, Dallas, TX"
    }
```

http:\\localhost:8080/persons/{ID}, method DELETE to delete person


http:\\localhost:8080/persons, method PUT to update person

```
{
        "id": 1,
        "userid": 10000,
        "firstname": "John",
        "lennon": "Lenon",
        "address": "111 ABC Way, Dallas, TX"
    }
```

http:\\localhost:8080/person, method POST to create new person

```
{
        "id": 1,
        "userid": 10000,
        "firstname": "John",
        "lennon": "Lenon",
        "address": "111 ABC Way, Dallas, TX"
    }
```

http:\\localhost:8080/persons/{id}, method GET to get persons by ID

http:\\localhost:8080/persons/user/{id}, method GET to get person by user ID


myRouter.HandleFunc("/", getAllPerson).Methods(constants.GET)
myRouter.HandleFunc("/persons", getAllPerson).Methods(constants.GET)
myRouter.HandleFunc("/persons/{id}", getPersonById).Methods(constants.GET)
myRouter.HandleFunc("/persons/user/{id}", getPersonByUserId)
myRouter.HandleFunc("/person", createPerson).Methods(constants.POST)
myRouter.HandleFunc("/persons", updatePerson).Methods(constants.PUT)
myRouter.HandleFunc("/persons/{id}", delete).Methods(constants.DELETE)