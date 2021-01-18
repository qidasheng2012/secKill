package dao

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"seckill/admin/common"
	"seckill/admin/model"
)

type ISeckillDao interface {
	Insert(*model.Seckill) (int64, error)
	Delete(int64) error
	Update(*model.Seckill) error
	SelectById(int64) (*model.Seckill, error)
	SelectAll() ([]*model.Seckill, error)
	selectEtcdById(string) (*model.Seckill, error)
}

type SeckillDao struct {
}

func (s *SeckillDao) Insert(seckill *model.Seckill) (int64, error) {
	conn, err := common.DB.Begin()
	if err != nil {
		log.Println("db begin failed :", err)
		return 0, err
	}

	result, err := common.DB.Exec("INSERT INTO seckill(`name`,`number`,`start_time`,`end_time`,`create_time`) VALUE(?,?,?,?,?)",
		seckill.Name, seckill.Number, seckill.StartTime, seckill.EndTime, time.Now())

	if err != nil {
		conn.Rollback()
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		conn.Rollback()
		return 0, err
	}

	// 同步到etcd
	seckill.SeckillId = id
	seckillIdStr := fmt.Sprintf("%d", id)
	seckillJson, _ := json.Marshal(seckill)
	err = common.EtcdPut(seckillIdStr, string(seckillJson))
	if err != nil {
		conn.Rollback()
		return 0, err
	}

	conn.Commit()
	return id, nil
}

func (s *SeckillDao) Delete(seckillId int64) error {
	conn, err := common.DB.Begin()
	if err != nil {
		return fmt.Errorf("db begin failed :%v", err)
	}

	_, err = common.DB.Exec("DELETE FROM seckill WHERE seckill_id = ?", seckillId)
	if err != nil {
		conn.Rollback()
		return err
	}

	// 同步到etcd
	seckillIdStr := fmt.Sprintf("%d", seckillId)
	err = common.EtcdDelete(seckillIdStr)
	if err != nil {
		conn.Rollback()
		return err
	}

	conn.Commit()
	return nil
}

func (s *SeckillDao) Update(seckill *model.Seckill) error {
	conn, err := common.DB.Begin()
	if err != nil {
		log.Println("db begin failed :", err)
		return err
	}

	_, err = common.DB.Exec("UPDATE seckill SET name = ? , number = ? , start_time = ? , end_time = ? WHERE seckill_id = ?",
		seckill.Name, seckill.Number, seckill.StartTime, seckill.EndTime, seckill.SeckillId)
	if err != nil {
		conn.Rollback()
		return err
	}

	// 同步到etcd
	seckillIdStr := fmt.Sprintf("%d", seckill.SeckillId)
	seckillJson, _ := json.Marshal(seckill)
	err = common.EtcdPut(seckillIdStr, string(seckillJson))
	if err != nil {
		conn.Rollback()
		return err
	}

	conn.Commit()
	return nil
}

func (s *SeckillDao) SelectEtcdById(seckillIdStr string) (*model.Seckill, error) {
	seckillStr, err := common.EtcdGet(seckillIdStr)
	if err != nil {
		return nil, err
	}

	if seckillStr == "" {
		return nil, nil
	}

	seckill := model.Seckill{}
	err = json.Unmarshal([]byte(seckillStr), &seckill)
	if err != nil {
		log.Panicf("json error %v", err)
	}

	return &seckill, nil
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
