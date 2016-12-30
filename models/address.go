package models

type address struct {
	ID             string `json:"id"`
	BuildingNumber string `json:"building_number"`
	BuildingName   string `json:"building_name"`
	StreetNumber   string `json:"street_number"`
	StreetName     string `json:"street_name"`
	PostCode       string `json:"post_code"`
	Latitude       string `json:"latitude"`
	Longitude      string `json:"longitude"`
	GroupOwner     string `json:"groupOwner"`
}