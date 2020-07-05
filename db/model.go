package db

//Fillable is the interface that should be implementd by everyone that wants to use Scan.
type Fillable interface {
	Fields() []interface{}
}

//Scannable is implemented by sql.Rows and sql.Row so we can use them in Scan function.
type Scannable interface {
	Scan(dest ...interface{}) error
}

//Scan given scannable into the fillable pointers.
func Scan(scannable Scannable, fillable Fillable) error {
	return scannable.Scan(fillable.Fields()...)
}
