package main

import "fmt"

var currentId int

var memos Memos

// Give us some seed data
func init() {
	RepoCreateMemo(Memo{Title: "Write presentation"})
	RepoCreateMemo(Memo{Title: "Host meetup"})
}

func RepoFindMemo(id int) Memo {
	for _, m := range memos {
		if m.Id == id {
			return m
		}
	}
	// return empty Memo if not found
	return Memo{}
}

func RepoCreateMemo(m Memo) Memo {
	currentId += 1
	m.Id = currentId
	memos = append(memos, m)
	return m
}

func RepoDestroyMemo(id int) error {
	for i, m := range memos {
		if m.Id == id {
			memos = append(memos[:i], memos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Memo with id of %d to delete", id)
}
