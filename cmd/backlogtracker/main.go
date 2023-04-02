package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/fsantand/backlog-tracker"
)

const (
  ListChoice string = "list"
  AddChoice string = "add"
  StatusChangeChoice string = "status_change"
)

type model struct {
  backlog backlogtracker.Backlog
  counter uint32
  filter, newStory textinput.Model
  statusChoice backlogtracker.Statuses 
  cursor int
  actualView string
}

func initialModel() model {
  state := backlogtracker.GetBacklog()
  
  return model {
    backlog: state.Backlog,
    counter: state.Counter,
    filter: textinput.New(),
    newStory: textinput.New(),
    statusChoice: backlogtracker.Statuses{
      backlogtracker.ToDo,
      backlogtracker.OnGoing,
      backlogtracker.OnReview,
      backlogtracker.Monitoring,
    },
    cursor: 0,
    actualView: ListChoice,
  }
}

func (m model) Init() tea.Cmd {
	return nil
}

func QuitUpdate(m model) (tea.Model, tea.Cmd) {
  backlogtracker.SaveBacklog(
    m.backlog,
    m.counter,
  )
  return m, tea.Quit
}

func ListViewUpdate(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd = nil

  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+q", "esc", "q":
      return QuitUpdate(m)

    case "down", "j": 
      m.cursor++
      if m.cursor >= len(m.backlog) {
          m.cursor = 0
      }

    case "up", "k":
      m.cursor--
      if m.cursor < 0 {
        m.cursor = len(m.backlog) - 1
      }
    case "a":
      m.actualView = AddChoice
      m.newStory.Focus()
      cmd = textinput.Blink
    }
  }

  return m, cmd
}

func AddViewUpdate(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd

  switch msg:= msg.(type) {
  case tea.KeyMsg:
    switch msg.Type {
    case tea.KeyEnter:
      backlogtracker.AddStoryToBacklog(
        m.counter,
        m.newStory.Value(),
        &m.backlog,
        )
      m.counter++
      m.actualView = ListChoice
    case tea.KeyEsc:
      m.actualView = ListChoice
    }
  }

  m.newStory, cmd = m.newStory.Update(msg)
  return m, cmd
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch m.actualView {
  case AddChoice:
    return AddViewUpdate(msg, m)
  case ListChoice:
    return ListViewUpdate(msg, m)
  }

  return QuitUpdate(m)
}

func ListView(m model) string {
  s := strings.Builder{}
  s.WriteString("Active backlog\n\n")

  for i, task := range(backlogtracker.FilterBacklog(m.backlog, m.statusChoice)) {
    actualTaskTitle := backlogtracker.DebugStory(task, true)
    if m.cursor == i {
      s.WriteString(fmt.Sprintf("  > %s", backlogtracker.Underline(actualTaskTitle)))
    } else {
      s.WriteString(fmt.Sprintf("  %s", actualTaskTitle))
    }
    s.WriteString("\n")
  }

  s.WriteString(backlogtracker.Subtle("\n(q quit)/(a add)/(enter change status)\n"))

  return s.String()
}

func AddView(m model) string {
  return fmt.Sprintf(
    "New story title\n\n%s\n\n%s",
    m.newStory.View(),
    backlogtracker.Subtle("(esc cancel)/(enter add)\n"),
  )
}

func (m model) View() string {
  switch m.actualView {
  case AddChoice:
    return AddView(m)
  case StatusChangeChoice:
    m.actualView = ListChoice
    return ListView(m)
  }
  return ListView(m)
}

func main() {
	if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
