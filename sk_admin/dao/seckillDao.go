package dao

import (
	"database/sql"
	"log"
	"time"

	"seckill/admin/common"
	"seckill/admin/model"
)

type ISeckillDao interface {
	Insert(*model.Seckill) (int64, error)
	Delete(int64) bool
	Update(*model.Seckill) error
	SelectById(int64) (*model.Seckill, error)
	SelectAll() ([]*model.Seckill, error)
}

type SeckillDao struct {
}

func (s *SeckillDao) Insert(seckill *model.Seckill) (int64, error) {
	result, err := common.DB.Exec("INSERT INTO seckill(`name`,`number`,`start_time`,`end_time`,`create_time`) VALUE(?,?,?,?,?)",
		seckill.Name, seckill.Number, seckill.StartTime, seckill.EndTime, time.Now())

	if err != nil {
		return 0, err
	}

	id, e := result.LastInsertId()
	if e != nil {
		return 0, e
	}
	return id, nil
}

func (s *SeckillDao) Delete(seckillId int64) bool {
	_, err := common.DB.Exec("DELETE FROM seckill WHERE seckill_id = ?", seckillId)
	if err != nil {
		return false
	}

	return true
}

func (s *SeckillDao) Update(seckill *model.Seckill) error {
	_, err := common.DB.Exec("UPDATE seckill SET name = ? , number = ? , start_time = ? , end_time = ? WHERE seckill_id = ?",
		seckill.Name, seckill.Number, seckill.StartTime, seckill.EndTime, seckill.SeckillId)
	if err != nil {
		return err
	}

	return nil
}

func (s *SeckillDao) SelectById(seckillId int64) (*model.Seckill, error) {
	rows, err := common.DB.Query("SELECT seckill_id,`name`,`number`,start_time,end_time,create_time FROM seckill WHERE seckill_id = ?", seckillId)
	if err != nil {
		return nil, err
	}

	seckills, err := processSeckillRows(rows)
	if err != nil {
		return nil, err
	}

	return seckills[0], nil
}

func (s *SeckillDao) SelectAll() ([]*model.Seckill, error) {
	rows, err := common.DB.Query("SELECT seckill_id,`name`,`number`,start_time,end_time,create_time FROM seckill")
	if err != nil {
		return nil, err
	}

	return processSeckillRows(rows)
}

func processSeckillRows(rows *sql.Rows) ([]*model.Seckill, error) {
	var seckills []*model.Seckill

	for rows.Next() {
		var seckill model.Seckill
		err := rows.Scan(&seckill.SeckillId, &seckill.Name, &seckill.Number, &seckill.StartTime, &seckill.EndTime, &seckill.CreateTime)
		if err != nil {
			log.Printf("product  SelectById error: %v", err)
			continue
		}
		seckills = append(seckills, &seckill)
	}
	rows.Close()

	return seckills, nil
}
