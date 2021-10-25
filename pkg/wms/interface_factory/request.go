package interface_factory


type Request interface {
	ToXML() string
	Check() (Response,error)
	GetMethod() string
}
