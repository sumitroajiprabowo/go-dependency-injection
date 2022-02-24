//go:build wireinject
// +build wireinject

package example

import (
	"io"
	"os"

	"github.com/google/wire"
)

/*
Create injector functions for each of the dependencies
and return a pointer to the struct
*/
func InizializeSimpleService() *SimpleService {
	wire.Build(NewSimpleService, NewSimpleRepository)
	return &SimpleService{}
}

/*
Create injector functions for each of the dependencies
with parameters isError is boolean and return a pointer to the struct and error
*/
func InizializeSimpleServiceWithError(isError bool) (*SimpleServiceWithError, error) {
	wire.Build(NewSimpleServiceWithError, NewSimpleRepositoryWithError)
	return &SimpleServiceWithError{}, nil
}

/*
Create injector InitializeDatabaseRepository function
with wire build NewDatabaseRepository, NewDatabasePostgresSQL, NewDatabaseMySQL and return a nil pointer
*/
func InitializeDatabaseRepository() *DatabaseRepository {
	wire.Build(NewDatabaseRepository, NewDatabasePostgresSQL, NewDatabaseMySQL)
	return nil
}

/*
Grouping providers for each of the dependencies
Example of grouping providers for each of the dependencies
*/

// Create provider set with name foo set with parameters foo repository and foo service
var FooSet = wire.NewSet(NewFooRepository, NewFooService)

// Create provider set with name bar set with parameters bar repository and bar service
var BarSet = wire.NewSet(NewBarRepository, NewBarService)

// Create injector InitializeFooBarService function with wire build FooSet BarSet and NewFooBarService and return a nil pointer
func InitializeFooBarService() *FooBarService {
	wire.Build(FooSet, BarSet, NewFooBarService)
	return nil
}

// Create injector InizializedHelloService function with return HelloService and wire build NewHelloService and NewHelloImpl

// Error
// func InizializedHelloService() *HelloService {
// 	wire.Build(NewHelloService, NewHelloImpl)
// 	return nil
// }

/*
Implementation Create Provider Set for bind interface
Create provider set helloSet
*/
var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

// Create injector InizializedHelloService function with return HelloService and wire build helloSet and NewHelloService
func InizializedHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}

/*
Implementation Structs Provider Set for bind interface
*/

// Create injector InizializedFooBar function with return FooBar and wire build FooSet BarSet and NewFooBar function with structs Foo and Bar
func InizializedFooBar() *FooBar {
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "Foo", "Bar"))
	return nil
}

// Another example using wire.Struct
// func InizializedFooBar() *FooBar {
// 	wire.Build(NewBar, wire.Struct(new(FooBar), "Bar"))
// 	return nil
// }

// Another example using wire.Struct
// func InizializedFooBar() *FooBar {
// 	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "*"))
// 	return nil
// }

/*
Implementation Create Provider Set for bind value
*/

var fooValue = &Foo{} // Value
var barValue = &Bar{} // Value

//Create injector InizializedFooBarUsingBindValue function with return FooBar and wire build with value fooValue and barValue and NewFooBar function with structs Foo and Bar
func InizializedFooBarUsingBindValue() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"))
	return nil
}

/*
Injector Interface Value
*/

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

/*
Struct Field Provider Set
*/

// Create injector InizializedCustomer function with return Person and wire build with struct field of Customer and Person
func InizializedCustomer() *Person {
	wire.Build(NewCustomer, wire.FieldsOf(new(*Customer), "Person"))
	return nil
}

/*
Cleanup function
*/

// Create injector InizializedCleanup function with return Cleanup and wire build with Cleanup function
func InizializedCleanup(name string) (*FileSystem, func()) {
	wire.Build(NewFileSystem, NewFile)
	return nil, nil
}
