package math

import "errors"

// создаем структуры для для каждого типа математической операции
type AddOperation struct{}
type SubOperation struct{}
type MulOperation struct{}
type DivOperation struct{}

func (AddOperation) GetCalculations(pam1 *float32, pam2 *float32) (*float32, error) {
	res := *pam1 + *pam2
	return &res, nil
}

func (SubOperation) GetCalculations(pam1 *float32, pam2 *float32) (*float32, error) {
	res := *pam1 - *pam2
	return &res, nil
}

func (MulOperation) GetCalculations(pam1 *float32, pam2 *float32) (*float32, error) {
	res := *pam1 * *pam2
	return &res, nil
}

func (DivOperation) GetCalculations(pam1 *float32, pam2 *float32) (*float32, error) {
	if *pam2 == 0 {
		return nil, errors.New("you cannot divide by 0. The second parameter should not be 0")
	}
	res := *pam1 / *pam2
	return &res, nil
}
