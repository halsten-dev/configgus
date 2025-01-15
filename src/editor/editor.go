package editor

import (
	"log"

	"github.com/halsten-dev/configgus/data"

	tea "github.com/charmbracelet/bubbletea"
)

type Editor struct {
	config *data.Config
}

var _ tea.Model = (*Editor)(nil)

func newEditor(config *data.Config) *Editor {
	return &Editor{
		config: config,
	}
}

func (m *Editor) Init() tea.Cmd {
	return nil
}

func (m *Editor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *Editor) View() string {
	return "Hello world"
}

func Run(config *data.Config) {
	p := tea.NewProgram(newEditor(config), tea.WithAltScreen())

	if err, _ := p.Run(); err != nil {
		log.Fatal(err)
	}
}
