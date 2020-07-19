package pool

type Pool interface{
	Factory() interface{}
	Acquire(dest interface{})
	Release(resource interface{})
}


type GeneralPool struct {
	factory func() interface{}
	pool chan interface{}
}

func (g *GeneralPool) Factory() interface{} {
	return g.factory()
}

func (g *GeneralPool) Acquire(dest interface{}) {
	select {
	case dest=<-g.pool:
		return
	default:
		dest = g.factory()
		return
	}
}

func (g *GeneralPool) Release(resource interface{}) {
	g.pool<-resource
}

