CREATE TABLE `app_review_records` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一主键ID（自增）',
  `ver` varchar(64) NOT NULL COMMENT '应用版本号（格式如 1.2.0）',
  `pkg` varchar(255) NOT NULL COMMENT '应用包名（Android是包名，如com.example.app；Ios是应用落地页id字符后面那一串数字 如id1234567890）',
  `platform` varchar(32) NOT NULL COMMENT '平台类型（iOS/Android/Web）',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '审核状态（0:待审, 1:通过, 2:拒绝，3：版本已过期（线上有更新的版本）',
  `approve_ts` int(11) NOT NULL COMMENT '审核通过的时间戳（Unix秒级）',
  `task_status` int(11) NOT NULL DEFAULT '0' COMMENT '任务状态（0:未开始, 1:进行中,2:完成,3:停止）',
  `task_create_ts` int(11) NOT NULL COMMENT '任务创建的时间戳（Unix秒级）',
  `channel` varchar(128) NOT NULL DEFAULT '' COMMENT '分发渠道标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='应用审核记录表';