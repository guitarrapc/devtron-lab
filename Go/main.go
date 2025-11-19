package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"time"
)

type MachineInformation struct {
	MachineName    string `json:"machineName"`
	OSDescription  string `json:"osDescription"`
	ProcessorCount int    `json:"processorCount"`
	UsedMemoryInMB int64  `json:"usedMemoryInMB"`
}

type WeatherForecast struct {
	Date         string `json:"date"`
	TemperatureC int    `json:"temperatureC"`
	TemperatureF int    `json:"temperatureF"`
	Summary      string `json:"summary"`
}

var summaries = []string{
	"Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching",
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/healthz", handleHealth)
	http.HandleFunc("/weatherforecast", handleWeatherForecast)

	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	hostname, _ := os.Hostname()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	info := MachineInformation{
		MachineName:    hostname,
		OSDescription:  runtime.GOOS + " " + runtime.GOARCH,
		ProcessorCount: runtime.NumCPU(),
		UsedMemoryInMB: int64(m.Alloc / (1024 * 1024)),
	}

	json.NewEncoder(w).Encode(info)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Healthy"))
}

func handleWeatherForecast(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	forecasts := make([]WeatherForecast, 5)
	for i := 0; i < 5; i++ {
		date := time.Now().AddDate(0, 0, i+1)
		tempC := rand.Intn(75) - 20 // -20 to 55
		forecast := WeatherForecast{
			Date:         date.Format("2006-01-02"),
			TemperatureC: tempC,
			TemperatureF: 32 + int(float64(tempC)/0.5556),
			Summary:      summaries[rand.Intn(len(summaries))],
		}
		forecasts[i] = forecast
	}

	json.NewEncoder(w).Encode(forecasts)
}
