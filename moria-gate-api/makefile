.PHONY: env start swagger grafana

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
