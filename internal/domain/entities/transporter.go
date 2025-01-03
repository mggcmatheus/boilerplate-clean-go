package entities

import "github.com/google/uuid"

type Transporter struct {
	Uuid    uuid.UUID          `json:"uuid"`
	Sapiens TransporterSapiens `bson:"sapiens"`
	Gatec   TransporterGatec   `bson:"gatec"`
	VetoRH  TransporterVetoRH  `bson:"vetorh"`
	Hub     TransporterHub     `bson:"hub"`
}

type TransporterSapiens struct {
	Id          int    `bson:"id"`
	Description string `bson:"description"`
}
type TransporterGatec struct {
	Id          int    `bson:"id"`
	Description string `bson:"description"`
}
type TransporterVetoRH struct {
	Id          int    `bson:"id"`
	Description string `bson:"description"`
}
type TransporterHub struct {
	Id          int    `bson:"id"`
	Description string `bson:"description"`
}
