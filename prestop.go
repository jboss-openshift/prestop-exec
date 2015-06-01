package main

import (
	"flag"
	"os"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/jboss-openshift/prestop-exec/sources"
	"fmt")

var argPollDuration = flag.Duration("poll_duration", 10*time.Second, "Polling duration")
var argPollAttempts	= flag.Int("poll_attempts", 1000, "Polling attempts")

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
	counter := 0
	for {
		select {
		case <-ticker.C:
			if counter > *argPollAttempts {
				return fmt.Errorf("Timed out!")
			}
			counter++

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
