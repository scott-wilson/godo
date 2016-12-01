package main

import (
	"fmt"

	"github.com/scott-wilson/godo/task"
)

func main() {
	t, err := task.New("Test", "This is a test message.\nTest test test")

	if err != nil {
		fmt.Printf("No can do, boss. %s\n", err)
		return
	}

	fmt.Printf("ID '%x' %d\n", t.Id, len(t.Id))
	fmt.Printf("PARENTID '%x' %d\n", t.ParentId, len(t.ParentId))
	fmt.Printf("STATUS %d\n", t.Status)
	fmt.Printf("SUBMISSION %s\n", t.Submission)
	fmt.Printf("SUBMITTER %s\n", t.Submitter)
	fmt.Printf("TITLE %s\n", t.Title)
	fmt.Printf("MESSAGE %s\n", t.Message)

	err = t.Save()

	if err != nil {
		fmt.Printf("No Savey. %s\n", err)
		return
	}

	fmt.Println("Saved.")
}
