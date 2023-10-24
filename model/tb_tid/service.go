package tb_tid

type Service interface {
	Create(tbTidInput TbTidCreateInput) (TbTid, error)
	GetOneByID(input GetOneByIDInput) (TbTid, error)
}

type service struct {
	tbTidRepository Repository
}

func NewService(tbTidRepository Repository) *service {
	return &service{tbTidRepository}
}

func (s *service) Create(tbTidInput TbTidCreateInput) (TbTid, error) {
	tbTid := TbTid{
		Tid:        tbTidInput.Tid,
		IpAddress:  tbTidInput.IpAddress,
		SnMiniPc:   tbTidInput.SnMiniPc,
		LocationId: tbTidInput.LocationId,
	}

	newTbTid, err := s.tbTidRepository.Create(tbTid)
	if err != nil {
		return newTbTid, err
	}

	return newTbTid, nil
}

func (s *service) GetOneByID(input GetOneByIDInput) (TbTid, error) {
	tbTid, err := s.tbTidRepository.GetOneByID(input.ID)
	if err != nil {
		return tbTid, err
	}
	if tbTid.ID == 0 {
		return tbTid, nil
	}

	return tbTid, nil
}
