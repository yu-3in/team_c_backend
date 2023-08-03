package request

type ReqUpdateUser struct {
	ID    int    `param:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ReqLoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReqUpdateUserGenre struct {
	GenreID []int `json:"genre_ids"`
}

type ReqCreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
