package representation

type PhotoRepresentation struct {
	ID int `json:"id"`
	DateTimeOriginal string `json:"dateTimeOriginal"`
	Manufacturer string	`json:"manufacturer"`
	Model	string	`json:"model"`
}