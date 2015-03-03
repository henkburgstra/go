package fbp

type IComponent interface {
	AddPort(string, PortType) *Port
	GetPorts() []*Port
	GetPort(string) *Port
	HandleData(*Port)
}

type Component struct {
	ports map[string]*Port
	Owner IComponent // referentie naar de aanroepende instantie
}

func NewComponent() *Component {
	component := new(Component)
	component.ports = make(map[string]*Port)
	return component
}

func (c *Component) AddPort(name string, portType PortType) *Port {
	port := NewPort(name, portType, c)
	c.ports[name] = port
	return port
}

func (c *Component) GetPorts() []*Port {
	ports := make([]*Port, 0, 0)
	for _, port := range c.ports {
		ports = append(ports, port)
	}
	return ports
}

func (c *Component) GetPort(name string) *Port {
	return c.ports[name]
}

func (c *Component) HandleData(port *Port) {
	for _, port := range c.Owner.GetPorts() {
		if port.PortType() == Out {
			port.Push(port.Data())
		}
	}
}
