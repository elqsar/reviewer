package entity

type Review struct {
	Id    int    `json:"id,omitempty"`
	Title string `json:"title"`
	Body  string `json:"body"`
	User  int    `json:"user"`
}
