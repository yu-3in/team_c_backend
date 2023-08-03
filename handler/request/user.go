package request

type ReqUpdateUser struct {
	ID             int    `param:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	DepartmentName string `json:"departmentName"`
	ProductName    string `json:"productName"`
}

type ReqLoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReqUpdateUserGenre struct {
	GenreID []int `json:"genreIds"`
}

type ReqCreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
