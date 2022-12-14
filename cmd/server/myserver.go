package main

import (
	"fmt"
	"gRPCcalculator/pkg/api"
	"gRPCcalculator/pkg/grpcserv"
	"gRPCcalculator/pkg/logging"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func main() {
	logging.Ent.Info("start program")
	//Задаем параметры из окружения для последущего использования в docker контейнере
	port := os.Getenv("PORT")
	if port == "" {
		port = "10500"
	}

	//Задаем службу для прослушивания порта
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		logging.Ent.WithFields(logrus.Fields{"error": err}).Error("error listening server")
		log.Fatal(err)
	}
	logging.Ent.WithFields(logrus.Fields{"port": port}).Debug("set server listening port")
	//Создаем новый gRPC сервер
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	//Регистрируем собственную конфигурацию сервера и заем ее прослушивающей службе
	api.RegisterDatabusServiceServer(grpcServer, &grpcserv.GRPCServer{})
	if err := grpcServer.Serve(listener); err != nil {
		logging.Ent.WithFields(logrus.Fields{"error": err}).Error("server registration error")
		log.Fatal(err)
	}
}
