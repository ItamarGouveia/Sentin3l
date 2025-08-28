
# Sentin3l - Remote Monitoring Agent

**Sentin3l** é um agente de monitoramento remoto (RMM) desenvolvido em Go, projetado para coletar métricas de sistemas e enviá-las para um servidor central. Ele é leve, modular e configurável via JSON ou flags de linha de comando.

---

## Funcionalidades

- Coleta de métricas:
  - CPU, memória, disco
  - Rede (bytes enviados/recebidos)
  - Processos ativos
  - Serviços e softwares instalados
  - Bateria e temperatura (quando disponível)
- AgentID persistente
- Configuração via arquivo `rmm.json`
- Flags de linha de comando para sobrescrever configurações
- Intervalo de coleta configurável
- Logs detalhados opcionais

---

## Instalação

1. Clone o repositório:

```bash
git clone https://github.com/seu-usuario/sentin3l.git
cd sentin3l
```

2. Instale dependências:

```bash
go get github.com/shirou/gopsutil/...
go get github.com/google/uuid
```

3. Compile o agente:

```bash
go build -o sentin3l
```

---

## Configuração

Crie um arquivo `rmm.json` na mesma pasta do executável:

```json
{
  "server_url": "http://localhost:8080/api/metrics",
  "auth_token": "abcdef123456",
  "interval_seconds": 5,
  "retry_interval_seconds": 10,
  "verbose": true,
  "collect_cpu": true,
  "collect_memory": true,
  "collect_disk": true,
  "collect_network": true,
  "collect_processes": true,
  "collect_services": true,
  "collect_software": true,
  "collect_battery": true,
  "collect_temperature": true
}
```

---

## Uso

Executando o agente com configurações padrão do `rmm.json`:

```bash
./sentin3l
```

Sobrescrevendo configurações via flags:

```bash
./sentin3l -server http://meuservidor:8080/api/metrics -interval 10 -verbose
```

### Flags disponíveis

- `-interval <segundos>` → Intervalo de envio das métricas (sobrescreve o JSON)  
- `-server <url>` → URL do servidor (sobrescreve o JSON)  
- `-verbose` → Ativa logs detalhados  
- `-version` → Mostra versão do agente  
- `-help` → Mostra ajuda  

---

## Estrutura do Projeto

```
sentin3l/
├── main.go       # Loop principal do agente
├── help.go       # Processa flags e mostra ajuda
├── config.go     # Lê rmm.json e aplica configuração
├── metrics.go    # Coleta métricas do sistema
├── agentid.go    # Gera e mantém AgentID persistente
├── go.mod
└── rmm.json      # Arquivo de configuração
```

---

## Exemplos de Saída

### No servidor

```text
[2025-08-28T09:00:00] Métricas recebidas: {AgentID:PC-1234 Hostname:itapc OS:linux CPUPercent:12.5 MemPercent:28.5 DiskPercent:8.3 Timestamp:2025-08-28T09:00:00}
```

### No agente (verbose)

```text
Servidor: http://localhost:8080/api/metrics
Intervalo: 5 segundos
Métricas enviadas com status: 200
```

---

## Roadmap

- [ ] Adicionar autenticação com token no envio de métricas  
- [ ] Coleta de softwares e serviços específicos do SO  
- [ ] Implementar dashboard web (React ou Fyne)  
- [ ] Compilação multiplataforma (Windows, Linux, MacOS)  
- [ ] Inicialização automática como serviço  

---

## Licença

Este projeto está licenciado sob a licença MIT.
