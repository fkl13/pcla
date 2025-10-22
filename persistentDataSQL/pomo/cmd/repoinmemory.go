//go:build inmemory || containers
// +build inmemory containers

package cmd

import (
	"github.com/fkl13/pcla/persistentDataSQL/pomo/pomodoro"
	"github.com/fkl13/pcla/persistentDataSQL/pomo/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
	return repository.NewInMemoryRepo(), nil
}
