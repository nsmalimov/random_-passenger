package drivers_processor

import (
	"math/rand"

	"random_passenger_driver/internal/coordinate_gen"

	"github.com/google/uuid"
)

type Driver struct {
	id        string
	carModel  string
	latitude  float64
	longitude float64
}

type DriversProcessor struct {
	automobiles   []string
	coordinateGen *coordinate_gen.CoordinateGen
}

func New(coordinateGen *coordinate_gen.CoordinateGen) *DriversProcessor {
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
		coordinateGen: coordinateGen,
	}
}

func (t *DriversProcessor) Driver() (driver Driver) {
	driver.id = uuid.New().String()
	driver.carModel = t.automobiles[rand.Intn(len(t.automobiles))]

	driver.latitude, driver.longitude = t.coordinateGen.GenCoordinates()

	return
}
