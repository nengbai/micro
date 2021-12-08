
--
--
-- 数据库： ``
--

-- --------------------------------------------------------

--
-- 表的结构 `users`
--
CREATE TABLE `users` (
  `userId` bigint UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'id',
  `role` tinyint UNSIGNED NOT NULL DEFAULT '0' COMMENT '角色: 0 普通用户，1 管理员',
  `type` tinyint UNSIGNED NOT NULL DEFAULT '0' COMMENT '类型: 0 未付费 ，1 付费',
  `name` varchar(500) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(20) NOT NULL DEFAULT '' COMMENT '密码',
  `email` varchar(20) NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '电话号码',
  `person_id` varchar(18) NOT NULL DEFAULT '' COMMENT '生份证号',
  `emp_id` varchar(100) NOT NULL DEFAULT '' COMMENT '员工号',
  `gender` tinyint UNSIGNED NOT NULL DEFAULT '0' COMMENT '性别：0 女，1 男',
  `buDomain` varchar(100) NOT NULL DEFAULT '' COMMENT '部门',
  `typeDevice` varchar(100) NOT NULL DEFAULT '' COMMENT '设备类型',
  `createTime` datetime NOT NULL DEFAULT '2020-11-19 00:00:00' COMMENT '创建时间',
  `isHot` tinyint UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否热榜,0:非热，1，热',
  `approvalStaffId` int UNSIGNED NOT NULL DEFAULT '0' COMMENT '审批人',
  `restSum` int UNSIGNED NOT NULL DEFAULT '0' COMMENT '重置密码次数',
  `failSum` int UNSIGNED NOT NULL DEFAULT '0' COMMENT '登陆失败次数',
  `isLogin` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0,离线，1，在线'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户';

--

INSERT INTO `Users` (`userId`,`name`,`password`,`email`,`phone`) VALUES (0,'baineng','zaq12wsx','nengbai@aliyun.com','18601369917')  