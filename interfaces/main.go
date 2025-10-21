package main

import (
	"fmt"
	commands "interfaces/commands"
	interfaces "interfaces/interfaces"
	notes "interfaces/notes"
	todos "interfaces/todo"
	validations "interfaces/validations"
)

func getNoteData() *notes.Note {
	title, _ := commands.PromptUntilValid("Please enter your note title: ", validations.ValidateNonEmpty)

	description, _ := commands.PromptUntilValid("Please enter your note description: ", validations.ValidateNonEmpty)

	fmt.Println(title, description)
	return notes.NewNote(title, description)
}

func getTodoData() *todos.Todo {
	text, _ := commands.PromptUntilValid("Please enter your todo text: ", validations.ValidateNonEmpty)
	return todos.New(text)
}

func outputData(data interfaces.Output) error{
	fmt.Println(data.String())
	return data.Save()
}

func main() {
	note := getNoteData()
	todo := getTodoData()

	err := outputData(todo)
	if err != nil {
		return
	}

	outputData(note)
}
