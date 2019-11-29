package agency

import (
	"errors"
	"fmt"
	"strconv"
)
// Defines Agency structure
type Agency struct{
	Cars[]Car
}
// Creates a new Agency instance
func NewAgency()Agency{
	var aux[]Car
	return Agency{
		Cars:aux,
	}
}
// Adds new Car to slice
func (a *Agency) AddCar(a1 Car){
	a.Cars=append(a.Cars, a1)
}
// Searches a Car by a given Plate number
func (a *Agency) SearchCar(p string)(*Car, error){
	for _, car := range a.Cars {
		if car.Plate==p{
			return &car, nil
		}
	}
	return nil, errors.New("car not found")
}
// Searches a Car by a given Plate number and returns its related index in the slice (used by Delete function)
func (a *Agency) SearchCarIndex(p string) int{
	aux:=a.Cars
	for i := 0; i<len(aux); i++ {
		if aux[i].Plate==p{
			return i
		}
	}
	return -1
}
// Prints all Cars in slice
func (a *Agency) PrintAutos(){
	for i := 0; i<len(a.Cars); i++ {
		fmt.Println("Plate is "+a.Cars[i].Plate+", Brand is "+a.Cars[i].Brand+", Model is "+strconv.Itoa(a.Cars[i].Model))
	}
}
// Updates all Car fields
func (a *Agency) UpdateCarAll(p, ap, am string, am2 int) error{
	if a.UpdateCarPlate(p, ap)!=nil && a.UpdateCarBrand(p, am)!=nil && a.UpdateCarModel(p, am2)!=nil{
		return nil
	}
	return errors.New("car not found")
}
// Updates Car Plate only
func (a *Agency) UpdateCarPlate(p, ap string) error{
	for _, car := range a.Cars {
		if car.Plate==p{
			car.Plate=ap
			return nil
		}
	}
	return errors.New("car not found")
}
// Updates Car Brand only
func (a *Agency) UpdateCarBrand(p, m string) error{
	for _, car := range a.Cars {
		if car.Plate==p{
			car.Brand=m
			return nil
		}
	}
	return errors.New("car not found")
}
// Updates Car Model only
func (a *Agency) UpdateCarModel(p string, m int) error{
	for _, car := range a.Cars {
		if car.Plate==p{
			car.Model=m
			return nil
		}
	}
	return errors.New("car not found")
}
// Deletes a Car by a given Plate number (unarranged deletion)
func (a *Agency) DeleteAuto(p string) error{
	aux:=a.Cars
	i:=a.SearchCarIndex(p)
	if i>=0 {
		aux[i] = aux[len(aux)-1]
		a.Cars = aux[:len(aux)-1]
		return nil
	}
	return errors.New("car not found")
}