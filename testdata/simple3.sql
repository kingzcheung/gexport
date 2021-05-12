CREATE TABLE `tracking` (
                            `tracking_id` int(11) NOT NULL AUTO_INCREMENT,
                            `url_id` int(11) NOT NULL,
                            `ip` varchar(20) DEFAULT '',
                            `source` varchar(200) DEFAULT '',
                            `region` varchar(200) DEFAULT '',
                            `os` varchar(50) DEFAULT '',
                            `device` varchar(50) DEFAULT '',
                            `browser` varchar(50) DEFAULT '',
                            `created_at` int(11) NOT NULL,
                            PRIMARY KEY (`tracking_id`),
                            KEY `url_id` (`url_id`)
) ENGINE=InnoDB AUTO_INCREMENT=387 DEFAULT CHARSET=utf8mb4