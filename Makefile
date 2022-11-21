# Makefile

myserver:
	docker build -t grpcserver  ./docker/.
	docker run --name grpcserver -d -e PORT=$(port) -p $(port):$(port) -e MATHOPERATION=$(oper)  grpcserver

