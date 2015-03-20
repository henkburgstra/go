package application

import ()

type IService interface {
	Name() string
	Start(args ...interface)
}
