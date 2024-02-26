package main

import (
	"mini-go-app/internal/service/ginrouter"
	"mini-go-app/internal/util/envconf"
	"mini-go-app/internal/util/zaplog"
)

func main() {
	env := envconf.New()
	log := zaplog.New()
	r := ginrouter.New(env, log)

	log.Info(env.AppName + " is starting")
	if err := r.Run(":" + env.Port); err != nil {
		log.Sugar().Panic(err)
	}
}
