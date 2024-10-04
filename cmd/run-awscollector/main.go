package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	err := os.MkdirAll("/logs/adot", 0755)
	if err != nil {
		log.Fatal(err)
	}

	logfile, err := os.OpenFile("/logs/adot/adot.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()

	enableOtel, exists := os.LookupEnv("ENABLE_OTEL")
	if !exists || strings.ToLower(enableOtel) == "false" {
		log.Println("OpenTelemetry collector is disabled")
		for {
			time.Sleep(time.Duration(1<<63 - 1))
		}
	}

	cmd := exec.Command("./awscollector", "--config", "/etc/otel-config.yaml")
	cmd.Stdout = logfile
	cmd.Stderr = logfile
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
