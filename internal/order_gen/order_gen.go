package order_gen

import (
	"bufio"
	"log"
	"math/rand"
	"os"

	"random_passenger_driver/internal/coordinate_gen"

	"github.com/google/uuid"
)

type OrderGenService struct {
	userNames     []string
	coordinateGen *coordinate_gen.CoordinateGen
}

type Order struct {
	id        string
	username  string
	latitude  float64
	longitude float64
}

func New(pathToNamesData string, coordinateGen *coordinate_gen.CoordinateGen) *OrderGenService {
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

func (t *OrderGenService) GenOrder() (order Order) {
	order.id = uuid.New().String()

	order.latitude, order.longitude = t.coordinateGen.GenCoordinates()

	order.username = t.userNames[rand.Intn(len(t.userNames))]

	return
}
