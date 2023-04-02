package backlogtracker

import (
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
  dat, err:= os.ReadFile("./tasks.yaml")
  if err != nil {
    panic(err)
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
  os.WriteFile("./tasks.yaml", dat, 0644)
  return nil
}
