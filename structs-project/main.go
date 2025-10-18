package main

import (
	"fmt"
	commands "structs-project/commands"
	notes "structs-project/notes"
	validations "structs-project/validations"
)

func getNoteData() (string, string) {
	title, _ := commands.PromptUntilValid("Please enter your note title: ", validations.ValidateNonEmpty)

	description, _ := commands.PromptUntilValid("Please enter your note description: ", validations.ValidateNonEmpty)

	fmt.Println(title, description)
	return title, description
}

func main() {
	title, description := getNoteData()

	fmt.Println(title, description)
	note := notes.NewNote(title, description)

	err := note.Save()
	if err != nil {
		fmt.Println("Error saving note:", err)
	} else {
		fmt.Println("Note saved successfully")
	}
}
