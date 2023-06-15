package service

func (s *service) UpdateByID(id int , brand string) error {
	err := s.store.UpdateByID(id,brand)
	if err!= nil {
        return err
    }
	return err
}