.PHONY: broker server1 server2 server3 jeth malkor comandante

broker:
	sudo docker build -t broker-go -f DockerfileBroker .
	sudo docker run -it --name broker-go -p 50050:50050 broker-go

server1:
	sudo docker build -t server1-go -f DockerfileServer1 .
	sudo docker run -it --name server1-go -p 50051:50051 server1-go

server2:
	sudo docker build -t server2-go -f DockerfileServer2 .
	sudo docker run -it --name server2-go -p 50052:50052 server2-go

server3:
	sudo docker build -t server3-go -f DockerfileServer3 .
	sudo docker run -it --name server3-go -p 50053:50053 server3-go

jeth:
	sudo docker build -t jeth-go -f DockerfileJeth .
	sudo docker run -it --name jeth-go jeth-go

malkor:
	sudo docker build -t malkor-go -f DockerfileMalkor .
	sudo docker run -it --name malkor-go malkor-go

comandante:
	sudo docker build -t comandante-go -f DockerfileComandante .
	sudo docker run -it --name comandante-go comandante-go