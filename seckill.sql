CREATE TABLE `User` (
  `id` bigint(20) NOT NULL COMMENT '用户ID ,手机号码',
  `nickname` varchar(255) DEFAULT NULL COMMENT '用户ID',
  `password` varchar(32) DEFAULT NULL COMMENT 'MD5(MD5(pass+固定salt) + salt)',
  `register_date` datetime DEFAULT NULL COMMENT '注册时间',
  `last_login_date` datetime DEFAULT NULL COMMENT '上次登录时间',
  `login_count` int(11) DEFAULT NULL COMMENT '登录次数',
  PRIMARY KEY (`id`,`head`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户信息表';


CREATE TABLE `BlackList` (
  `id` bigint(20) NOT NULL COMMENT '用户ID ,手机号码',
  `type` bit(1) NOT NULL DEFAULT b'0'  COMMENT '0:白名单, 1:黑名单'
 _`create_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '黑名单信息表';


--秒杀成功信息记录表：
CREATE TABLE success_killed(
  `seckill_id` BIGINT NOT NULL COMMENT '秒杀商品id',
  `user_phone` BIGINT NOT NULL COMMENT '用户手机号',
  `state` TINYINT NOT NULL DEFAULT -1 COMMENT '状态标识：-1无效，0成功，1已付款',
  `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '秒杀成功时间',
  PRIMARY KEY (seckill_id, user_phone), /*联合主键，防止重复秒杀*/
  KEY idx_create_time (create_time)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='秒杀成功明细表';


CREATE TABLE `product` (
  `product_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '商品Id',
  `product_name` varchar(50) NOT NULL DEFAULT '' COMMENT '商品名称',
  `total` int(5) unsigned NOT NULL DEFAULT '0' COMMENT '商品数量',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '商品状态',
  PRIMARY KEY (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='商品数据表';


CREATE TABLE seckill(
  `seckill_id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '商品库存id',
  `name` VARCHAR(120) NOT NULL COMMENT '商品名称',
  `number` INT NOT NULL COMMENT '库存数量',
  `start_time` TIMESTAMP NOT NULL COMMENT '秒杀开始时间',
  `end_time` TIMESTAMP NOT NULL COMMENT '秒杀结束时间',
  `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (seckill_id),
  KEY idx_start_time (start_time),
  KEY idx_end_time (end_time),
  KEY idx_create_time (create_time)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET= utf8mb4  COMMENT='秒杀库存表';

-- 初始化数据
insert into seckill(name,number,start_time,end_time)
values
  ('1000元秒杀iphone7s',100,'2017-04-14 00:00:00','2017-04-15 00:00:00'),
  ('800元秒杀ipadAir',150,'2017-04-14 00:00:00','2017-04-15 00:00:00'),
  ('500元秒杀华为P9',200,'2017-04-14 00:00:00','2017-04-15 00:00:00'),
  ('200元秒杀红米NOTE',300,'2017-04-14 00:00:00','2017-04-15 00:00:00');
