build:
	docker build -f build/Dockerfile -t tech-challenge .

run:
	docker-compose -f deployments/docker-compose.yml up -d