package bnnd

type BNNDService interface {
	GetAllBananaDucks() ([]BananaDuck, error)
	CreateBananaDuck(d BananaDuck) error
}

type BNNDManager struct {
	bnndRepo BNNDRepo
}

func NewBNNDManager(bnndRepo BNNDRepo) BNNDService {
	return &BNNDManager{
		bnndRepo: bnndRepo,
	}
}

func (s *BNNDManager) GetAllBananaDucks() ([]BananaDuck, error) {
	ducks, err := s.bnndRepo.FindAll()
	if err != nil {
		return []BananaDuck{}, err
	}
	return ducks, nil
}

func (s *BNNDManager) CreateBananaDuck(d BananaDuck) error {
	err := s.bnndRepo.Insert(d)
	if err != nil {
		return err
	}
	return nil
}
