package domain

type Customer struct {
	Id          string `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	PaymentMode string `json:"payment_mode"`
	Location    string `json:"location"`
}
