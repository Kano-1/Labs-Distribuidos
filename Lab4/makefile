docker-director:
	docker build -t director-build -f DockerfileServer .
	docker run -d --name director-container -p 50051:50051 director-build

	docker logs -f --tail 10 director-container

mercenarios-file:
	go run Mercenarios/mercenarios.go