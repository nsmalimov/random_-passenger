package order_gen_service

import (
	"math"
	"math/rand"
)

type OrderGenService struct {
	pathToNamesData string
}

type Order struct {
	username  string
	latitude  float64
	longitude float64
}

func New(pathToNamesData string) *OrderGenService {
	return &OrderGenService{
		pathToNamesData: pathToNamesData,
	}
}

func (t *OrderGenService) randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func (t *OrderGenService) Round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}

func (t *OrderGenService) GenCoordinates(x0, y0, radius float64) (float64, float64) {
	var radiusInDegrees float64
	radiusInDegrees = float64(25000 / 111000.0)
	u := t.randFloat(0, 1)
	v := t.randFloat(0, 1)

	w := radiusInDegrees * math.Sqrt(u)
	tN := 2 * math.Pi * v
	x := w * math.Cos(tN)
	y := w * math.Sin(tN)

	new_x := x / math.Cos(math.Pi*y0/180.0)

	// 55.742245, 37.709603
	foundLongitude := t.Round(new_x+x0, 6)
	foundLatitude := t.Round(y+y0, 6)

	return foundLongitude, foundLatitude
}

func (t *OrderGenService) GenOrder() (order Order) {
	order.latitude = 0
	order.longitude = 0
	order.username = ""

	return
}
