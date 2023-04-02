package backlogtracker_test

import (
  "testing"
  "github.com/fsantand/backlog-tracker"
)

func TestColorizeString(t *testing.T) {
  expected := "\033[32mHello World\033[0m"
  s := backlogtracker.ColorizeString(
    "Hello World",
    backlogtracker.Green,
  )

  if s != expected{
    t.Errorf("Expected %s do not match actual %s", expected, s)
  }
}
