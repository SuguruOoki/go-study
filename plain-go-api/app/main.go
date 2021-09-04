package main

import (
	"fmt"
	"log"
	"net/http"

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
	fmt.Fprintln(w, "Memo Index!")
}

func MemoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memoId := vars["memoId"]
	fmt.Fprintln(w, "Memo show:", memoId)
}