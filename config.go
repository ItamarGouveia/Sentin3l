package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ServerURL          string `json:"server_url"`
	AuthToken          string `json:"auth_token"`
	IntervalSeconds    int    `json:"interval_seconds"`
	RetryInterval      int    `json:"retry_interval_seconds"`
	Verbose            bool   `json:"verbose"`
	CollectCPU         bool   `json:"collect_cpu"`
	CollectMemory      bool   `json:"collect_memory"`
	CollectDisk        bool   `json:"collect_disk"`
	CollectNetwork     bool   `json:"collect_network"`
	CollectProcesses   bool   `json:"collect_processes"`
	CollectServices    bool   `json:"collect_services"`
	CollectSoftware    bool   `json:"collect_software"`
	CollectBattery     bool   `json:"collect_battery"`
	CollectTemperature bool   `json:"collect_temperature"`
}

func LoadConfig(path string) Config {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Não foi possível ler o arquivo de configuração, usando valores padrão:", err)
		return Config{
			ServerURL:        "http://localhost:8080/api/metrics",
			IntervalSeconds:  5,
			Verbose:          false,
			CollectCPU:       true,
			CollectMemory:    true,
			CollectDisk:      true,
			CollectNetwork:   true,
			CollectProcesses: true,
			CollectServices:  true,
			CollectSoftware:  true,
		}
	}

	var cfg Config
	if err := json.Unmarshal(file, &cfg); err != nil {
		fmt.Println("Erro ao parsear JSON, usando valores padrão:", err)
		return Config{
			ServerURL:        "http://localhost:8080/api/metrics",
			IntervalSeconds:  5,
			Verbose:          false,
			CollectCPU:       true,
			CollectMemory:    true,
			CollectDisk:      true,
			CollectNetwork:   true,
			CollectProcesses: true,
			CollectServices:  true,
			CollectSoftware:  true,
		}
	}

	return cfg
}
