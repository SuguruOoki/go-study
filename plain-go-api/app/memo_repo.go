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
	for _, t := range memos {
		if t.Id == id {
			return t
		}
	}
	// return empty Memo if not found
	return Memo{}
}

func RepoCreateMemo(t Memo) Memo {
	currentId += 1
	t.Id = currentId
	memos = append(memos, t)
	return t
}

func RepoDestroyMemo(id int) error {
	for i, t := range memos {
		if t.Id == id {
			memos = append(memos[:i], memos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Memo with id of %d to delete", id)
}
