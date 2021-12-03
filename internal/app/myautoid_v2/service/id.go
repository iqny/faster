package service

func (s *Service) GetID(key int) (int64, error) {
	return s.db.GetAutoId(key)
}
