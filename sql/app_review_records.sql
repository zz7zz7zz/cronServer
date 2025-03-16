CREATE TABLE `app_review_records` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一主键ID（自增）',
  `ver` varchar(64) NOT NULL COMMENT '应用版本号（格式如 1.2.0）',
  `pkg` varchar(255) NOT NULL COMMENT '应用包名（如 com.example.app）',
  `platform` varchar(32) NOT NULL COMMENT '平台类型（iOS/Android/Web）',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '审核状态（0:待审, 1:通过, 2:拒绝）',
  `time_stamp` int(11) NOT NULL COMMENT '记录时间戳（Unix秒级）',
  `task_status` int(11) NOT NULL DEFAULT '0' COMMENT '任务状态（0:未开始, 1:进行中,2:完成,3:停止）',
  `channel` varchar(128) NOT NULL DEFAULT '' COMMENT '分发渠道标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='应用审核记录表';