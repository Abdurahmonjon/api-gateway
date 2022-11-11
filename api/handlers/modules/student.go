package modules

type Student struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
}

type StudentList struct {
	Students []Student `json:"students"`
}
