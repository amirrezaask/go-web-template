package config

import "github.com/golobby/config"

var C *config.Config

func Init(feeders ...config.Feeder) error {
	var err error
	C, err = config.New(feeders...)
	if err != nil {
		return err
	}
	return nil
}
