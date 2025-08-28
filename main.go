package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func sendMetrics(serverURL string, m Metrics) {
	data, _ := json.Marshal(m)
	resp, err := http.Post(serverURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Erro ao enviar métricas:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Métricas enviadas com status:", resp.Status)
}

func main() {
	// Carrega configuração do JSON
	config := LoadConfig("rmm.json")

	// Sobrescreve com flags, se passadas
	serverURL, interval, verbose := parseFlags()
	if serverURL != "" {
		config.ServerURL = serverURL
	}
	if interval != 0 {
		config.IntervalSeconds = interval
	}
	config.Verbose = config.Verbose || verbose

	agentID := getAgentID()

	if config.Verbose {
		fmt.Println("Servidor:", config.ServerURL)
		fmt.Println("Intervalo:", config.IntervalSeconds, "segundos")
	}

	for {
		metrics := collectMetrics(agentID)
		sendMetrics(config.ServerURL, metrics)
		time.Sleep(time.Duration(config.IntervalSeconds) * time.Second)
	}
}
