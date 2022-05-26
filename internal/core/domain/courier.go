package domain

type Courier struct {
	Id            string `json:"id"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Email         string `json:"email"`
	TransportMode string `json:"transport_mode"`
}
