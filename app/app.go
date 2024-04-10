package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type AppConfig struct {
	Port          string
	LogLevel      string
	MetricsEnable bool
	MetricsPort   string
	MetricsPath   string
}

var defaultAppConfig = &AppConfig{
	Port:          "3000",
	LogLevel:      "debug",
	MetricsEnable: false,
	MetricsPort:   "9090",
	MetricsPath:   "/metrics",
}

func MergeAppConfigsWithEnv() *AppConfig {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultAppConfig.Port
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = defaultAppConfig.LogLevel
	}

	metricEnable := os.Getenv("METRICS_ENABLE") == "true"
	if metricEnable {
		metricEnable = defaultAppConfig.MetricsEnable
	}

	metricPort := os.Getenv("METRICS_PORT")
	if metricPort == "" {
		metricPort = defaultAppConfig.MetricsPort
	}

	metricPath := os.Getenv("METRICS_PATH")
	if metricPath == "" {
		metricPath = defaultAppConfig.MetricsPath
	}

	return &AppConfig{
		Port:          port,
		LogLevel:      logLevel,
		MetricsEnable: metricEnable,
		MetricsPort:   metricPort,
		MetricsPath:   metricPath,
	}
}

type App struct {
	Configs *AppConfig
}

func NewApp(configs *AppConfig) *App {
	return &App{Configs: configs}
}

func (a *App) Start() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	fmt.Println("Listening on port " + a.Configs.Port)
	http.ListenAndServe(":"+a.Configs.Port, r)
}
