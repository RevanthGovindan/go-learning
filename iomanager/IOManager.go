package iomanager

type IoManger interface {
	ReaLines() ([]string, error)
	WriteResult(data any) error
}
