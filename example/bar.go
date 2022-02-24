package example

// Create struct bar repository
type BarRepository struct {
}

// Create provider functions for each of the dependencies barRepository
func NewBarRepository() *BarRepository {
	return &BarRepository{}
}

// Create struct bar service with dependency barRepository
type BarService struct {
	*BarRepository
}

// Create provider functions for each of the dependencies barService with parameters barRepository
func NewBarService(barRepository *BarRepository) *BarService {
	return &BarService{
		BarRepository: barRepository,
	}
}
