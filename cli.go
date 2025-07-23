package main

import (
	"fmt"
	"log"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

const playbookBasePath = "ansible/"

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

func (o os) PlaybookPath() string {
	switch o {
	case wsl:
		return "windows"
	case arch:
		return "arch"
	default:
		return "none"
	}
}

func (p profile) PlaybookPath() string {
	switch p {
	case dev:
		return "dev"
	case hyprland:
		return "hyprland"
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
	err            error
	done           bool
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

type playbookFinishedMsg struct{ err error }

func execPlaybook(o os, p profile) tea.Cmd {
	c := exec.Command("ansible-playbook", playbookBasePath+o.PlaybookPath()+"/"+p.PlaybookPath()+".yml", "--ask-become-pass")
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return playbookFinishedMsg{err}
	})
}

func updateOsChosen(m model, msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
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
			return m, execPlaybook(m.osChosen, m.profileChosen)
		}

	case playbookFinishedMsg:
		if msg.err != nil {
			m.err = msg.err
			fmt.Println(m.err)
			// return m, tea.Quit
			m.done = false
		} else {
			m.prompt = "Playbook finished! You can close this with CTRL+Q or ESC"
			fmt.Println(m.prompt)
			m.done = true
			m.err = nil
			// return m, tea.Quit
		}
		m.profileChosen = 0
		m.cursor = 0
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
	if m.done {
		return m.prompt + "\n\n" + viewOsChosen(m)
	} else if m.err != nil {
		return "Error occured: " + m.err.Error() + "\n\n" + viewOsChosen(m)
	} else if m.osChosen == 0 {
		return viewChooseOs(m)
	} else {
		return viewOsChosen(m)
	}
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
