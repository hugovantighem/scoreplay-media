package main

import (
	"context"
	"log"
	"myproject/infra"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	var conf infra.Config
	if err := envconfig.Process(ctx, &conf); err != nil {
		log.Fatal(err)
	}

	logrus.SetLevel(logrus.DebugLevel)

	done := make(chan os.Signal)

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	stop := infra.RunApplication(conf)
	defer stop()

	<-done

}
