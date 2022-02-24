package example

// Create interface SayHello with func Hello with parameter name and string return value string
type SayHello interface {
	Hello(name string) string
}

// Create struct hello service with dependency sayHello
type HelloService struct {
	SayHello
}

// Create struct SayHelloImpl
type SayHelloImpl struct {
}

// Create implementation method SayHelloImpl Hello with parameter name and string return value string
func (s *SayHelloImpl) Hello(name string) string {
	return "Hello " + name
}

// Create new provider with name NewSayHelloImpl and return a pointer to the struct
func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

// Create new provider NewHelloService with parameters sayHello and return a pointer to the struct HelloService with dependency sayHello
func NewHelloService(sayHello SayHello) *HelloService {
	return &HelloService{
		SayHello: sayHello,
	}
}
