package entities

import "github.com/google/uuid"

type Product struct {
	Uuid    uuid.UUID      `json:"uuid"`
	Sapiens ProductSapiens `bson:"sapiens"`
	Gatec   ProductGatec   `bson:"gatec"`
	VetoRH  ProductVetoRH  `bson:"vetorh"`
	Hub     ProductHub     `bson:"hub"`
}

type ProductSapiens struct {
	Id          int    `bson:"id"`
	Description string `bson:"description"`
}
type ProductGatec struct {
	Id          int    `bson:"id"`
	Description string `bson:"description"`
}
type ProductVetoRH struct {
	Id          int    `bson:"id"`
	Description string `bson:"description"`
}
type ProductHub struct {
	Id          int    `bson:"id"`
	Description string `bson:"description"`
}
