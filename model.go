package main

import (
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/crookshanks003/snake/cons"
	"github.com/crookshanks003/snake/models"
)

type TickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Duration(cons.Fps)*time.Millisecond, func(t time.Time) tea.Msg {
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
		screen:   models.NewScreen(cons.Height, cons.Width),
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
		score, gameOver := m.screen.UpdateSnakePos()
		if score {
			m.score++
		}
		if gameOver {
			m.gameOver = true
			return m, nil
		}
		return m, tick()
	}

	return m, nil
}

func (m model) View() string {
	lastLine := "\nScore: " + strconv.Itoa(m.score)
	if m.gameOver {
		lastLine += "\t\t\tGame Over!!"
	} else {
		lastLine += "\t\t\t\t"
	}
	lastLine += "\t\t     q to quit"
	return m.screen.RenderScreen() + lastLine
}
