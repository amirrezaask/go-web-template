package pool

type ResourcePool interface{
	Factory() interface{}
	Acquire(dest interface{})
	Release(resource interface{})
}

type GenericResourcePool struct {
	factory func() interface{}
	pool chan interface{}
}

func (g *GenericResourcePool) Factory() interface{} {
	return g.factory()
}

func (g *GenericResourcePool) Acquire(dest interface{}) {
	select {
	case dest=<-g.pool:
		return
	default:
		dest = g.factory()
		return
	}
}

func (g *GenericResourcePool) Release(resource interface{}) {
	g.pool<-resource
}

