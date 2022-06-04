package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func InitialModel() *model {
	return &model{}
}

func main() {
	p := tea.NewProgram(InitialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}