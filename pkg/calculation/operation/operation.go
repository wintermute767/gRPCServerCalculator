package operation

import (
	"gRPCcalculator/pkg/calculation"
	"gRPCcalculator/pkg/calculation/math"
	"os"
)

func GetChooseOperation() calculation.Calculations {

	//Узнаем какой тип мат. операции из переменной окружения
	mathOperation := os.Getenv("MATHOPERATION")

	//Выберем тип мат операции и создадим соответсвующий интерфес
	switch mathOperation {
	case "Add":
		return calculation.Calculations(math.AddOperation{})
	case "Sub":
		return calculation.Calculations(math.SubOperation{})
	case "Mul":
		return calculation.Calculations(math.MulOperation{})
	case "Div":
		return calculation.Calculations(math.DivOperation{})
	}
	//дефолтное значение сложение
	return calculation.Calculations(math.AddOperation{})
}
