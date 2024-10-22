package Activity

import "errors"

var (
	ErrNilExpense     = errors.New("nil expense")
	ErrInvalidExpense = errors.New("invalid expense")
	ErrScan           = errors.New("scan error")
	ErrNoExpense      = errors.New("no expense")
	ErrCreateFile     = errors.New("error creating file")
	ErrWriteFile      = errors.New("error writing to file")
	ErrOpenFile       = errors.New("error opening file")
)
