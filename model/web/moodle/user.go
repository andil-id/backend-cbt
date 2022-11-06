package moodle

type UserCreateRequest struct {
	Users []Users `json:"users"`
}

type Users struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}
