package backlogtracker_test

import (
  "testing"
  "github.com/fsantand/backlog-tracker"
)

func TestDebugStory(t *testing.T) {
  expected := "#1: Hello world On going"
  s := backlogtracker.Story{
    Id: 1,
    Title: "Hello world",
    Status: "On going",
    Done: false,
  }
  description := backlogtracker.DebugStory(s, false)

  if description != expected{
    t.Errorf("Expected %s do not match actual %s", expected, description)
  }
}
