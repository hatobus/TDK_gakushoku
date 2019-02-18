CREATE DATABASE IF NOT EXISTS helpnamiki;
DROP TABLE IF EXISTS `student`;

CREATE TABLE `student`(
    `no` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(64) NOT NULL,
    `sumofcoin` int(11) NOT NULL,
    `lastworked` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`no`)
);
