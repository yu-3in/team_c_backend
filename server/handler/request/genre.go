package request

type ReqCreateGenre struct {
	Title string `json:"title"`
	Color string `json:"color"`
}

type ReqUpdateGenre struct {
	ID    int    `param:"id"`
	Title string `json:"title"`
	Color string `json:"color"`
}
