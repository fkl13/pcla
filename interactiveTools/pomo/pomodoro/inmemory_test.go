package pomodoro_test

import (
	"testing"

	"github.com/fkl13/pcla/interactiveTools/pomo/pomodoro"
	"github.com/fkl13/pcla/interactiveTools/pomo/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()

	return repository.NewInMemoryRepo(), func() {}
}
