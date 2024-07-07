package car

import (
	"log"
	"proto-course/protogen/car"

	"github.com/google/uuid"
)

func ValidateCar() {
	car := car.Car{
		CarId:           uuid.New().String(),
		Brand:           "BMW",
		Model:           "X5",
		Price:           100000,
		ManufactureYear: 2021,
	}

	err := car.ValidateAll()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Car is valid")
	log.Println(&car)
}
