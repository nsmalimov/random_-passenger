package drivers_processor

import (
	"math/rand"

	"github.com/google/uuid"
)

type Driver struct {
	id       string
	carModel string
}

type DriversProcessor struct {
	automobiles []string
}

func New() *DriversProcessor {
	return &DriversProcessor{
		automobiles: []string{
			"Kia X-Line",
			"Hyundai Creta",
			"Renault Kapture",
			"Kia Rio",
			"Skoda Octavia",
			"Volkswagen Polo",
			"Nissan Qashqai",
			"Skoda Rapid",
			"Hyundai Solaris",
			"Volkswagen Transporter",
			"Citroen Jumpy",
			"Peugeot Expert",
			"Ford Transit",
			"Volkswagen Caravelle",
			"Volkswagen Kombi",
		},
	}
}

func (t *DriversProcessor) Driver() (driver Driver) {
	driver.id = uuid.New().String()
	driver.carModel = t.automobiles[rand.Intn(len(t.automobiles))]

	return
}
