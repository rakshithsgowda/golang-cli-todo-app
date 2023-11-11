package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

// simple item row in the todo list
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

// functions
func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	ls := *t
	if index <= 0 || index >= len(ls) {
		return errors.New("invalid index")
	}
	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true
	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")

	}
	*t = append(ls[:index-1], ls[index:]...)
	return nil
}

// read the file unmarshal the data
func (t *Todos) Load(filename string) error {
	filedata, err := os.ReadFile(filename)
	// filenot found and reading problem and file returned without data
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(filedata) == 0 {
		return err
	}
	err = json.Unmarshal(filedata, t)
	if err != nil {
		return err
	}
	return nil
}

// write back to the file
func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// print func
// Count if pending
