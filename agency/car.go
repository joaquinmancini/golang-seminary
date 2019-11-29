package agency
// Defines Car structure
type Car struct{
	Plate string
	Brand string
	Model int
}
// Creates a new Car instance
func NewCar (Plate, Brand string, Model int)Car{
	return Car{
		Plate:Plate,
		Brand:Brand,
		Model:Model,
	}
}