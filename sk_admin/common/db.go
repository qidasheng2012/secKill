package common

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// 连接数据库配置
const (
	username   = "root"
	password   = "123456"
	ip         = "localhost"
	port       = "3306"
	dbName     = "seckill"
	driverName = "mysql"
)

// DB 数据库连接池
var DB *sql.DB

func ConnectMySQL() error {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=uft8"
	//注意：要想解析time.Time类型，必须要设置parseTime=True
	path := strings.Join([]string{username, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8&parseTime=True&loc=Local"}, "")
	//打开数据库，前者是驱动名，所以要导入:_"github.com/go-sql-driver/mysql"
	DB, _ = sql.Open(driverName, path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		return err
	}

	log.Print("mysql database connect success")
	return nil
}

func CreateTable() {
	productTable := "CREATE TABLE IF NOT EXISTS `product` (" +
		"`product_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '商品Id'," +
		"`product_name` varchar(50) NOT NULL DEFAULT '' COMMENT '商品名称'," +
		"`total` int(5) unsigned NOT NULL DEFAULT '0' COMMENT '商品数量'," +
		"`status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '商品状态: 0:在售，1：下架'," +
		"PRIMARY KEY (`product_id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品数据表';"

	_, err := DB.Exec(productTable)
	if err != nil {
		log.Printf("create productTable error: %v", err)
	}
	log.Println("create productTable success")

	seckillTable := "CREATE TABLE IF NOT EXISTS seckill(" +
		"`seckill_id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '商品库存id'," +
		"`name` VARCHAR(120) NOT NULL COMMENT '商品名称'," +
		"`number` INT NOT NULL COMMENT '库存数量'," +
		"`start_time` TIMESTAMP NOT NULL COMMENT '秒杀开始时间'," +
		"`end_time` TIMESTAMP NOT NULL COMMENT '秒杀结束时间'," +
		"`create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'," +
		"PRIMARY KEY (seckill_id)," +
		"KEY idx_start_time (start_time)," +
		"KEY idx_end_time (end_time)," +
		"KEY idx_create_time (create_time)" +
		")ENGINE=InnoDB DEFAULT CHARSET= utf8mb4  COMMENT='秒杀库存表';"

	_, err = DB.Exec(seckillTable)
	if err != nil {
		log.Printf("create productTable error: %v", err)
	}
	log.Println("create productTable success")
}

func InitTable() {
	initProductData := "insert into product(product_name,total) values " +
		"('1000元秒杀iphone7s',100)," +
		"('800元秒杀ipadAir',150)," +
		"('500元秒杀华为P9',200)," +
		"('200元秒杀红米NOTE',300);"

	_, err := DB.Exec(initProductData)
	if err != nil {
		log.Printf("init product data error: %v", err)
	}
	log.Println("init product data success")

	initSeckillData := "insert into seckill(name,number,start_time,end_time) values " +
		"('1000元秒杀iphone7s',100,'2017-04-14 00:00:00','2017-04-15 00:00:00')," +
		"('800元秒杀ipadAir',150,'2017-04-14 00:00:00','2017-04-15 00:00:00')," +
		"('500元秒杀华为P9',200,'2017-04-14 00:00:00','2017-04-15 00:00:00')," +
		"('200元秒杀红米NOTE',300,'2017-04-14 00:00:00','2017-04-15 00:00:00');"

	_, err = DB.Exec(initSeckillData)
	if err != nil {
		log.Printf("init seckill data error: %v", err)
	}
	log.Println("init seckill data success")
}
