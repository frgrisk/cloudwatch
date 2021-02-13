package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/google/uuid"
	"github.com/hskiba/cloudwatch"
	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
	"os"
)

func main() {

	stream := uuid.New().String()

	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{FullTimestamp: true, DisableColors: true})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(log.TraceLevel)
	e := log.NewEntry(logger)

	cfg, err := config.LoadDefaultConfig(context.TODO())

	g := cloudwatch.NewGroup("github.com/hskiba/cloudwatch", cloudwatchlogs.NewFromConfig(cfg))

	w, err := g.Create(stream)
	if err != nil {
		log.Fatal(err)
	}
	logger.AddHook(&writer.Hook{
		Writer: w,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
			log.InfoLevel,
			log.DebugLevel,
			log.TraceLevel,
		},
	})

	e.Trace("Something very low level.")
	e.Debug("Useful debugging information.")
	e.Info("Something noteworthy happened!")
	e.Warn("You should probably take a look at this.")
	e.Error("Something failed but I'm not quitting.")

	w.Close()
}
