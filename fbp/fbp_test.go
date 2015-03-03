package fbp

import (
	//	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	component1 := NewComponent()
	component2 := NewComponent()
	port1 := component1.AddPort("out", Out)
	port2 := component2.AddPort("in", In)
	Connect(port1, port2)
}

func TestCounter(t *testing.T) {
	start := NewComponent()
	startPort := start.AddPort("start.out", Out)

	counter := NewCounter()
	counter.AddPort("counter.activate", In)
	counter.AddPort("counter.out", Out)

	counter2 := NewCounter()
	counter2.AddPort("counter2.activate", In)
	counter2.AddPort("counter2.out", Out)

	counter3 := NewCounter()
	counter3.AddPort("counter3.activate", In)

	Connect(startPort, counter.GetPort("counter.activate"))
	Connect(counter.GetPort("counter.out"), counter2.GetPort("counter2.activate"))
	Connect(counter.GetPort("counter.out"), counter3.GetPort("counter3.activate"))

	for i := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		startPort.Push(i)
	}
	//startPort.Push("startdata")

}
