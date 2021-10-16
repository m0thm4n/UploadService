package request_models

type CreatePhotoModel struct {
	DateTimeOriginal string `json:"dateTimeOriginal"`
	Manufacturer string	`json:"manufacturer"`
	Model string	`json:"model"`
}