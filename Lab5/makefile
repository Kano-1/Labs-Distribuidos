.PHONY: broker server1 server2 server3

broker:
	docker build -t broker-go -f DockerfileBroker .
	docker run -it --name broker-go -p 50050:50050 broker-go

server1:
	docker build -t server1-go -f DockerfileServer1 .
	docker run -it --name server1-go -p 50051:50051 server1-go

server2:
	docker build -t server2-go -f DockerfileServer2 .
	docker run -it --name server2-go -p 50052:50052 server2-go

server3:
	docker build -t server3-go -f DockerfileServer3 .
	docker run -it --name server3-go -p 50053:50053 server3-go

jeth:
	docker build -t jeth-go -f DockerfileJeth .
	docker run -it --name jeth-go jeth-go

malkor:
	docker build -t malkor-go -f DockerfileMalkor .
	docker run -it --name malkor-go malkor-go

comandante:
	docker build -t comandante-go -f DockerfileComandante .
	docker run -it --name comandante-go comandante-go