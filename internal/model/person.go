package model

type Person struct {
	UserID    int    `json:"userid"`
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Address   string `json:"address"`
}
