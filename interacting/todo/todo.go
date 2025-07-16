package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, js, 0644)
}

func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}

// String prints out a formatted list
func (l *List) String() string {
	return l.Format(false)
}

func (l *List) Format(exclude bool) string {
	formatted := ""
	for k, t := range *l {
		if exclude && t.Done {
			continue
		}
		prefix := "  "
		if t.Done {
			prefix = "X "
		}
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}
	return formatted
}

func (l *List) StringVerbose(exclude bool) string {
	formatted := ""
	dateTimeFormat := "2 Jan 2006 15:04"

	maxTaskLenth := 0
	for _, t := range *l {
		if exclude && t.Done {
			continue
		}
		if len(t.Task) > maxTaskLenth {
			maxTaskLenth = len(t.Task)
		}
	}

	for k, t := range *l {
		if exclude && t.Done {
			continue
		}

		prefix := "  "
		completed := ""
		if t.Done {
			prefix = "X "
			completed += fmt.Sprintf(" Completed: %s", t.CompletedAt.Format(dateTimeFormat))
		}
		formatted += fmt.Sprintf("%s%d: %-*s Created: %s%s\n", prefix, k+1, maxTaskLenth, t.Task, t.CreatedAt.Format(dateTimeFormat), completed)
	}
	return formatted
}
