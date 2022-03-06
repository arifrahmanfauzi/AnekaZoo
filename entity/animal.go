package entity

//Person object for REST(CRUD)
type Animal struct {
	ID        string    `json:"id"`
	Name 	  string 	`json:"name"`
	Class     string 	`json:"class"`
	Legs      int    	`json:"legs"`
}