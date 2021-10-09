package bnnd

type BNNDService interface {
	GetAllBananaDucks() ([]BananaDuck, error)
	CreateBananaDuck(d BananaDuck) error
}

type BNNDManager struct {
	bnndRepo BNNDRepo
}

// NewBNNDManager creates a new BNNDService for managing ducks
func NewBNNDManager(bnndRepo BNNDRepo) BNNDService {
	return &BNNDManager{
		bnndRepo: bnndRepo,
	}
}

// GetAllBananaDucks returns all ducks found of type []BananaDuck
func (s *BNNDManager) GetAllBananaDucks() ([]BananaDuck, error) {
	ducks, err := s.bnndRepo.FindAll()
	if err != nil {
		return []BananaDuck{}, err
	}
	return ducks, nil
}

// CreateBananaDuck returns the created duck of type BananaDuck
func (s *BNNDManager) CreateBananaDuck(d BananaDuck) error {
	err := s.bnndRepo.Insert(d)
	if err != nil {
		return err
	}
	return nil
}
