package service

func (s *Service) GetID(key int16) (int64, error) {
	return s.db.GetAutoId(key)
}
