package main

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

type SuccessResponseType[T any] struct {
	Message string `json:"message"`
	Code    uint8  `json:"code"`
	Data    T      `json:"data"`
}
