package main

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title" binding:"required"`
	IsDone bool   `json:"isDone"`
}
