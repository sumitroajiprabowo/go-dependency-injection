package example

//Create struct foo repository
type FooRepository struct {
}

// Create provider functions for each of the dependencies fooRepository
func NewFooRepository() *FooRepository {
	return &FooRepository{}
}

// Create struct foo service with dependency fooRepository
type FooService struct {
	*FooRepository
}

// Create provider functions for each of the dependencies fooService with parameters fooRepository
func NewFooService(fooRepository *FooRepository) *FooService {
	return &FooService{
		FooRepository: fooRepository,
	}
}
