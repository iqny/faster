package interface_factory

type Wms interface {
	Execute(req Request) (res Response, err error)
}
