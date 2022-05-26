package domain

type Restaurant struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Phone    string `json:"Phone"`
	Image    string `json:"image"`
}
