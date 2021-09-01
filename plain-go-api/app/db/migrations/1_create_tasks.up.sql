-- たたきとして作成
-- 最初はここについてもinit.sql側で対応する。
CREATE TABLE `tasks` (
	`id` bigint(20) NOT NULL AUTO_INCREMENT,
	`title` varchar(50) NOT NULL,
	`description` varchar(50) NOT NULL,
	`deadline_at` datetime NOT NULL,
	`created_at` datetime NOT NULL,
	`updated_at` datetime NOT NULL,
	PRIMARY KEY (`id`)
) DEFAULT CHARSET = utf8mb4;