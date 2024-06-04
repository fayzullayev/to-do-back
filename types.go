package main

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}
