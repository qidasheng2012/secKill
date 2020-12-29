package service

import (
	"seckill/admin/dao"
	"seckill/admin/model"
)

var seckillDao = new(dao.SeckillDao)

type SeckillService struct {
}

func (s *SeckillService) Insert(seckill *model.Seckill) (int64, error) {
	return seckillDao.Insert(seckill)
}

func (s *SeckillService) Delete(seckillId int64) bool {
	return seckillDao.Delete(seckillId)
}

func (s *SeckillService) Update(seckill *model.Seckill) error {
	return seckillDao.Update(seckill)
}

func (s *SeckillService) SelectById(seckillId int64) (*model.Seckill, error) {
	return seckillDao.SelectById(seckillId)
}

func (s SeckillService) SelectAll() ([]*model.Seckill, error) {
	return seckillDao.SelectAll()
}
