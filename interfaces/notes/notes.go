package notes

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewNote(title, description string) *Note {
	return &Note{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (n *Note) Save() error {
	nowStr := n.CreatedAt.Format("20060102-150405")
	normalizedTitle := strings.ToLower(strings.ReplaceAll(n.Title, " ", "-"))
	// Remove all non-alphanumeric and non-hyphen characters
	processedTitle := ""
	for _, r := range normalizedTitle {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			processedTitle += string(r)
		}
	}
	fmt.Println(processedTitle)
	filename := fmt.Sprintf("%s-%s.json", processedTitle, nowStr)

	json, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)
}

func (n *Note) Update(title string, description string) {
	if title != "" {
		n.Title = title
	}
	if description != "" {
		n.Description = description
	}
	n.UpdatedAt = time.Now()
}

func (n *Note) String() string {
	return fmt.Sprintf("Title: %s\nDescription: %s\nCreatedAt: %s\nUpdatedAt: %s",
		n.Title, n.Description, n.CreatedAt.Format(time.RFC3339), n.UpdatedAt.Format(time.RFC3339))
}
