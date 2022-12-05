package models

type Location struct {
	Longitude float64 `bson:"longitude"json:"longitude"` // Kinh do
	Latitude  float64 `bson:"latitude"json:"latitude"`   //vi do
	Address   string  `bson:"address"json:"address"`
	TimeStamp int64   `bson:"timestamp"json:"timestamp"`
	Datestamp int64   `bson:"datestamp"json:"datestamp"`
}
