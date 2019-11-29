package main

import (
	"encoding/json"
	"fmt"
	ag "golang-seminary/agency"
)

func main(){

	miAgency := ag.NewAgency()
	miAuto1 := ag.NewCar("ABC123","Renault",2009)
	miAuto2 := ag.NewCar("SEE100","Mitsubishi",2016)
	miAuto3 := ag.NewCar("ABC145","Toyota",2005)
	miAuto4 := ag.NewCar("ABC156","Honda",2010)
	miAgency.AddCar(miAuto1)
	miAgency.AddCar(miAuto2)
	miAgency.AddCar(miAuto3)
	miAgency.AddCar(miAuto4)
	miAutoMarshalled, _ :=json.Marshal(miAuto1)
	fmt.Println(miAutoMarshalled)
	miAgency.PrintAutos()
	miAgency.DeleteAuto("SEE100")
	fmt.Println("---")
	miAgency.PrintAutos()
}