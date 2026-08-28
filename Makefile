.PHONY: up down logs


up:
	docker compose up -d

keycloak:
	docker compose up -d keycloak

kafka-test:
	docker exec -it nbm-kafka /opt/kafka/bin/kafka-topics.sh --list --bootstrap-server localhost:9092

down:
	docker compose down


logs:
	docker compose logs -f