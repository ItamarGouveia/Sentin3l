package main

import (
	"flag"
	"fmt"
	"os"
)

// Mostra ajuda
func showHelp() {
	fmt.Println("RMM Agent - Help")
	fmt.Println()
	fmt.Println("Parâmetros disponíveis:")
	fmt.Println("  -interval <segundos>   Intervalo de envio das métricas (sobrescreve rmm.json)")
	fmt.Println("  -server <url>          URL do servidor (sobrescreve rmm.json)")
	fmt.Println("  -verbose               Ativa logs detalhados")
	fmt.Println("  -version               Mostra a versão do agente")
	fmt.Println("  -help                  Mostra esta tela de ajuda")
	fmt.Println()
	os.Exit(0)
}

// Processa flags e retorna valores que sobrescrevem rmm.json
func parseFlags() (serverURL string, interval int, verbose bool) {
	showVersion := flag.Bool("version", false, "Mostra versão do agente")
	help := flag.Bool("help", false, "Mostra ajuda")
	verboseFlag := flag.Bool("verbose", false, "Ativa logs detalhados")
	server := flag.String("server", "", "URL do servidor (sobrescreve rmm.json)")
	intervalFlag := flag.Int("interval", 0, "Intervalo de envio em segundos (sobrescreve rmm.json)")

	flag.Parse()

	if *help {
		showHelp()
	}

	if *showVersion {
		fmt.Println("RMM Agent v1.0")
		os.Exit(0)
	}

	return *server, *intervalFlag, *verboseFlag
}
