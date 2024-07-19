package models

type Task struct {
	ID 			int			`json:"id"`
	Name 		string 		`json:"name"`
	Description string 		`json:"description"`
	Status 		string 		`json:"status"`
}

type Response struct {
	Message 	string
}