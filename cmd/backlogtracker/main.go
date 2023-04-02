package main

import (
  "github.com/fsantand/backlog-tracker"
)

func AddNewStory(title string, state *backlogtracker.YamlStorage) backlogtracker.Story {
  story := backlogtracker.AddStoryToBacklog(state.Counter, title, &state.Backlog)
  *&state.Counter++
  return story
}

func main() {
  state := backlogtracker.GetBacklog()

  filters := []backlogtracker.Status{
    backlogtracker.Monitoring,
    backlogtracker.ToDo,
    backlogtracker.OnReview,
    backlogtracker.OnGoing,
    backlogtracker.Completed,
  }

  filtered := backlogtracker.FilterBacklog(state.Backlog, filters)
  backlogtracker.PrintAllStories(filtered)

  backlogtracker.SaveBacklog(state.Backlog, state.Counter)
}
