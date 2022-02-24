package example

import "errors"

/*
Create example structs and interfaces to be used in the tests using dependency injection with Google Wire
*/

// Create a new struct Simple Service instance and return it as a pointer
type SimpleRepository struct {
}

//Create a new provider function for each of the dependencies
func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

/*
Create a new struct Simple Service instance and return it as a pointer simple repository
*/
type SimpleService struct {
	*SimpleRepository
}

// Create a new Simple Service instance and return it as a pointer
func NewSimpleService(repository *SimpleRepository) *SimpleService {
	return &SimpleService{
		SimpleRepository: repository,
	}
}

/*
Create example structs and interfaces with error to be used in the tests using dependency injection with Google Wire
*/

// Create a new Simple Service instance and return it as a pointer with error
type SimpleRepositoryWithError struct {
	Error bool
}

/*
Create provider functions for each of the dependencies with parameters isError is boolean and return a pointer to the struct and error
*/
func NewSimpleRepositoryWithError(isError bool) *SimpleRepositoryWithError {
	return &SimpleRepositoryWithError{
		Error: isError,
	}
}

/*
Create a new struct Simple Service With Error instance and return it as a pointer New Simple Service With Error
*/
type SimpleServiceWithError struct {
	*SimpleRepositoryWithError
}

/*
Create a new Simple Service with Error instance and return it as a pointer with error
*/
func NewSimpleServiceWithError(repository *SimpleRepositoryWithError) (*SimpleServiceWithError, error) {
	if repository.Error {
		return nil, errors.New("Error creating Simple Service")
	}
	return &SimpleServiceWithError{repository}, nil
}
