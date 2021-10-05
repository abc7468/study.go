package subtask

import (
	"fmt"
	"time"
)

const (
	UNKNOWN status = iota
	TODO
	DONE
)

type Task struct {
	Title    string    `json:"title,omitempty"`
	Status   status    `json:"status,omitempty"`
	Deadline *Deadline `json:"deadline,omitempty"`
	Priority int       `json:"priority,omitempty"`
	SubTasks []Task    `json:"subTasks,omitempty"`
}

type Deadline struct {
	time.Time
}

type status int

type IncludeSubTasks Task

func (t Task) String() string {
	check := "v"
	if t.Status != DONE {
		check = " "
	}
	return fmt.Sprintf("[%s] %s %s", check, t.Title, t.Deadline)
}

func (t IncludeSubTasks) indentedString(prefix string) string {
	str := prefix + Task(t).String()
	for _, st := range t.SubTasks {
		str += "\n" + IncludeSubTasks(st).indentedString(prefix+" ")
	}
	return str
}

func (t IncludeSubTasks) String() string {
	return t.indentedString("")
}
