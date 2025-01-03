package entities

import "github.com/google/uuid"

type Customer struct {
	Uuid       uuid.UUID       `bson:"uuid"`
	Identifier string          `bson:"identifier"`
	Name       string          `bson:"name"`
	Sapiens    CustomerSapiens `bson:"sapiens"`
	Gatec      CustomerGatec   `bson:"gatec"`
	VetoRH     CustomerVetoRH  `bson:"vetorh"`
	Hub        CustomerHub     `bson:"hub"`
}

type CustomerSapiens struct {
	Id   int    `bson:"id"`
	Name string `bson:"name"`
}

type CustomerGatec struct {
	Id   int    `bson:"id"`
	Name string `bson:"name"`
}

type CustomerVetoRH struct {
	Id   int    `bson:"id"`
	Name string `bson:"name"`
}

type CustomerHub struct {
	Id   int    `bson:"id"`
	Name string `bson:"name"`
}
