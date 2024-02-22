package logs

import "os"

var Log, _ = os.Create("logs/logs.log")
