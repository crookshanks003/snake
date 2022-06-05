package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/crookshanks003/snake/models"
)

const height = 34
const width = 70
const fps = 100

type TickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Duration(fps)*time.Millisecond, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

type model struct {
	screen   models.Screen
	score    int
	gameOver bool
}

func InitialModel() model {
	return model{
		screen:   models.NewScreen(height, width),
		score:    0,
		gameOver: false,
	}
}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "right":
			m.screen.Snake.ChangeDir(0)
			return m, nil
		case "down":
			m.screen.Snake.ChangeDir(1)
			return m, nil
		case "left":
			m.screen.Snake.ChangeDir(2)
			return m, nil
		case "up":
			m.screen.Snake.ChangeDir(3)
			return m, nil
		}

	case TickMsg:
		m.screen.UpdateSnakePos()
		return m, tick()
	}

	return m, nil
}

func (m model) View() string {
	return m.screen.Render()
}
