package utils

import (
	"log"
	"log/slog"
	"os"
	"strconv"
	"time"
)

var Logger *slog.Logger

// LogInit is meant to be run as a goroutine to create a new log file every day
// appending the file's creation timestamp in its name.
func LogInit() {
	var logs *os.File
	var jsonHandler *slog.JSONHandler
	for {
		filename := "logs/logs_" + time.Now().Format(time.DateOnly) + "_" + strconv.Itoa(time.Now().Hour()) + "h" + strconv.Itoa(time.Now().Minute()) + ".log"
		var err error
		logs, err = os.Open(filename)
		if err != nil {
			var createErr error
			logs, createErr = os.Create(filename)
			if createErr != nil {
				log.Println(GetCurrentFuncName(), slog.Any("output", createErr))
			}
		}
		jsonHandler = slog.NewJSONHandler(logs, nil)
		Logger = slog.New(jsonHandler)
		time.Sleep(time.Hour * 24)
	}
}
