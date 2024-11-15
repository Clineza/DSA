package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/harmonica"
)

type model struct {
	numbers  [2]int
	swapping bool
	anim     *harmonica.Animation
	err      error
}

func initialModel() model {
	return model{
		numbers:  [2]int{1, 2},
		swapping: false,
		anim:     nil,
		err:      nil,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen, tea.EnterMouseAllMode)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "s":
			if !m.swapping {
				m.swapping = true
				m.anim = harmonica.NewAnimation(
					harmonica.EaseInOutQuad,
					harmonica.NewFrame(0, 50, 1000*time.Millisecond), // Animation frames
				)
				return m, tea.Tick(time.Millisecond*10, func(t time.Time) tea.Msg {
					return t
				})
			}
		}
	case time.Time:
		if m.swapping {
			if m.anim.Advance() { // This was incorrectly named 'next' before
				m.numbers[0], m.numbers[1] = m.numbers[1], m.numbers[0]
				m.swapping = false
				m.anim = nil
				return m, nil
			}
			return m, tea.Tick(time.Millisecond*10, func(t time.Time) tea.Msg { return t })
		}
	}
	return m, nil
}

func (m model) View() string {
	var s string
	if m.swapping {
		position := int(m.anim.Position() * 50) // Animation goes from 0 to 100%, so we adjust for half the screen width
		s += fmt.Sprintf(" %s%d%s %s%d%s\n", strings.Repeat(" ", 50-position), m.numbers[0], strings.Repeat(" ", position), strings.Repeat(" ", position), m.numbers[1], strings.Repeat(" ", 50-position))
	} else {
		s += fmt.Sprintf(" %50d %50d\n", m.numbers[0], m.numbers[1])
	}
	s += "\nPress 's' to swap, 'q' to quit.\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
