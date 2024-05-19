docker-namenode:
	docker build -t namenode -f DockerfileNameNode .
	docker run -it -p 50051:50051 namenode

docker-datanode1:
	docker build -t datanode1 -f DockerfileDataNode1 .
	docker run -it -p 50052:50052 datanode1

docker-datanode2:
	docker build -t datanode2 -f DockerfileDataNode2 .
	docker run -it -p 50053:50053 datanode2

docker-datanode3:
	docker build -t datanode3 -f DockerfileDataNode3 .
	docker run -it -p 50054:50054 datanode3