# Moria Gate API (EN)

## Overview
Moria Gate API is a secure microservice that protects sensitive personal and business data by using blockchain principles (Merkle Tree), SHA256 hashing, and AES encryption. It works as a secure gateway that returns consistent, deterministic outputs without persisting or exposing the original documents.

## Motivation
In payment systems and compliance-heavy environments, leaking CPF or CNPJ can result in fraud, extortion, and identity theft. Moria Gate acts as a protective filter — the raw data is never stored or leaked, and external systems only see safe, predictable outputs.

## Architecture
- **Golang** with Clean Architecture
- **TLS (HTTPS)** exposed on port `8443`
- **PostgreSQL** used for mock document queries
- **SHA256 + AES-GCM** for encryption and hashing
- **Merkle Tree** for simulating inclusion proof
- **Docker + Makefile** for easy setup
- **Swagger**, **Prometheus**, and **Grafana** for documentation and monitoring

## Technologies
- `Go`, `gin`, `swaggo`, `testify`
- `Docker`, `Docker Compose`, `Makefile`
- `Prometheus`, `Grafana`, `Swagger UI`
- `pgx`, `merkle-tree` libs

## Endpoints
| Method | Path                | Description                       |
|--------|---------------------|-----------------------------------|
| POST   | /v1/secure-data     | Returns hashed and encrypted data |
| GET    | /metrics            | Prometheus metrics                |
| GET    | /swagger/index.html | Swagger API documentation         |

## Ports
| Service     | Port  |
|-------------|-------|
| HTTPS API   | 8443  |
| PostgreSQL  | 5432  |
| Grafana     | 3000  |
| Prometheus  | 9090  |

## Makefile Commands
```bash
make env        # Set up Docker, build containers, and seed PostgreSQL
make start      # Start the API manually (runs main.go)
make swagger    # Open Swagger UI in the browser
make grafana    # Open Grafana dashboard
make docs       # Generate Swagger documentation in English and Portuguese
make docs-en    # Generate Swagger docs in English only
make docs-pt    # Generate Swagger docs in Portuguese only
make test       # Run all unit tests with verbose output
make coverage   # Run tests and open HTML coverage report
```

## Monitoring
- Grafana dashboard: auto-provisioned
- Alert email: `YOUR-EMAIL-HERE@EMAIL.COM`
- Prometheus: http://localhost:9090
- Swagger UI: https://localhost:8443/swagger/index.html

## Example Requests
```bash
curl -k -s -X POST https://localhost:8443/v1/secure-data \
  -H "Content-Type: application/json" \
  -d '{"document": "12345678901"}' | jq

# Repeat for other documents:
# 23456789012, 34567890123, 87654321, 76543210, 65432109
```

---

# Moria Gate API (PT)

## Visão Geral
A Moria Gate API é um microsserviço seguro que protege dados sensíveis (CPF e CNPJ) utilizando princípios de blockchain (Árvore de Merkle), hashing com SHA256 e criptografia AES-GCM. Atua como uma camada segura que retorna saídas determinísticas sem armazenar nem expor os dados reais.

## Motivação
Em sistemas de pagamento e ambientes regulados, vazamentos de CPF ou CNPJ podem gerar fraudes e extorsões. A Moria Gate funciona como um filtro de segurança — os dados brutos nunca são armazenados nem expostos, e os sistemas externos veem apenas respostas seguras.

## Arquitetura
- **Golang** com Clean Architecture
- **TLS (HTTPS)** exposto na porta `8443`
- **PostgreSQL** simula base de dados segura
- **SHA256 + AES-GCM** para proteção
- **Árvore de Merkle** para simular prova de integridade
- **Docker + Makefile** para setup simplificado
- **Swagger**, **Prometheus** e **Grafana** para observabilidade

## Tecnologias
- `Go`, `gin`, `swaggo`, `testify`
- `Docker`, `Docker Compose`, `Makefile`
- `Prometheus`, `Grafana`, `Swagger UI`
- `pgx`, `merkle-tree`

## Endpoints
| Método | Caminho              | Descrição                          |
|--------|----------------------|------------------------------------|
| POST   | /v1/secure-data      | Retorna dado criptografado         |
| GET    | /metrics             | Métricas Prometheus                |
| GET    | /swagger/index.html  | Documentação Swagger               |

## Portas
| Serviço     | Porta |
|-------------|-------|
| API HTTPS   | 8443  |
| PostgreSQL  | 5432  |
| Grafana     | 3000  |
| Prometheus  | 9090  |

## Comandos Makefile
```bash
make env        # Configura o Docker, constrói os containers e popula o PostgreSQL
make start      # Inicia a API manualmente (executa o main.go)
make swagger    # Abre a interface do Swagger no navegador
make grafana    # Abre o painel do Grafana
make docs       # Gera a documentação Swagger em inglês e português
make docs-en    # Gera apenas a documentação Swagger em inglês
make docs-pt    # Gera apenas a documentação Swagger em português
make test       # Executa todos os testes unitários com saída detalhada
make coverage   # Gera e abre o relatório de cobertura dos testes em HTML
```

## Monitoramento
- Dashboard Grafana: auto importado
- Alerta: `YOUR-EMAIL-HERE@EMAIL.COM`
- Prometheus: http://localhost:9090
- Swagger: https://localhost:8443/swagger/index.html

## Exemplo de Requisições
```bash
curl -k -s -X POST https://localhost:8443/v1/secure-data \
  -H "Content-Type: application/json" \
  -d '{"document": "12345678901"}' | jq

# Também disponíveis:
# 23456789012, 34567890123, 87654321, 76543210, 65432109
```
