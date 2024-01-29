package backlogtracker

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type YamlStorage struct {
  Counter uint32
  Backlog Backlog
}

// Bad performant solution
func GetBacklog() YamlStorage {
  state := YamlStorage{}
  home, err := os.UserHomeDir()
  if err != nil {
    panic(err)
  }
  tasksDir := fmt.Sprintf("%s/.local/state/backlog-tracker/tasks.yaml", home)
  dat, err:= os.ReadFile(tasksDir)
  if err != nil {
    fmt.Println("Creating new backlog file ...")
    state.Counter = 1
    SaveBacklog(state.Backlog, state.Counter)
    fmt.Println("File created successfully")
  }
  err = yaml.Unmarshal(dat, &state)
  return state
}

func SaveBacklog(backlog Backlog, counter uint32) error {
  state := YamlStorage{
    Counter: counter,
    Backlog: backlog,
  }
  dat, err := yaml.Marshal(&state)
  if err != nil {
    panic(err)
  }
  home, err := os.UserHomeDir()
  if err != nil {
    panic(err)
  }
  tasksDir := fmt.Sprintf("%s/.local/state/backlog-tracker/tasks.yaml", home)
  os.WriteFile(tasksDir, dat, 0644)
  return nil
}
