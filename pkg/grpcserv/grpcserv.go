package grpcserv

import (
	"context"
	"gRPCcalculator/pkg/api"
	"gRPCcalculator/pkg/grpcserv/serveranswer"
	"gRPCcalculator/pkg/logging"
	"github.com/sirupsen/logrus"
)

// Задаем структуру сервера на основе автосгенерированного кода указанного ./pkg/api
// указываем пустую структуру будующих функций
type GRPCServer struct {
	api.UnimplementedDatabusServiceServer
}

// Указываем основную функцию gRPC
func (GRPCServer) Send(_ context.Context, req *api.SendRequest) (*api.SendResponse, error) {

	logging.Ent.Info("server began calculations")
	result, err := serveranswer.GetAnswerRequest(&req.Prm1, &req.Prm2)
	if err != nil {
		logging.Ent.WithFields(logrus.Fields{"error": err}).Error("error get answer request")
	}

	logging.Ent.Info("server send response")

	return &api.SendResponse{Result: *result}, err
}
