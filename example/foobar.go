package example

/*
 Create struct foo bar service with dependency foo bar service
 and return it as a pointer
*/
type FooBarService struct {
	*FooService
	*BarService
}

// Create provider functions for each of the dependencies fooBarService with parameters fooService and barService
func NewFooBarService(fooService *FooService, barService *BarService) *FooBarService {
	return &FooBarService{
		FooService: fooService,
		BarService: barService,
	}
}
