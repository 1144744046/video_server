

CREATE TABLE `comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `vidio_id` varchar(11) DEFAULT NULL,
  `author_id` int(11) DEFAULT NULL,
  `content` varchar(255) DEFAULT NULL,
  `time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
# Data for table "comment"
#


#
# Structure for table "session"
#

CREATE TABLE `session` (
  `session_id` varchar(255) NOT NULL DEFAULT '',
  `login_name` varchar(255) DEFAULT NULL,
  `ttl` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
# Data for table "session"
#

INSERT INTO `session` VALUES ('f28be68c-871e-4fc8-a10d-5994f73ce79f','siliang','15298060768443'),('509843ff-6800-4433-a358-71fe9a847974','test','15301933538061'),('fb09595b-07a7-4c33-b023-8f9c8d64abee','test','15301933538344'),('41e0aa24-b577-4fed-9174-ddef50f76f44','test','15301933538349'),('4bab171b-4136-4126-8996-a04539cb2626','test','15301933538354'),('7eece4bd-9616-41e7-9108-ec27419307f6','test','15301933538349'),('48f18367-e919-4f5e-82c5-784065e9837b','test','15302364334121'),('2288e7f1-5333-474c-ad29-1e12b6fb7449','test','15302368574164'),('c7b17ea9-095f-49d3-bf07-0f9890385a00','siliang','15302369798938'),('c525caf7-a27d-4066-b0cd-182b4679dffe','jdlaj','15302371721066'),('8e6e2251-5897-4545-892b-884883e19aff','djakl','15302373035451'),('20997f51-2203-4985-b764-acf5b0816ed6','1dvx','15302374157034'),('f693faff-b085-463f-8acd-1079d61e4ffe','aaav','15302374770517'),('c0b7724a-5de2-4465-b82f-c7a48215c733','adfc','15302376407114'),('f919611c-dd56-4b38-84e1-af20823a5cbc','cgfa','15302376720451'),('cde03fdd-c2b4-49f3-8777-d7403832a6b0','sf','15302384381920'),('7af955bd-a9a4-40b5-bf67-1ec5800ffac1','leyuzhuren','15302388736472'),('ce44bf4a-321e-4eed-9b65-902e3eb8910d','test','15302404065418'),('e8d95ee7-64e7-43ea-a623-9086ea4dcee7','test','15302405728468'),('62aff6cc-ffd7-42af-a029-64364ae5bbb0','test','15302405870001'),('209de042-4067-4198-bd75-a16fe5155f19','test','15302438147614'),('75a66456-4232-401d-91b6-c6566f11ca82','test','15302762794271'),('4c9a7c27-adb5-4993-ac95-196a1a2d3bfc','test','15303340698993'),('371a8971-05bb-4b50-91a7-d4c1118cd06a','test','15303352749520');

#
# Structure for table "users"
#

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) DEFAULT NULL,
  `pwd` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;

#
# Data for table "users"
#

INSERT INTO `users` VALUES (1,'siliang','123456'),(2,'test','123'),(3,'test','123'),(4,'test','123'),(5,'test','123'),(6,'test','123'),(7,'test','123'),(8,'test','wrtts123'),(9,'siliang','huahauf'),(10,'jdlaj','fgdkjal'),(11,'djakl','cvc'),(12,'1dvx','dag'),(13,'aaav','ced'),(14,'adfc','s234'),(15,'cgfa','dadgf'),(16,'sf','adfg'),(17,'leyuzhuren','dkajlhf'),(18,'test','123');

#
# Structure for table "vidio_del_rec"
#

CREATE TABLE `vidio_del_rec` (
  `video_id` varchar(256) NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
# Data for table "vidio_del_rec"
#


#
# Structure for table "vidio_info"
#

CREATE TABLE `vidio_info` (
  `id` varchar(255) NOT NULL DEFAULT '',
  `name` varchar(255) DEFAULT NULL,
  `author_id` int(11) DEFAULT NULL,
  `display_ctime` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
# Data for table "vidio_info"
#

INSERT INTO `vidio_info` VALUES ('42289d65-0256-42a4-b86f-872b4eabfc89','test',0,'2018-29-66 12:01:39',NULL),('6f8ae171-60c1-461d-9749-732e026b63e5','test',2,'2018-29-66 12:02:18',NULL),('20eb38e3-9d83-4332-b80c-191717bfcf2f','04.Docker基本概念和框架 -Docker 容器相关技术简介.mp4',2,'2018-29-66 17:14:20',NULL),('ba4278bf-457e-4156-b190-00e41577a58f','03.Docker基本概念和框架 -Docker 的基本组成.mp4',2,'2018-30-66 12:44:58',NULL);
