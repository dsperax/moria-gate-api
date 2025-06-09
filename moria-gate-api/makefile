.PHONY: env start swagger grafana docs docs-pt docs-en test coverage

env:
	docker compose up -d --build
	sleep 5
	docker exec postgres psql -U postgres -d securedata -f /docker-entrypoint-initdb.d/init.sql

start:
	go run ./cmd/api/main.go

swagger:
	xdg-open https://localhost:8443/swagger/index.html || open https://localhost:8443/swagger/index.html

grafana:
	xdg-open http://localhost:3000 || open http://localhost:3000

docs: docs-en docs-pt

docs-en:
	swag init -g cmd/api/main.go -o ./docs/en --parseDependency --parseInternal --outputTypes go,json,yaml --lang en

docs-pt:
	swag init -g cmd/api/main.go -o ./docs/pt --parseDependency --parseInternal --outputTypes go,json,yaml --lang pt

test:
	go test ./... -v

coverage:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
