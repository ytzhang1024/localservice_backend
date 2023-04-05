/*
SQLyog Ultimate v12.5.1 (64 bit)
MySQL - 5.7.27-log : Database - gorm
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`gorm` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `gorm`;

/*Table structure for table `additional_info` */

DROP TABLE IF EXISTS `additional_info`;

CREATE TABLE `additional_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) unsigned NOT NULL,
  `order_id` bigint(20) unsigned NOT NULL,
  `message` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_service_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;

/*Data for the table `additional_info` */

insert  into `additional_info`(`id`,`created_at`,`updated_at`,`deleted_at`,`user_id`,`order_id`,`message`) values 
(1,'2023-04-01 16:56:25.441','2023-04-01 16:56:25.441',NULL,0,0,'testestest'),
(2,'2023-04-01 16:56:44.758','2023-04-01 16:56:44.758',NULL,0,0,'testestest'),
(3,'2023-04-01 16:57:22.484','2023-04-01 16:57:22.484',NULL,1,10,'testestest'),
(4,'2023-04-01 16:57:28.159','2023-04-01 16:57:28.159',NULL,1,10,''),
(5,'2023-04-01 16:57:30.175','2023-04-01 16:57:30.175',NULL,1,0,''),
(6,'2023-04-01 16:57:33.018','2023-04-01 16:57:33.018',NULL,1,10,'testestest'),
(7,'2023-04-01 16:59:49.511','2023-04-01 16:59:49.511',NULL,1,10,'testestest'),
(8,'2023-04-01 17:00:01.221','2023-04-01 17:00:01.221',NULL,1,10,'testestest');

/*Table structure for table `category` */

DROP TABLE IF EXISTS `category`;

CREATE TABLE `category` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

/*Data for the table `category` */

insert  into `category`(`id`,`name`) values 
(1,'Cleaning'),
(2,'Babysitting'),
(3,'Pest Control'),
(4,'Plumbing'),
(5,'Electrical Repairs'),
(6,'Beauty');

/*Table structure for table `order` */

DROP TABLE IF EXISTS `order`;

CREATE TABLE `order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `customer_id` bigint(20) unsigned NOT NULL,
  `service_id` bigint(20) unsigned NOT NULL,
  `requirements` longtext,
  `status` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_service_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=latin1;

/*Data for the table `order` */

insert  into `order`(`id`,`created_at`,`updated_at`,`deleted_at`,`customer_id`,`service_id`,`requirements`,`status`) values 
(1,'2023-04-01 14:12:03.866','2023-04-01 14:12:03.866',NULL,0,0,'',''),
(2,'2023-04-01 15:03:40.348','2023-04-01 15:03:40.348',NULL,1,0,'test',''),
(3,'2023-04-01 15:03:58.747','2023-04-01 15:03:58.747',NULL,1,0,'test',''),
(4,'2023-04-01 15:06:54.217','2023-04-01 15:06:54.217',NULL,1,0,'test',''),
(5,'2023-04-01 15:07:21.996','2023-04-01 15:07:21.996',NULL,11,0,'test',''),
(6,'2023-04-01 15:08:14.695','2023-04-01 15:08:14.695',NULL,11,0,'test',''),
(7,'2023-04-01 15:08:45.073','2023-04-01 15:08:45.073',NULL,11,12,'test',''),
(8,'2023-04-01 15:24:52.304','2023-04-01 15:24:52.304',NULL,11,12,'test','Pending'),
(9,'2023-04-01 15:26:54.923','2023-04-01 15:26:54.923',NULL,11,12,'test','Pending');

/*Table structure for table `review` */

DROP TABLE IF EXISTS `review`;

CREATE TABLE `review` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) unsigned NOT NULL,
  `service_id` bigint(20) unsigned NOT NULL,
  `content` longtext NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

/*Data for the table `review` */

insert  into `review`(`id`,`created_at`,`updated_at`,`deleted_at`,`user_id`,`service_id`,`content`) values 
(1,'2023-04-01 04:13:26.858','2023-04-01 04:13:26.858',NULL,0,0,''),
(2,'2023-04-01 04:13:32.770','2023-04-01 04:13:32.770',NULL,0,0,''),
(3,'2023-04-01 04:13:33.925','2023-04-01 04:13:33.925','2023-04-01 04:41:21.707',0,0,''),
(4,'2023-04-01 04:15:20.522','2023-04-01 04:15:20.522',NULL,1,1,'\"test\"'),
(5,'2023-04-01 04:15:38.702','2023-04-01 04:15:38.702',NULL,1,1,'this is a review'),
(6,'2023-04-01 04:19:21.407','2023-04-01 04:19:21.407','2023-04-01 04:25:52.099',1,1,'this is a review');

/*Table structure for table `service` */

DROP TABLE IF EXISTS `service`;

CREATE TABLE `service` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) unsigned NOT NULL,
  `city` longtext,
  `title` longtext,
  `description` longtext,
  `prices` bigint(20) DEFAULT NULL,
  `address` longtext,
  `areas_coverd` bigint(20) DEFAULT NULL,
  `category` longtext,
  `availibility` longtext,
  `mobile` longtext,
  `photos` longtext,
  `longitude_and_latitude` longtext,
  `status` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_service_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

/*Data for the table `service` */

insert  into `service`(`id`,`created_at`,`updated_at`,`deleted_at`,`user_id`,`city`,`title`,`description`,`prices`,`address`,`areas_coverd`,`category`,`availibility`,`mobile`,`photos`,`longitude_and_latitude`,`status`) values 
(3,'2023-04-01 03:31:14.490','2023-04-01 03:31:14.490',NULL,0,'','1','1',1,'1',0,'1','','1','','','Pending'),
(4,'2023-04-01 03:31:31.430','2023-04-01 03:31:31.430',NULL,0,'','1','1',1,'1',0,'1','','1','','','Pending');

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `email` varchar(128) NOT NULL,
  `password` varchar(128) NOT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `mobile` varchar(11) DEFAULT NULL,
  `nick_name` varchar(128) DEFAULT NULL,
  `role` varchar(128) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

/*Data for the table `user` */

insert  into `user`(`id`,`created_at`,`updated_at`,`deleted_at`,`email`,`password`,`avatar`,`mobile`,`nick_name`,`role`) values 
(1,'2023-04-01 02:55:27.778','2023-04-01 02:55:27.778',NULL,'test','$2a$10$OvXk7ixo94sFf9J2/UAOBew3tpDsFH8l1rkqpaEjlbrM9VGL3PwOW','','','','a'),
(2,'2023-04-01 04:33:46.118','2023-04-01 04:33:46.118','2023-04-01 04:33:53.644','test1','$2a$10$e5F9qfG3/NKQ7bXDGmDWPOzk5midLHcGnx5lNcX2.fpogmZaca0XG','','','','a'),
(3,'2023-04-05 14:22:43.407','2023-04-05 14:22:43.407',NULL,'test1','$2a$10$2d85HWhHkFiB2RLDi/O3BeYI7j7yTVrgagkST8OVJrVzAhNci5ezS','','','','a');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
