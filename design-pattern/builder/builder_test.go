package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	// Director
	manufacturingComplex := ManufacturingDirector{}

	// Car
	carBuilder := &CarBuilder{}
	// DirectorにCarを渡す
	manufacturingComplex.SetBuilder(carBuilder)
	// 組み立て
	manufacturingComplex.Construct()
	// Carの情報を取得
	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf(
			"Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf(
			"Structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}

	// Bike
	bikeBuilder := BikeBuilder{}
	manufacturingComplex.SetBuilder(&bikeBuilder)
	manufacturingComplex.Construct()

	motorbike := bikeBuilder.GetVehicle()

	if motorbike.Wheels != 2 {
		t.Errorf(
			"Wheels on a car must be 2 and they were %d\n", motorbike.Wheels)
	}

	if motorbike.Structure != "MotorBike" {
		t.Errorf(
			"Structure on a motorbike must be 'MotorBike' and was %s\n", motorbike.Structure)
	}

	if motorbike.Seats != 1 {
		t.Errorf("Seats on a motorbike must be 1 and they were %d\n", motorbike.Seats)
	}
}
