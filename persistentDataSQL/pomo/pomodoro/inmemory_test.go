//go:build inmemory
// +build inmemory

package pomodoro_test

import (
	"testing"

	"github.com/fkl13/pcla/persistentDataSQL/pomo/pomodoro"
	"github.com/fkl13/pcla/persistentDataSQL/pomo/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()

	return repository.NewInMemoryRepo(), func() {}
}
