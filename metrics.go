package main

import (
	"os"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

type Metrics struct {
	AgentID     string            `json:"agent_id"`
	Hostname    string            `json:"hostname"`
	OS          string            `json:"os"`
	Timestamp   string            `json:"timestamp"`
	CPUPercent  float64           `json:"cpu_percent"`
	MemPercent  float64           `json:"mem_percent"`
	DiskPercent float64           `json:"disk_percent"`
	NetIO       map[string]uint64 `json:"net_io"`
	Processes   []string          `json:"processes"`
	Services    []string          `json:"services"`
	Software    []string          `json:"software"`
	Temperature float64           `json:"temperature,omitempty"`
	Battery     float64           `json:"battery,omitempty"`
}

func collectMetrics(agentID string) Metrics {
	hostname, _ := os.Hostname()
	hostInfo, _ := host.Info()

	cpuPercents, _ := cpu.Percent(0, false)
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("/")

	netIOs, _ := net.IOCounters(true)
	netMap := make(map[string]uint64)
	for _, io := range netIOs {
		netMap[io.Name] = io.BytesSent + io.BytesRecv
	}

	procs, _ := process.Processes()
	procNames := []string{}
	for _, p := range procs {
		name, _ := p.Name()
		procNames = append(procNames, name)
	}

	// Serviços e software de exemplo
	services := []string{"service1", "service2"} // para produção, detecte serviços reais
	software := []string{"app1", "app2"}         // para produção, liste softwares reais

	// Temperatura e bateria (opcional, depende do SO)
	temp := 0.0
	bat := 0.0

	return Metrics{
		AgentID:     agentID,
		Hostname:    hostname,
		OS:          hostInfo.Platform,
		Timestamp:   time.Now().Format(time.RFC3339),
		CPUPercent:  cpuPercents[0],
		MemPercent:  vmStat.UsedPercent,
		DiskPercent: diskStat.UsedPercent,
		NetIO:       netMap,
		Processes:   procNames,
		Services:    services,
		Software:    software,
		Temperature: temp,
		Battery:     bat,
	}
}
