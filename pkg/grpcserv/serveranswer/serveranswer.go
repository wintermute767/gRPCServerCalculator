package serveranswer

import (
	"gRPCcalculator/pkg/calculation/operation"
	"gRPCcalculator/pkg/logging"
	"github.com/sirupsen/logrus"
	"reflect"
)

func GetAnswerRequest(pam1 *float32, pam2 *float32) (*float32, error) {

	//выберем тип операции
	operation := operation.GetChooseOperation()

	//Выполним вычисления
	logging.Ent.WithFields(logrus.Fields{"mathOperation": reflect.TypeOf(operation), "pam1": *pam1, "pam2": *pam2}).Info("parameters calculations")
	answ, err := operation.GetCalculations(pam1, pam2)
	if err != nil {
		logging.Ent.WithFields(logrus.Fields{"error": err}).Error("error get answer request")
	}
	logging.Ent.WithFields(logrus.Fields{"answer": *answ}).Info("result calculations")

	return answ, err
}
