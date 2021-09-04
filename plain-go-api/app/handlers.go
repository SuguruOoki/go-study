package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

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
