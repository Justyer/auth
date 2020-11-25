CREATE TABLE `auth_user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `nick` varchar(8) NOT NULL COMMENT '昵称',
  `phone` char(11) NOT NULL COMMENT '手机号',
  `name` varchar(15) NOT NULL COMMENT '账号',
  `pass` varchar(16) NOT NULL COMMENT '密码',
  `status` enum('enable','disable') NOT NULL DEFAULT 'enable' COMMENT '状态',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `idx_uni` (`phone`),
  KEY `idx_ctime` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='权限-用户表';

CREATE TABLE `auth_role` (
  `role_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(15) NOT NULL COMMENT '账号',
  `status` enum('enable','disable') NOT NULL DEFAULT 'enable' COMMENT '状态',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`role_id`),
  KEY `idx_ctime` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='权限-角色表';

CREATE TABLE `auth_user_role` (
  `user_id` int(11) NOT NULL,
  `role_id` int(11) NOT NULL,
  `status` enum('enable','disable') NOT NULL DEFAULT 'enable' COMMENT '状态',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`, `role_id`),
  KEY `idx_ctime` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='权限-用户角色关联表';