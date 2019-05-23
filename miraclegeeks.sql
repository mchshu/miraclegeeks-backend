DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`username` VARCHAR(64) NOT NULL,
	`email` VARCHAR(64) NOT NULL,
	`password` VARCHAR(64) NOT NULL,
	`created` DATETIME NOT NULL
);

DROP TABLE IF EXISTS `project`;
CREATE TABLE `project` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`name` VARCHAR(64) NOT NULL,
	`avatar` VARCHAR(64) NOT NULL,
	`description` VARCHAR(64) NOT NULL,
    `period` INTEGER NOT NULL,
	`created` DATETIME NOT NULL
);

INSERT INTO `project` (`name`, `avatar`, `description`, `period`, `created`) VALUES
(
	'e-commerce website of medicine', 
	'/images/bitmap0.png', 
	'This is a brief description of the project, no more than one lines.',
	1,
	"2019-05-25 10:00:00"
),
(
	'e-commerce website of medicine', 
	'/images/bitmap1.png', 
	'This is a brief description of the project, no more than one lines.',
	2,
	"2019-05-25 10:00:00"
),
(
	'e-commerce website of medicine', 
	'/images/bitmap2.png', 
	'This is a brief description of the project, no more than one lines.',
	3,
	"2019-05-25 10:00:00"
),
(
	'e-commerce website of medicine', 
	'/images/bitmap3.png', 
	'This is a brief description of the project, no more than one lines.',
	10,
	"2019-05-25 10:00:00"
),
(
	'e-commerce website of medicine', 
	'/images/bitmap4.png', 
	'This is a brief description of the project, no more than one lines.',
	11,
	"2019-05-25 10:00:00"
),
(
	'e-commerce website of medicine', 
	'/images/bitmap0.png', 
	'This is a brief description of the project, no more than one lines.',
	11,
	"2019-05-25 10:00:00"
),
(
	'e-commerce website of medicine', 
	'/images/bitmap5.png', 
	'This is a brief description of the project, no more than one lines.',
	15,
	"2019-05-25 10:00:00"
),
(
	'e-commerce website of medicine', 
	'/images/bitmap0.png', 
	'This is a brief description of the project, no more than one lines.',
	15,
	"2019-05-25 10:00:00"
),
(
	'e-commerce website of medicine', 
	'/images/bitmap1.png', 
	'This is a brief description of the project, no more than one lines.',
	10,
	"2019-05-25 10:00:00"
),
(
	'e-commerce website of medicine', 
	'/images/bitmap2.png', 
	'This is a brief description of the project, no more than one lines.',
	16,
	"2019-05-25 10:00:00"
),
(
	'e-commerce website of medicine', 
	'/images/bitmap3.png', 
	'This is a brief description of the project, no more than one lines.',
	18,
	"2019-05-25 10:00:00"
)