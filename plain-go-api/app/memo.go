package main

import "time"

type Memo struct {
	// NOTE: autoincrementの指定?
	Id    int    `json:"id"`
	Title string `json:"title"`
	// NOTE: longText, mediumTextの指定?
	Content    string    `json:"content"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Memos []Memo
