package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	prompt    string
	osChoices []string
	osChosen  string
	cursor    int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q", "esc":
			m.prompt = "exiting..."
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.osChoices)-1 {
				m.cursor++
			}

		case "enter":
			m.osChosen = m.osChoices[m.cursor]
		}
	}

	return m, nil
}

func (m model) View() string {
	var s string = ""

	if m.osChosen == "" {
		s += "Select OS -\n\n"
		for i, os := range m.osChoices {
			selected := " "
			if i == m.cursor {
				selected = "x"
			}
			s += fmt.Sprintf("[%s] %s\n", selected, os)
		}
	} else {
		s += "Great! You chose " + m.osChosen
	}

	return s
}

func main() {
	initialModel := model{
		prompt:    "test",
		osChoices: []string{"Arch Linux", "Windows(wsl)"},
	}
	p := tea.NewProgram(initialModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("dotfiles: %v", err)
	}
}
