# Makefile

myserver:
	docker build -t grpcserver  ./docker/.
	docker run -h grpcserver -d -e PORT=$(port) -P -e MATHOPERATION=$(ope)  grpcserver

ipserv: 
	docker exec grpcserver  hostname -I	
