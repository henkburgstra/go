package fbp

import (
	"fmt"
)

type Counter struct {
	*Component
}

func NewCounter() *Counter {
	counter := new(Counter)
	counter.Component = NewComponent()
	counter.Owner = counter
	return counter
}

func (c *Counter) HandleData(port *Port) {
	fmt.Println("Data uit Counter", port.name)
	// Hier komt de business logica

	// Onderstaande regel zorgt ervoor dat de data naar uitgaande poorten gepusht wordt.
	c.Component.HandleData(port)
}
