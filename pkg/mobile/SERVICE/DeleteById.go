package service


func (s *service) DeleteMobile(id int) error {
	err := s.store.DeleteByID(id)
	if err!= nil {
        return err
    }
	return err
}