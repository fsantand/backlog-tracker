package backlogtracker

import "fmt"

type Color string

const (
  Green Color = "\033[32m"
  Yellow Color = "\033[33m"
  Purple Color = "\033[35m"
  Cyan Color = "\033[36m"
  Gray Color = "\033[37m"
  White Color = ""
  Reset Color = "\033[0m"
)


func ColorizeString(to_colorize string, text_color Color) string {
  return fmt.Sprintf("%s%s%s", text_color, to_colorize, Reset)
}
