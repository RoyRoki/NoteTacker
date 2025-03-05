package model

type Note struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Note  string `json:"note"`
	Date  string `json:"date"`
}
