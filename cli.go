package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type os int

func (o os) String() string {
	switch o {
	case arch:
		return "Arch Linux"
	case wsl:
		return "Windows (WSL)"
	default:
		return "none"
	}
}

type profile int

func (p profile) String() string {
	switch p {
	case dev:
		return "Dev"
	case hyprland:
		return "Hyprland"
	default:
		return "none"
	}
}

const (
	dev profile = iota + 1
	hyprland
)

const (
	arch os = iota + 1
	wsl
)

type model struct {
	prompt         string
	osChoices      []os
	osChosen       os
	profileChoices map[os][]profile
	profileChosen  profile
	cursor         int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if msg, ok := msg.(tea.KeyMsg); ok {
		if msg.String() == "ctrl+q" || msg.String() == "esc" {
			return m, tea.Quit
		}
	}

	if m.osChosen == 0 {
		return updateChooseOs(m, msg)
	} else {
		return updateOsChosen(m, msg)
	}
}

func updateOsChosen(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.profileChoices[m.osChosen])-1 {
				m.cursor++
			}

		case "enter":
			m.profileChosen = m.profileChoices[m.osChosen][m.cursor]
		}
	}
	return m, nil
}

func updateChooseOs(m model, msg tea.Msg) (tea.Model, tea.Cmd) {

	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
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
			m.cursor = 0
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.osChosen == 0 {
		return viewChooseOs(m)
	} else if m.profileChosen == 0 {
		return viewOsChosen(m)
	} else {
		return viewProfileChosen(m)
	}
}

func viewProfileChosen(m model) string {
	return "You Chose " + m.profileChosen.String() + "\n\n"
}

func viewOsChosen(m model) string {
	s := "Please choose profile for " + m.osChosen.String() + "\n\n"

	profiles := m.profileChoices[m.osChosen]

	for i, p := range profiles {
		selected := " "
		if m.cursor == i {
			selected = "x"
		}
		s += fmt.Sprintf("[%s] %s\n", selected, p)
	}
	return s
}

func viewChooseOs(m model) string {
	s := "Select OS -\n\n"
	for i, os := range m.osChoices {
		selected := " "
		if i == m.cursor {
			selected = "x"
		}
		s += fmt.Sprintf("[%s] %s\n", selected, os)
	}
	return s
}

func main() {
	initialModel := model{
		prompt:    "TEST PROMPT",
		osChoices: []os{arch, wsl},
		profileChoices: map[os][]profile{
			arch: {dev, hyprland},
			wsl:  {dev},
		},
	}
	p := tea.NewProgram(initialModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("dotfiles: %v", err)
	}
}
