package application

import ()

type IService interface {
	String() string
}

type IConfigService interface {
	IService
	Config() struct{}
	SetConfig(struct{})
}
