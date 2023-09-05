build-app:
	docker build -f build/Dockerfile -t mauricio1998/tech-challenge .

build-db:
	docker build -f build/Dockerfile.db -t mauricio1998/tech-db ./build

build-all: build-app build-db
	
run:
	docker-compose -f deployments/docker-compose.yml up -d