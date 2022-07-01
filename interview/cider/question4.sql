CREATE TABLE `student` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `score` int DEFAULT NULL,
  `course` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `student` VALUES ('1', '张三', '122', '语文');
INSERT INTO `student` VALUES ('2', '张三', '123', '英语');
INSERT INTO `student` VALUES ('3', '李四', '131', '数学');
INSERT INTO `student` VALUES ('4', '李四', '132', '语文');
INSERT INTO `student` VALUES ('5', '李四', '133', '英语');
INSERT INTO `student` VALUES ('6', '张三', '121', '数学');

/*求学生表中每个学生的平均分数*/
SELECT name,AVG(score) from student GROUP BY name