package fbp

import (
	"fmt"
)

type PortStatus int

const (
	Open PortStatus = iota
	Closed
)

type PortType int

const (
	In PortType = iota
	Out
)

type Port struct {
	data     interface{}
	name     string
	portType PortType
	parent   IComponent
	pipes    []*Pipe
	status   PortStatus
}

func NewPort(name string, portType PortType, parent IComponent) *Port {
	port := new(Port)
	port.name = name
	port.portType = portType
	port.parent = parent
	port.pipes = make([]*Pipe, 0, 0)
	port.status = Open
	return port
}

func (p *Port) PortType() PortType {
	return p.portType
}

func (p *Port) GetData() interface{} {
	return p.data
}

func (p *Port) AddPipe(pipe *Pipe) {
	p.pipes = append(p.pipes, pipe)
}

func (p *Port) Listen(to chan interface{}) {
	fmt.Println("Listen", p.name)
	for {
		ip := <-to
		fmt.Println("Receive", p.name)
		p.SetData(ip)
		p.parent.(*Component).Owner.HandleData(p)
	}
}

func (p *Port) Status() PortStatus {
	return Open
}

func (p *Port) SetStatus(status PortStatus) {
	p.status = status
}

func (p *Port) Data() interface{} {
	return p.data
}

func (p *Port) SetData(data interface{}) {
	p.data = data
}

func (p *Port) Push(data interface{}) {
	for i, pipe := range p.pipes {
		fmt.Println("Push", p.name, "pipe", i)
		pipe.Push(data)
	}
}
