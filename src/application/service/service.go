package service

import (
	"fmt"
	"guide_go/src/internal/launch"
	"sync"
)

type baseService struct {
	*launch.Launcher
}

type service struct {
	*baseService
	UserService *UserService
}

type ServicePool struct {
	pool *sync.Pool
}

func NewServicePool(launch *launch.Launcher) *ServicePool {
	return &ServicePool{
		pool: &sync.Pool{
			New: func() any {
				return newService(launch)
			},
		},
	}
}

func newService(launch *launch.Launcher) *service {
	base := &baseService{launch}
	return &service{
		base,
		&UserService{
			base,
			launch.Mysql.UserRepositoryImpl,
		},
	}
}

func (p *ServicePool) Get() *service {
	fmt.Println("p", p)
	s := p.pool.Get().(*service)
	fmt.Println("???????", s)
	return s
}

func (p *ServicePool) Return(service *service) {
	p.pool.Put(service)
}
