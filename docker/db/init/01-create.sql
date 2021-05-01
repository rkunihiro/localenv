DROP TABLE IF EXISTS `User`;
CREATE TABLE `User` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,

    `name` varchar(255) NOT NULL,

    `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `User` (`name`) VALUES
('foo'), ('bar');

