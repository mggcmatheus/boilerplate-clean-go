package entities

import "github.com/google/uuid"

type Provider struct {
	Uuid       uuid.UUID       `bson:"uuid"`
	Identifier string          `bson:"identifier"`
	Sapiens    ProviderSapiens `bson:"sapiens"`
	Gatec      ProviderGatec   `bson:"gatec"`
	VetoRH     ProviderVetoRH  `bson:"vetorh"`
	Hub        ProviderHub     `bson:"hub"`
}

type ProviderSapiens struct {
	Id   int    `bson:"id"`
	Name string `bson:"name"`
}

type ProviderGatec struct {
	Id   int    `bson:"id"`
	Name string `bson:"name"`
}

type ProviderVetoRH struct {
	Id   int    `bson:"id"`
	Name string `bson:"name"`
}

type ProviderHub struct {
	Id   int    `bson:"id"`
	Name string `bson:"name"`
}
