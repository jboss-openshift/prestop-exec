package main

import (
	"flag"
	"os"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/jboss-openshift/prestop-exec/sources"
)

var argPollDuration = flag.Duration("poll_duration", 10*time.Second, "Polling duration")

func main() {
	flag.Parse()
	glog.Infof(strings.Join(os.Args, " "))
	glog.Infof("Pre-stop exec version %v", prestopVersion)

	err := doWork()
	if err != nil {
		glog.Error(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func doWork() error {
	source := sources.NewSource()
	ticker := time.NewTicker(*argPollDuration)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			done, err := source.CheckProgress()
			if err != nil {
				return err
			}
			if done {
				glog.Info("Pre-stop check DONE")
				return nil
			}
		}
	}
	return nil
}
