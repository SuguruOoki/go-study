package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/memos", MemoIndex)
	router.HandleFunc("/memos/{memoId}", MemoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func MemoIndex(w http.ResponseWriter, r *http.Request) {
	memos := Memos{
		Memo{Title: "Write presentation"},
		Memo{Title: "Host meetup"},
	}

	json.NewEncoder(w).Encode(memos)
}

func MemoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memoId := vars["memoId"]
	fmt.Fprintln(w, "Memo show:", memoId)
}

type Memo struct {
	// NOTE: autoincrementの指定?
	Id    uint   `json:"id"`
	Title string `json:"title"`
	// NOTE: longText, mediumTextの指定?
	Content    string    `json:"content"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Memos []Memo
