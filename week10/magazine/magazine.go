package magazine

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
	//HomeAddress Address
}
type Employee struct {
	Name string
	Rate float64
	Address
	//HomeAddress Address
}
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
