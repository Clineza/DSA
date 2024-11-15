package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/harmonica"
	"os"
	"strconv"
)

type model struct {
	cursor         int
	numbers        []int
	spring         harmonica.Spring
	numProjectiles []*harmonica.Projectile
}

func initialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	for _, letter := range m.numbers {
		s += strconv.Itoa(letter) + " "
	}
	s += "\n"

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func main() {

	initPos := harmonica.Point{X: 0, Y: 0}
	initVel := harmonica.Vector{X: 0, Y: 0}
	initAcc := harmonica.TerminalGravity
	pro := *harmonica.NewProjectile(harmonica.FPS(60), initPos, initVel, initAcc)
	m := model{
		numbers:        []int{1, 5, 2, 3, 6, 12},
		spring:         harmonica.NewSpring(harmonica.FPS(60), 6.0, 0.5),
		numProjectiles: []*harmonica.Projectile{&pro},
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

/*

so i need an array of projectiles.
each projectile will have a value equal to what's the base

*/
