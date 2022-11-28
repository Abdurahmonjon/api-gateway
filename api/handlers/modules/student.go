package modules

type Student struct {
	UserName  string `json:"user_name,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

type StudentList struct {
	Students []Student `json:"students"`
}
