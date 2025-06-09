# Moria Gate API — Technical E-Book (EN)

## 1. Introduction
Moria Gate API is inspired by the gates of Moria in Lord of the Rings: a hidden, secure entry that only opens with the right phrase. Similarly, our API protects sensitive documents and only provides derived results when invoked properly.

This e-book aims to help you fully understand how we implemented this project and why.

---

## 2. Motivation

- Avoid exposing sensitive data (e.g., CPF, CNPJ) in service-to-service calls
- Simulate real blockchain behavior (e.g., Merkle trees)
- Make every input always return the same output (determinism)
- Eliminate storage or caching: security = stateless + reproducible

---

## 3. Architecture Overview

### Components:

- **PostgreSQL** → Stores raw data
- **Go API (HTTPS)** → Query gateway
- **SHA256** → Consistent hashing of documents
- **AES** → Symmetric encryption of query response
- **Merkle Tree** → Generates root & allows inclusion proof
- **Prometheus & Grafana** → Observability
- **Swagger** → Docs

---

## 4. Technologies

- `gin-gonic`: Web framework
- `pgx`: PostgreSQL client
- `swaggo/swag`: Swagger documentation
- `stretchr/testify`: Unit testing
- `merkle`: Merkle Tree lib (GitHub open source)
- `crypto/sha256`, `crypto/aes`: Go native crypto
- `Prometheus client_golang`, `Grafana dashboards`

---

## 5. Secure Flow

1. Input (e.g. CPF `12345678901`)
2. Hash → `sha256(input)`
3. Use hash in Merkle Tree for proof generation
4. Query PostgreSQL and build encrypted response
5. Return AES-encrypted JSON
6. All logs redact the original input (use hash instead)

---

## 6. Key Code Snippets

### Merkle Tree Inclusion (simplified)
```go
leaf := sha256.Sum256([]byte(input))
proof := tree.GenerateProof(leaf[:])
```

### AES Encryption
```go
block, _ := aes.NewCipher(key)
stream := cipher.NewCFBEncrypter(block, iv)
stream.XORKeyStream(encrypted, plaintext)
```

### Endpoint Handler
```go
func (h *Handler) SecureData(ctx *gin.Context) {
    input := getDocument(ctx)
    hash := sha256(input)
    if !tree.Has(hash) {
        return 404
    }
    result := queryDB(input)
    response := encryptAES(result)
    ctx.JSON(200, response)
}
```

---

## 7. Observability

- Grafana dashboard: `/grafana`
- Metrics: `/metrics`
- Alerts via email if response time > 1s
- All via Prometheus queries

---

## 8. Final Thoughts

- Highly secure → no data exposed
- Scalable → stateless, fast, simple
- Ready for integration with ZKPs

# Moria Gate API — Livro Técnico (PT-BR)

## 1. Introdução
A Moria Gate API foi inspirada nos portões de Moria de Senhor dos Anéis: uma entrada segura e protegida que só se abre com a frase certa. Da mesma forma, esta API protege dados sensíveis e apenas fornece resultados derivados sob demanda.

Este e-book explica em detalhes a implementação, as decisões e os conceitos envolvidos.

---

## 2. Motivação

- Evitar exposição de dados como CPF e CNPJ
- Simular comportamento real de blockchain (árvore de Merkle)
- Ter respostas determinísticas (mesmo input → mesmo output)
- Eliminar persistência: segurança via estateless e reproducibilidade

---

## 3. Arquitetura

### Componentes:

- **PostgreSQL** → Armazena dados reais
- **API Go (HTTPS)** → Funil de consulta
- **SHA256** → Hash consistente do documento
- **AES** → Criptografia simétrica da resposta
- **Árvore de Merkle** → Geração de raiz e prova de inclusão
- **Prometheus & Grafana** → Monitoramento
- **Swagger** → Documentação

---

## 4. Tecnologias

- `gin-gonic`: Framework web
- `pgx`: Cliente PostgreSQL
- `swaggo/swag`: Documentação Swagger
- `stretchr/testify`: Testes unitários
- `merkle`: Lib de árvore de Merkle
- `crypto/sha256`, `crypto/aes`: Criptografia Go nativa
- `Prometheus`, `Grafana`

---

## 5. Fluxo Seguro

1. Entrada (ex: CPF `12345678901`)
2. Geração de hash → `sha256(input)`
3. Verificação na árvore de Merkle
4. Consulta ao PostgreSQL
5. Criptografia da resposta (AES)
6. Resposta protegida em JSON
7. Logs nunca mostram input original

---

## 6. Exemplos de Código

### Árvore de Merkle
```go
leaf := sha256.Sum256([]byte(input))
proof := tree.GenerateProof(leaf[:])
```

### Criptografia AES
```go
block, _ := aes.NewCipher(key)
stream := cipher.NewCFBEncrypter(block, iv)
stream.XORKeyStream(encrypted, plaintext)
```

### Endpoint Seguro
```go
func (h *Handler) SecureData(ctx *gin.Context) {
    input := getDocument(ctx)
    hash := sha256(input)
    if !tree.Has(hash) {
        return 404
    }
    result := queryDB(input)
    response := encryptAES(result)
    ctx.JSON(200, response)
}
```

---

## 7. Observabilidade

- Dashboard Grafana: `/grafana`
- Métricas: `/metrics`
- Alertas por e-mail se tempo > 1s
- Baseado em Prometheus

---

## 8. Considerações Finais

- Altamente seguro
- Escalável
- Pronto para futuras integrações com ZKP