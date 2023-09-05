build:
	docker build -f build/Dockerfile -t mauricio1998/tech-challenge .

run:
	docker-compose -f deployments/docker-compose.yml up -d