package service

type HelloService interface {
	GetMessage() string
}
type helloService struct{}

func NewHelloService() HelloService {
	s := &helloService{}
	return s
}

func (s *helloService) GetMessage() string {
	return "hello"
}
