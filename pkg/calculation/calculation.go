package calculation

type Calculations interface {
	GetCalculations(*float32, *float32) (*float32, error)
}
