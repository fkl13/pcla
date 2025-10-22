//go:build !inmemory && !containers
// +build !inmemory,!containers

package cmd

import (
	"github.com/fkl13/pcla/persistentDataSQL/pomo/pomodoro"
	"github.com/fkl13/pcla/persistentDataSQL/pomo/pomodoro/repository"
	"github.com/spf13/viper"
)

func getRepo() (pomodoro.Repository, error) {
	repo, err := repository.NewSQLite3Repo(viper.GetString("db"))
	if err != nil {
		return nil, err
	}
	return repo, nil
}
