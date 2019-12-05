package drivers_processor

import (
	"math/rand"

	"random_passenger_driver/internal/coordinate_gen"

	"github.com/google/uuid"
)

type Driver struct {
	ID        string
	CarModel  string
	Latitude  float64
	Longitude float64
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

func (d *DriversProcessor) Driver() (driver Driver) {
	driver.ID = uuid.New().String()
	driver.CarModel = d.automobiles[rand.Intn(len(d.automobiles))]

	driver.Latitude, driver.Longitude = d.coordinateGen.GenCoordinates()

	return
}
