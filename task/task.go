package task

import (
	"crypto/rand"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/scott-wilson/godo/config"
	"github.com/scott-wilson/godo/user"
)

const (
	UNCOMPLETED = iota
	COMPLETED
)

const idsize = 32

type Task struct {
	Id         [idsize]byte
	ParentId   [idsize]byte
	Status     int
	Submission time.Time
	Submitter  user.User
	Title      string
	Message    string
}

func New(title string, message string) (Task, error) {
	config, err := config.Read()

	if err != nil {
		return Task{}, err
	}

	submitter := config.User
	t := Task{Submitter: submitter, Title: title, Message: message}

	// Set unique ID value
	id := make([]byte, idsize)
	_, err = rand.Read(id)

	if err != nil {
		return t, err
	}

	copy(t.Id[:], id[:idsize])

	// Set submission time
	t.Submission = time.Now()

	// Set status
	t.Status = UNCOMPLETED

	return t, nil
}

func (t *Task) Load() error {
	config, err := config.Read()

	if err != nil {
		return err
	}

	fmt.Println(config.RootPath)

	return nil
}

func (t *Task) Save() error {
	config, err := config.Read()

	if err != nil {
		return err
	}

	f, err := os.Create(path.Join(config.RootPath, fmt.Sprintf("%x", t.Id)))
	defer f.Close()
	// Id         [idsize]byte
	// ParentId   [idsize]byte
	// Status     int
	// Submission time.Time
	// Submitter  user.User
	// Title      string
	// Message    string
	template := `---
id: %x
parent: %x
status: %d
submission %s
submitter: %s
---

%s

%s
`
	data := fmt.Sprintf(template, t.Id, t.ParentId, t.Status, t.Submission, t.Submitter, t.Title, t.Message)

	_, err = f.WriteString(data)

	if err != nil {
		return err
	}

	return nil
}
