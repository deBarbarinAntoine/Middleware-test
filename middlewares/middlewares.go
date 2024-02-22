package middlewares

import (
	"Middleware-test/logs"
	"Middleware-test/models"
	"Middleware-test/utils"
	"log"
	"log/slog"
	"net/http"
)

var jsonHandler = slog.NewJSONHandler(logs.Log, nil)
var Logger = slog.New(jsonHandler)
var logId = 0

func Log() models.Middleware {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			Logger.Info("Log() Middleware", slog.Int("reqId", logId), slog.String("clientIP", utils.GetIP(r)), slog.String("reqMethod", r.Method), slog.String("reqURL", r.URL.String()))
			logId++
			handler.ServeHTTP(w, r)
		}
	}
}

func Guard() models.Middleware {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println("Guard()")
			handler.ServeHTTP(w, r)
		}
	}
}
