package express

type Address struct {
	Original   string `json:"original"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Detail     string `json:"detail"`
	PostalCode string `json:"postalCode"`
}
