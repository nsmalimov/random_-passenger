package ordergen

import (
	"bufio"
	"log"
	"math/rand"
	"os"

	"random_passenger_driver/internal/coordinategen"

	"github.com/google/uuid"
)

type OrderGenService struct {
	userNames     []string
	coordinateGen *coordinategen.CoordinateGen
}

type Order struct {
	ID            string
	Username      string
	LatitudeFrom  float64
	LongitudeFrom float64
	LatitudeTo    float64
	LongitudeTo   float64
}

func New(pathToNamesData string, coordinateGen *coordinategen.CoordinateGen) *OrderGenService {
	file, err := os.Open(pathToNamesData)
	if err != nil {
		log.Fatalf("Error when try os.Open, err: %s", err)
	}

	defer func() {
		err = file.Close()

		if err != nil {
			log.Fatalf("Error when try file.Close, err: %s", err)
		}
	}()

	var userNames []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		userNames = append(userNames, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatalf("Error when try scanner.Err, err: %s", err)
	}

	return &OrderGenService{
		userNames:     userNames,
		coordinateGen: coordinateGen,
	}
}

func (o *OrderGenService) GenOrder() (order Order) {
	order.ID = uuid.New().String()

	order.LatitudeTo, order.LongitudeTo = o.coordinateGen.GenCoordinates()
	order.LatitudeFrom, order.LongitudeFrom = o.coordinateGen.GenCoordinates()

	order.Username = o.userNames[rand.Intn(len(o.userNames))]

	return
}
