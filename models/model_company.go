package models

type Company struct {
	Name        *string `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Amount      *int    `json:"amount" bson:"amount"`
	Registered  *bool   `json:"registered" bson:"registered"`
	Type        *string `json:"type" bson:"type"`
}
