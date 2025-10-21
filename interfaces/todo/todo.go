package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Text     string `json:"text"`
}

func (t Todo) String() string {
	return fmt.Sprintf("Todo: %s", t.Text)
}

func (t Todo) Save() error {
	filename := "todo.json"

	json, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)
}

func New(text string) *Todo {
	return &Todo{Text: text}
}