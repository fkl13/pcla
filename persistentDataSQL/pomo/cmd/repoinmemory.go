package cmd

import (
	"github.com/fkl13/pcla/interactiveTools/pomo/pomodoro"
	"github.com/fkl13/pcla/interactiveTools/pomo/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
	return repository.NewInMemoryRepo(), nil
}
