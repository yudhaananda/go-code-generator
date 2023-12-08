package repositories

import (
	"generator/repositories/generate"
)

type Repositories struct {
	Generate generate.Interface
}

type Param struct {
}

func Init(param Param) *Repositories {
	return &Repositories{
		Generate: generate.Init(),
	}
}
