package validation

import (
	"fmt"
	"person/internal/model"
)

func IsValidPerson(p model.Person) bool {
	fmt.Println(" is valid person", p)
	if p.Firstname == "" || p.Id == 0 || p.UserID == 0 {
		fmt.Println("firstname ", p.Firstname, " id ", p.Id, " userid ", p.UserID)
		return false
	}
	return true
}
