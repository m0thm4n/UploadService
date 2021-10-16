package photo

type Photo struct {
	ID int `bson:"_id"`
	// Location string `bson:"Location"`
	DateTimeOriginal string `bson:"DateTimeOriginal"`
	Manufacturer	string	`bson:"Manufacturer"`
	Model	string	`bson:"Model"`
	PhotoLocation string `bson:"PhotoLocation"`
}