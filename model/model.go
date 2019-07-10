package model

//UserModel adalah model dari user kalau bisa sesuai dengan databasenya
type UserModel struct {
	ID        int    `json:"id"`
	UserName  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"-"`
}

//ResponseUser untuk hasil response dari user
type ResponseUser struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []UserModel
}
