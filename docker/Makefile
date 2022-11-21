# Makefile

myserver:
	docker build -t grpcserver  .
	docker run --name grpcserver -d -e PORT=$(port) -p $(port):$(port) -e MATHOPERATION=$(oper)  grpcserver

