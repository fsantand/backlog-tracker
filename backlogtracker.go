package backlogtracker

import (
	"fmt"
	"golang.org/x/exp/slices"
)

type Story struct {
  Id uint32
  Title string
  Status Status
  Done bool
}

func (s Story) GetStatusDisplay() string {
  return ColorizeString(
    fmt.Sprintf("(%s)", s.Status),
    StatusColorCode[s.Status],
  )
}

type Status string

const (
  ToDo Status = "To Do"
  OnGoing Status = "On Going"
  OnReview Status = "On Review"
  Monitoring Status = "Monitoring"
  Completed Status = "Completed"
)

type Statuses []Status
type Backlog []Story

func AddStoryToBacklog(id uint32, title string, backlog *Backlog) Story {
  story := Story{
    Id: id,
    Title: title,
    Status: ToDo,
    Done: false,
  }
  *backlog = append(*backlog, story)
  return story
}

func PrintAllStories(backlog Backlog)  {
  for _, story := range backlog {
    fmt.Println(DebugStory(story, true))
  }
}

var StatusColorCode = map[Status]Color {
  ToDo: Gray,
  OnGoing: Cyan,
  OnReview: Yellow,
  Monitoring: Purple,
  Completed: Green,
}

func DebugStory(story Story, colorized bool) string {
  status := string(story.Status)
  if colorized {
    status = story.GetStatusDisplay()
  }
  return fmt.Sprintf(
    "#%d: %s %s",
    story.Id,
    story.Title,
    status,
  )
}

func FilterBacklog(backlog Backlog, include []Status) Backlog {
  var filtered Backlog
  for _, story := range backlog {
    if slices.Contains(include, story.Status) {
      filtered = append(filtered, story)
    }
  }

  return filtered
}
