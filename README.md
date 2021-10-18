
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

## GET Requests

### 1. Get all persons.
###       Request URL:
            ```
            Request URL: http:\\localhost:8080/
            Request URL: http:\\localhost:8080/persons
            ```

### 2. Get person by id.
###       Request URL:
            ```
            Request URL: http:\\localhost:8080/persons/{id}, method GET to get persons by ID
            ```

### 3. Get person by userid.
        ```
        Request URL: http:\\localhost:8080/persons/user/{id}, method GET to get person by user ID
        ```


## POST Requests

### 1. Create new person.
###       Request URL:
            ```
            http:\\localhost:8080/person, method POST to create new persons
            ```
###       Body:
            ```
            {
                    "id": 1,
                    "userid": 10000,
                    "firstname": "John",
                    "lennon": "Lenon",
                    "address": "111 ABC Way, Dallas, TX"
                }
            ```


## PUT Requests

### 1. Update person details.
###       Request URL:
            ```
            http:\\localhost:8080/persons
            ```
###       Body:
            ```
            {
                    "id": 1,
                    "userid": 10000,
                    "firstname": "George",
                    "lennon": "Washington",
                    "address": "222 BBC Way, Miami, FL"
                }
            ```


## DELETE Requests

### 1. Update person details.
###       Request URL:
            ```
            http:\\localhost:8080/persons/{ID}
            ```