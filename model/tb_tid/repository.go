package tb_tid

import "gorm.io/gorm"

type Repository interface {
	Create(tbTid TbTid) (TbTid, error)
	GetOneByID(id int) (TbTid, error)
	GetOneByTid(tid string) (TbTid, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(tbTid TbTid) (TbTid, error) {
	err := r.db.Create(&tbTid).Error
	if err != nil {
		return TbTid{}, err
	}

	return tbTid, nil
}

func (r *repository) GetOneByID(id int) (TbTid, error) {
	var tbTid TbTid

	err := r.db.Where("id = ?", id).First(&tbTid).Error
	if err != nil {
		return tbTid, err
	}

	return tbTid, nil
}

func (r *repository) GetOneByTid(tid string) (TbTid, error) {
	var tbTid TbTid

	err := r.db.Where("tid = ?", tid).First(&tbTid).Error
	if err != nil {
		return tbTid, err
	}

	return tbTid, nil
}
