package main

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

type SpeedInput struct {
	Focus bool
}

func NewSpeedInput() SpeedInput {
	return SpeedInput{Focus: true}
}

func (s *SpeedInput) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "1", "2", "3":
			val, err := strconv.Atoi(msg.String())
			if err != nil {
				return tea.Quit
			}
			fps = 50 * (4 - val)
			s.Focus = false
			return tick()
		default:
			return tea.Quit
		}
	}
	return nil
}

func (s *SpeedInput) View() string {
	ret := "Select speed, press (1: slow, 2: medium, 3: fast)\n\npress any other key to quit"
	return ret
}
