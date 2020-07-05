package db

type Fillable interface {
	Fields() []interface{}
}

type Scannable interface {
	Scan(dest ...interface{}) error
}

func Scan(scannable Scannable, fillable Fillable) error {
	return scannable.Scan(fillable.Fields()...)
}
