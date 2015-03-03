package fbp

type Pipe struct {
	from    *Port
	to      *Port
	channel chan interface{}
}

func NewPipe() *Pipe {
	pipe := new(Pipe)
	pipe.channel = make(chan interface{}, 1)
	return pipe
}

func (p *Pipe) Connect(from *Port, to *Port) {
	p.ConnectFrom(from)
	p.ConnectTo(to)
}

func (p *Pipe) ConnectFrom(from *Port) {
	p.from = from
	from.AddPipe(p)
}

func (p *Pipe) ConnectTo(to *Port) {
	p.to = to
	to.AddPipe(p)
	go to.Listen(p.channel)
}

func (p *Pipe) From() *Port {
	return p.from
}

func (p *Pipe) To() *Port {
	return p.to
}

func (p *Pipe) Push(data interface{}) {
	p.channel <- data
}
