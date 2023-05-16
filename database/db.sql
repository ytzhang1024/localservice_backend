/*
SQLyog Ultimate
MySQL - 5.7.41-log : Database - db6251
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`db6251` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `db6251`;

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
  `status` varchar(100) NOT NULL,
  `request` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_service_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=latin1;

/*Data for the table `additional_info` */

insert  into `additional_info`(`id`,`created_at`,`updated_at`,`deleted_at`,`user_id`,`order_id`,`message`,`status`,`request`) values 
(14,'2023-05-14 23:29:54.938','2023-05-14 23:55:19.193',NULL,76,43,'I need addr','Completed',''),
(15,'2023-05-14 23:31:34.300','2023-05-14 23:57:08.190',NULL,76,43,'I need addr','Completed',''),
(16,'2023-05-15 00:00:08.715','2023-05-15 00:00:38.260',NULL,76,44,'123Cest','Completed',''),
(17,'2023-05-15 02:35:05.266','2023-05-15 02:43:35.423',NULL,76,56,'Please enter more...','Completed',''),
(18,'2023-05-15 02:36:41.032','2023-05-15 02:36:41.032',NULL,76,56,'Please enter more....','Pending',''),
(19,'2023-05-15 02:37:26.452','2023-05-15 02:37:26.452',NULL,76,56,'please enter more....','Pending',''),
(20,'2023-05-15 03:23:19.892','2023-05-15 03:24:18.385',NULL,79,59,'Please enter more detailed description','Completed',''),
(21,'2023-05-15 03:47:04.621','2023-05-15 03:47:23.560',NULL,63,62,'Please','Completed',''),
(22,'2023-05-15 04:00:04.493','2023-05-15 04:00:18.407',NULL,63,63,'Please888888','Completed',''),
(23,'2023-05-15 04:00:48.038','2023-05-15 04:01:36.006',NULL,63,63,'Please add review','Completed',''),
(24,'2023-05-15 13:28:28.485','2023-05-15 13:44:31.158',NULL,63,66,'Please...','Completed',''),
(25,'2023-05-15 13:28:38.330','2023-05-15 14:31:47.652',NULL,63,67,'Please add review','Completed',''),
(26,'2023-05-15 14:19:32.289','2023-05-15 14:19:32.289',NULL,63,66,'I want to..','Pending',''),
(27,'2023-05-15 14:19:53.706','2023-05-15 14:19:53.706',NULL,63,67,'Please add review','Pending','');

/*Table structure for table `category` */

DROP TABLE IF EXISTS `category`;

CREATE TABLE `category` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

/*Data for the table `category` */

/*Table structure for table `message` */

DROP TABLE IF EXISTS `message`;

CREATE TABLE `message` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) unsigned NOT NULL,
  `order_id` bigint(20) unsigned NOT NULL,
  `message` longtext,
  `status` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_service_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=latin1;

/*Data for the table `message` */

insert  into `message`(`id`,`created_at`,`updated_at`,`deleted_at`,`user_id`,`order_id`,`message`,`status`) values 
(22,'2023-05-15 02:35:05.330','2023-05-15 03:33:55.303',NULL,63,56,'Please update descrition','Completed'),
(23,'2023-05-15 02:35:13.878','2023-05-15 02:35:13.878',NULL,76,57,'Please add review','Completed'),
(24,'2023-05-15 02:36:41.035','2023-05-15 02:36:41.035',NULL,76,56,'Please update descrition','Pending'),
(25,'2023-05-15 02:37:26.452','2023-05-15 02:37:26.452',NULL,76,56,'Please update descrition','Pending'),
(26,'2023-05-15 03:23:19.901','2023-05-15 03:23:19.901',NULL,79,59,'Please update descrition','Pending'),
(27,'2023-05-15 03:23:32.889','2023-05-15 03:23:32.889',NULL,79,58,'Please add review','Completed'),
(28,'2023-05-15 03:47:04.621','2023-05-15 03:47:04.621',NULL,63,62,'Please update descrition','Update'),
(29,'2023-05-15 03:47:29.756','2023-05-15 03:47:29.756',NULL,63,62,'Please add review','Completed'),
(30,'2023-05-15 04:00:04.494','2023-05-15 04:00:58.096',NULL,63,63,'Please update descrition','Completed'),
(31,'2023-05-15 04:00:48.036','2023-05-15 04:01:15.199',NULL,63,63,'Please add review','Completed'),
(32,'2023-05-15 13:28:28.496','2023-05-15 13:45:12.354',NULL,63,66,'Please update descrition','Completed'),
(33,'2023-05-15 13:28:38.315','2023-05-15 13:50:06.850',NULL,63,67,'Please add review','Completed'),
(34,'2023-05-15 14:19:32.272','2023-05-15 14:32:33.770',NULL,63,66,'Please update descrition','Completed'),
(35,'2023-05-15 14:19:53.706','2023-05-15 14:19:53.706',NULL,63,67,'Please add review','Pending');

/*Table structure for table `order` */

DROP TABLE IF EXISTS `order`;

CREATE TABLE `order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `customer_id` bigint(20) unsigned NOT NULL,
  `service_id` bigint(20) unsigned NOT NULL,
  `provider_id` bigint(20) unsigned NOT NULL,
  `customer_name` varchar(100) DEFAULT NULL,
  `customer_email` varchar(100) DEFAULT NULL,
  `customer_phone` varchar(100) DEFAULT NULL,
  `service_title` varchar(100) DEFAULT NULL,
  `postcode` varchar(100) DEFAULT NULL,
  `address` varchar(100) DEFAULT NULL,
  `city` varchar(100) DEFAULT NULL,
  `date` varchar(100) DEFAULT NULL,
  `description` longtext,
  `status` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_service_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=70 DEFAULT CHARSET=latin1;

/*Data for the table `order` */

insert  into `order`(`id`,`created_at`,`updated_at`,`deleted_at`,`customer_id`,`service_id`,`provider_id`,`customer_name`,`customer_email`,`customer_phone`,`service_title`,`postcode`,`address`,`city`,`date`,`description`,`status`) values 
(48,'2023-05-15 00:40:30.879','2023-05-15 00:40:30.879','2023-05-15 02:48:38.000',76,91,64,'Tim','123@123.com','1234567890','Sparkle Clean','NE2 HN2','addr','Newcastle','Fri May 19 2023 01:40:17 GMT+0100','Desc','Pending'),
(51,'2023-05-15 00:51:26.999','2023-05-15 00:51:26.999','2023-05-15 02:48:38.000',76,103,75,'Tim','123@123.com','01234567890','Service','NE2 HN2','addr','Newcastle','Mon May 15 2023 00:07:00 GMT+0100','asd','Pending'),
(52,'2023-05-15 00:56:10.313','2023-05-15 00:56:10.313',NULL,76,103,75,'Tim','123@123.com','01234567890','Service','NE2 HN2','addr','Newcastle','Mon May 15 2023 00:06:00 GMT+0100','123','Pending'),
(53,'2023-05-15 01:08:03.324','2023-05-15 01:08:03.324',NULL,76,103,75,'Tim','123@123.com','01234567890','Service','NE2 HN2','addr','Newcastle','Mon May 15 2023 00:06:00 GMT+0100','adr','Pending'),
(63,'2023-05-15 03:59:34.556','2023-05-15 04:01:34.616',NULL,63,110,64,'Shuaiyue Xie','xieshuaiyue@gmail.com','07536919873','Pest control','SO14 0GE','Flat 604, Orions Point','Southampton','Thu May 11 2023 04:59:35 GMT+0100','huuuuuuuuu','Completed'),
(64,'2023-05-15 04:03:25.096','2023-05-15 14:17:01.905',NULL,63,110,64,'Shuaiyue Xie','xieshuaiyue@gmail.com','07536919873','Pest control','SO14 0GE','Flat 604, Orions Point','Southampton','Thu May 25 2023 05:03:07 GMT+0100','I want my house no pesst','Accepted'),
(65,'2023-05-15 04:03:50.959','2023-05-15 14:17:12.054',NULL,63,118,64,'Shuaiyue Xie','xieshuaiyue@gmail.com','07536919973','Babysitting','SO14 0GE','Flat 604, Orions Point','Southampton','Mon May 15 2023 02:00:00 GMT+0100','take care my baby','Rejected'),
(66,'2023-05-15 04:04:20.418','2023-05-15 14:19:32.289',NULL,63,119,64,'Shuaiyue Xie','xieshuaiyue@gmail.com','07536919873','Beauty','SO14 0GE','Flat 604, Orions Point','Southampton','Tue May 23 2023 05:04:05 GMT+0100','7777777777777777777','Update'),
(67,'2023-05-15 04:04:41.133','2023-05-15 14:31:46.135',NULL,63,120,64,'Shuaiyue Xie','xieshuaiyue@gmail.com','07536919873','Cleaning','SO14 0GE','Flat 604, Orions Point','Southampton','Thu May 18 2023 05:08:40 GMT+0100','66666666666666666','Completed'),
(68,'2023-05-15 13:43:02.015','2023-05-15 13:43:02.015',NULL,63,112,65,'Shuaiyue Xie','xieshuaiyue@gmail.com','07536919873','Sparkle Clean','SO14 0GE','Flat 604, Orions Point','Southampton','Mon May 22 2023 16:42:35 GMT+0100','I want my house ....','Pending'),
(69,'2023-05-15 14:30:25.517','2023-05-15 14:30:25.517',NULL,63,110,64,'Shuaiyue Xie','xieshuaiyue@gmail.com','07536919873','Pest control','SO14 0GE','Flat 604, Orions Point','Southampton','Wed May 17 2023 17:30:17 GMT+0100','test','Pending');

/*Table structure for table `provider_msg` */

DROP TABLE IF EXISTS `provider_msg`;

CREATE TABLE `provider_msg` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) unsigned NOT NULL,
  `message` longtext,
  `status` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_service_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=latin1;

/*Data for the table `provider_msg` */

insert  into `provider_msg`(`id`,`created_at`,`updated_at`,`deleted_at`,`user_id`,`message`,`status`) values 
(9,'2023-05-14 22:22:22.257','2023-05-14 22:22:22.257',NULL,64,'Please tell me more detail',''),
(10,'2023-05-15 13:32:57.787','2023-05-15 13:32:57.787',NULL,66,'PLease give me detail',''),
(11,'2023-05-15 14:22:15.218','2023-05-15 14:22:15.218',NULL,66,'We want more detail','');

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
  `rating` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=latin1;

/*Data for the table `review` */

insert  into `review`(`id`,`created_at`,`updated_at`,`deleted_at`,`user_id`,`service_id`,`content`,`rating`) values 
(19,'2023-05-15 03:33:52.931','2023-05-15 03:33:52.931',NULL,63,119,'ITS SO GOOD!!!!!!!',5),
(20,'2023-05-15 04:00:56.548','2023-05-15 04:00:56.548',NULL,63,110,'111111111111111111111',5),
(21,'2023-05-15 04:01:13.829','2023-05-15 04:01:13.829','2023-05-15 14:26:51.435',63,110,'4444444444444444',4),
(22,'2023-05-15 13:45:10.007','2023-05-15 13:45:10.007',NULL,63,119,'very good!',5),
(23,'2023-05-15 13:50:05.608','2023-05-15 13:50:05.608',NULL,63,120,'Very bad!!!!!!!!!!!!!!!!!!!!!!111',1),
(24,'2023-05-15 13:50:05.608','2023-05-15 13:50:05.608',NULL,65,112,'Not bad',3),
(25,'2023-05-15 13:50:05.608','2023-05-15 13:50:05.608',NULL,65,112,'Rubbish',2),
(26,'2023-05-15 14:32:31.869','2023-05-15 14:32:31.869',NULL,63,119,'test',4);

/*Table structure for table `service` */

DROP TABLE IF EXISTS `service`;

CREATE TABLE `service` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  `city` longtext,
  `title` longtext,
  `description` longtext,
  `prices` bigint(20) DEFAULT NULL,
  `address` longtext,
  `areas_coverd` bigint(20) DEFAULT NULL,
  `category` longtext,
  `availibility` longtext,
  `mobile` longtext,
  `photos` varchar(1000) DEFAULT NULL,
  `longitude_and_latitude` longtext,
  `status` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_service_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=129 DEFAULT CHARSET=latin1;

/*Data for the table `service` */

insert  into `service`(`id`,`created_at`,`updated_at`,`deleted_at`,`user_id`,`city`,`title`,`description`,`prices`,`address`,`areas_coverd`,`category`,`availibility`,`mobile`,`photos`,`longitude_and_latitude`,`status`) values 
(99,'2023-05-14 23:01:11.998','2023-05-14 23:01:11.998','2023-05-14 23:58:33.389',64,'1','1','1',1,'1',0,'Cleaning','1','07462264811','upload/1684156107152315637Cleaning3.jpg','1','Pending'),
(100,'2023-05-14 23:03:04.109','2023-05-14 23:03:04.109','2023-05-14 23:58:35.943',64,'1','1','1',1,'1',0,'Cleaning','1','04746378222','upload/1684156107152315637Cleaning3.jpg','1','Pending'),
(101,'2023-05-14 23:13:55.559','2023-05-14 23:13:55.559','2023-05-14 23:58:37.636',64,'1','1','1',11,'1',0,'Cleaning','1','08238383333','upload/1684156107152315637Cleaning3.jpg','1','Pending'),
(102,'2023-05-14 23:15:28.241','2023-05-14 23:15:28.241','2023-05-15 00:10:02.113',64,'123','1','1',1,'123',0,'Cleaning','1','01231231232','upload/1684156107152315637Cleaning3.jpg','1','Pending'),
(104,'2023-05-15 00:06:40.675','2023-05-15 00:06:40.675','2023-05-15 00:09:45.117',64,'London','123','87',5,'123',0,'Babysitting','787','01231233232','upload/1684156060868644062Babysitting2.jpg','London','Pending'),
(105,'2023-05-15 00:23:34.260','2023-05-15 00:23:34.260','2023-05-15 00:09:45.117',64,'London','The Gym Group Southampton','The Gym Group Southampton Portswood, Portswood Road, Southampton',10,'The Gym Group Southampton Portswood',0,'Plumbing','24/7','08236784721','upload/logo.png','50.91102832164159, -1.3979382100044861','Pending'),
(106,'2023-05-15 00:24:41.367','2023-05-15 00:24:41.367','2023-05-15 00:09:45.117',64,'London','The Gym','The Gym Group Southampton Portswood, Portswood Road, Southampton',10,'Portswood Road, Southampton',0,'Babysitting','24/7','03746383713','upload/1684156060868644062Babysitting2.jpg','London','Pending'),
(107,'2023-05-15 00:32:03.687','2023-05-15 00:32:03.687','2023-05-15 00:09:45.117',64,'London','The Gym Group Southampton','The Gym Group Southampton',10,'The Gym Group Southampton',0,'Beauty','24/7','03827382732','upload/1684156089864589129Beauty2.jpg','London','Pending'),
(108,'2023-05-15 01:14:07.372','2023-05-15 01:14:07.372','2023-05-15 02:22:34.451',64,'London','Service3','r43423',10,'123452343',0,'Pest control','24/7','03434324234','upload/1684156030507167043pestcontrol.jpg','50.91102832164159, -1.3979382100044861','Pending'),
(109,'2023-05-15 01:19:09.520','2023-05-15 01:19:09.520','2023-05-15 02:24:58.335',64,'Southampton','Service2','321312334',10,'1231231',0,'Babysitting','24/7','03877323223','upload/1684156060868644062Babysitting2.jpg','50.91102832164159, -1.3979382100044861','Pending'),
(110,'2023-05-15 01:56:23.557','2023-05-15 13:07:10.511',NULL,64,'Southampton','Pest control','Pest control',88,'Orin Potin St Mary\'s Road, Southampton, England SO14 0GE',0,'Pest control','24/7','+447536144283','upload/1684156030507167043pestcontrol.jpg','48.88217501285378, 2.3644679623716547','Approved'),
(112,'2023-05-14 22:25:52.984','2023-05-14 22:51:31.615',NULL,65,'Bath','Sparkle Clean','We specialize in comprehensive residential cleaning services, including regular cleaning, deep cleaning, move-in/move-out cleaning, and post-construction cleaning. ',56,'pton  Marchwood',0,'Cleaning','24/7','07536919875','upload/1684156030507167043pestcontrol.jpg','50.89388022632951, -1.4491725313288744','Approved'),
(113,'2023-05-14 22:27:32.295','2023-05-14 22:52:23.348',NULL,67,'Bath','Fresh Sweep','We understand the importance of a clean and organized workspace. Our office cleaning services include dusting, vacuuming, sanitizing, trash removal, and restocking supplies. ',60,'SeaCity Museum',0,'Cleaning','24/7','07536919874','upload/1684156107152315637Cleaning3.jpg','50.9106647435273, -1.410430666321995','Approved'),
(114,'2023-05-14 22:28:50.593','2023-05-14 22:28:50.593',NULL,66,'London','Pure Shine','We cater to the cleaning needs of various commercial establishments, such as retail stores, restaurants, healthcare facilities, and educational institutions. ',23,'Ealing',0,'Cleaning','24/5','07536919973','upload/logo.png','51.506929484025356, -0.3315805284867823','Approved'),
(115,'2023-05-14 22:30:14.663','2023-05-14 22:53:13.709',NULL,66,'Bristol','Crystal Clear','Our window cleaning service ensures crystal-clear views and enhances the overall appearance of your space. ',78,'Clissold Park',0,'Cleaning','24/5','07536919973','upload/1684156107152315637Cleaning3.jpg','51.56384799367909, -0.0833705826380879','Approved'),
(116,'2023-05-14 22:31:48.908','2023-05-14 22:52:50.609',NULL,67,'Southampton','Shine Bright Cleaning','We offer specialized cleaning services for specific areas or items that require extra care.',67,'Romford',0,'Cleaning','24/5','07536919973','upload/1684156089864589129Beauty2.jpg','51.59308512673066, 0.20561617834583212','Approved'),
(117,'2023-05-14 22:33:50.116','2023-05-14 22:54:16.453',NULL,67,'Southampton','Crystal Maid','We offer specialized cleaning services for specific areas or items that require extra care.',89,'JORVIK Viking Centre',0,'Cleaning','24/7','07536919973','upload/1684156107152315637Cleaning3.jpg','53.96030642385765, -1.0999900446842084','Approved'),
(118,'2023-05-15 02:15:47.883','2023-05-15 13:07:40.872',NULL,64,'Southampton','Babysitting','Babysitting',77,'Orin Potin St Mary\'s Road, Southampton, England SO14 0GE',0,'Babysitting','24/7','+447536144283','upload/1684156060868644062Babysitting2.jpg','48.86823059789598, 2.2960593667699745','Approved'),
(119,'2023-05-15 02:17:31.592','2023-05-15 13:08:09.868',NULL,64,'Southampton','Beauty','Beauty',66,'Orin Potin St Mary\'s Road, Southampton, England SO14 0GE',0,'Beauty','24/7','+447536144283','upload/1684156089864589129Beauty2.jpg','48.86823059789598, 2.2960593667699745','Approved'),
(120,'2023-05-15 02:18:48.697','2023-05-15 13:08:27.156',NULL,64,'Southampton','Cleaning','Cleaning',55,'Orin Potin St Mary\'s Road, Southampton, England SO14 0GE',0,'Cleaning','24/7','+447536144283','upload/1684156107152315637Cleaning3.jpg','48.853625699529516, 2.2699320535838696','Approved'),
(121,'2023-05-15 03:04:11.170','2023-05-15 03:04:11.170',NULL,80,'Bath','Sparkle Clean Services','I can provider cleaning service',40,'Avon, 1A Forum Buildings, Bath BA1 1UG',0,'Cleaning','24/7','07364827122','upload/1684156107152315637Cleaning3.jpg','51.38138325249295, -2.373802454036227','Approved'),
(122,'2023-05-15 03:06:12.197','2023-05-15 03:06:12.197',NULL,80,'Bath','PowerPro Electricians','I can do electrical repairs',50,'Avon, 1A Forum Buildings, Bath BA1 1UG',0,'Electrical repairs','24/7','08237648123','upload/logo.png','51.381135491511955, -2.3626337366051016','Approved'),
(123,'2023-05-15 03:09:07.192','2023-05-15 03:09:07.192',NULL,80,'Bath','Tiny Treasures Babysitters','We are good babysitters',30,'1 North St, St Paul\'s, Bristol BS1 3PR',0,'Babysitting','24/7','07246438113','upload/1684156060868644062Babysitting2.jpg','51.4612242096892, -2.5900739261755708','Approved'),
(124,'2023-05-15 03:11:36.228','2023-05-15 03:11:36.228',NULL,80,'Newcastle','PowerPro Electricians','PowerPro Electricians offers expert electrical repair services to keep your home or business powered up and safe',70,'131 Sandyford Rd, Jesmond, Newcastle upon Tyne NE2 1QR',0,'Electrical repairs','24/7','02374658123','upload/logo.png','54.97537247882519, -1.6066097547691927','Approved'),
(126,'2023-05-15 03:15:46.902','2023-05-15 03:15:46.902','2023-05-15 14:25:01.010',80,'Bristol','FlowFix Plumbing Services','FlowFix Plumbing Services is your go-to provider for all plumbing needs.',20,'Callington Rd, Brislington, Bristol BS4 5AY',0,'Plumbing','24/7','03783232345','upload/logo.png','51.43320355770655, -2.5544133804974964','Approved'),
(128,'2023-05-15 14:14:58.040','2023-05-15 14:25:27.690',NULL,64,'','Test','Update',34,'Orin Potin St Mary\'s Road, Southampton, England SO14 0GE',0,'Plumbing','24/7','+447536144283','upload/1684160325438375733Plumbing2.jpg','48.88217501285378, 2.3644679623716547','Approved');

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
  `text` varchar(128) NOT NULL,
  `status` varchar(128) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=84 DEFAULT CHARSET=latin1;

/*Data for the table `user` */

insert  into `user`(`id`,`created_at`,`updated_at`,`deleted_at`,`email`,`password`,`avatar`,`mobile`,`nick_name`,`role`,`text`,`status`) values 
(1,NULL,NULL,NULL,'admin1@admin.com','$2a$10$tFAEEneNi.mz1e.nIsNGUejMsjCE.sS5k80rhT02dnClT8J6ApZRO','upload/profile.jpg',NULL,'Admin','Admin',' ',' '),
(63,'2023-05-14 22:00:41.153','2023-05-15 14:33:11.286',NULL,'xieshuaiyue@gmail.com','$2a$10$Y/isL5og2wxTw5BDSEzX3.SWBfOAJZ3gH2wltBf3Z864n6Kf/QnhW','upload/1684161191280981195profile4.png','09685746543','uu','Customer','','Pending'),
(64,'2023-05-14 22:07:07.630','2023-05-15 13:30:22.784',NULL,'wuyushu_u@163.com','$2a$10$cLWUMuYuMgAZTdd5mv0kwu.JhARohQi3c/cykT3RRqTLD.57Xjo7u','upload/1684157422780114967profile1.png','07846576432','Sparkling Home Services ','Provider','update','Approved'),
(65,'2023-05-14 22:07:07.630','2023-05-15 13:32:22.919',NULL,'1@163.com','$2a$10$cLWUMuYuMgAZTdd5mv0kwu.JhARohQi3c/cykT3RRqTLD.57Xjo7u','upload/profile.jpg','07846576432','Sparkling Home Services ','Provider','Our highly trained and experienced cleaners use eco-friendly products and state','Approved'),
(66,'2023-05-14 22:07:07.630','2023-05-15 14:23:05.103',NULL,'2@163.com','$2a$10$cLWUMuYuMgAZTdd5mv0kwu.JhARohQi3c/cykT3RRqTLD.57Xjo7u','upload/profile.jpg','07846576432','Sparkling Home Services ','Provider','Our highly trained and experienced cleaners use eco-friendly products and state','Approved'),
(67,'2023-05-14 22:07:07.630','2023-05-14 22:24:22.416',NULL,'3@163.com','$2a$10$cLWUMuYuMgAZTdd5mv0kwu.JhARohQi3c/cykT3RRqTLD.57Xjo7u','upload/profile.jpg','07846576432','Sparkling Home Services ','Provider','Our highly trained and experienced cleaners use eco-friendly products and state','Pending'),
(68,'2023-05-14 22:07:07.630','2023-05-14 22:24:22.416','2023-05-15 14:22:49.511','4@163.com','$2a$10$cLWUMuYuMgAZTdd5mv0kwu.JhARohQi3c/cykT3RRqTLD.57Xjo7u','upload/profile.jpg','07846576432','Sparkling Home Services ','Provider','Our highly trained and experienced cleaners use eco-friendly products and state','Pending'),
(69,'2023-05-14 22:07:07.630','2023-05-14 22:24:22.416',NULL,'5@163.com','$2a$10$cLWUMuYuMgAZTdd5mv0kwu.JhARohQi3c/cykT3RRqTLD.57Xjo7u','upload/profile.jpg','07846576432','Sparkling Home Services ','Provider','Our highly trained and experienced cleaners use eco-friendly products and state','Pending'),
(70,'2023-05-14 22:07:07.630','2023-05-14 22:24:22.416',NULL,'6@163.com','$2a$10$cLWUMuYuMgAZTdd5mv0kwu.JhARohQi3c/cykT3RRqTLD.57Xjo7u','upload/profile.jpg','07846576432','Sparkling Home Services ','Provider','Our highly trained and experienced cleaners use eco-friendly products and state','Pending'),
(71,'2023-05-14 22:07:07.630','2023-05-14 22:24:22.416',NULL,'7@163.com','$2a$10$cLWUMuYuMgAZTdd5mv0kwu.JhARohQi3c/cykT3RRqTLD.57Xjo7u','upload/profile.jpg','07846576432','Sparkling Home Services ','Provider','Our highly trained and experienced cleaners use eco-friendly products and state','Pending'),
(72,'2023-05-14 22:07:07.630','2023-05-14 22:24:22.416',NULL,'8@163.com','$2a$10$cLWUMuYuMgAZTdd5mv0kwu.JhARohQi3c/cykT3RRqTLD.57Xjo7u','upload/profile.jpg','07846576432','Sparkling Home Services ','Provider','Our highly trained and experienced cleaners use eco-friendly products and state','Pending'),
(73,'2023-05-14 22:07:07.630','2023-05-14 22:24:22.416',NULL,'9@163.com','$2a$10$cLWUMuYuMgAZTdd5mv0kwu.JhARohQi3c/cykT3RRqTLD.57Xjo7u','upload/profile.jpg','07846576432','Sparkling Home Services ','Provider','Our highly trained and experienced cleaners use eco-friendly products and state','Pending'),
(74,'2023-05-14 22:07:07.630','2023-05-14 22:24:22.416',NULL,'10@163.com','$2a$10$cLWUMuYuMgAZTdd5mv0kwu.JhARohQi3c/cykT3RRqTLD.57Xjo7u','upload/profile.jpg','07846576432','Sparkling Home Services ','Provider','Our highly trained and experienced cleaners use eco-friendly products and state','Pending'),
(75,'2023-05-14 23:20:37.618','2023-05-14 23:24:10.566',NULL,'provider@123.com','$2a$10$Mpe8c8S10aYhMyu7X9CZounb22p5qpZxjH8PZrfmpro0FPckYRplu','upload/profile.jpg','','','Provider','1234567890','Approved'),
(76,'2023-05-14 23:27:13.584','2023-05-15 13:07:35.414',NULL,'customer@123.com','$2a$10$ZgnGHh4AdPejijtMG29cEeE5A5/EUerMqBEApKTRlOinJbqpwDKtK','upload/168415605541041163020230509235645.jpg','07821628939','Nick','Customer','','Pending'),
(77,'2023-05-15 01:36:41.134','2023-05-15 01:36:41.134',NULL,'123@gmail.com','$2a$10$uMsyo0GJ5joxe20EeRODGeBkddinSD1rAi8DPiSbVwXhLKbdonItG','upload/profile.jpg','07536144283','provider','Provider','I want create a provider account','Pending'),
(80,'2023-05-15 03:01:23.501','2023-05-15 03:01:23.501',NULL,'adam@163.com','$2a$10$6FTjmATgSP.I/.S0krAvL.Unqu0DC6f91t0QyBj2DcaWKhf06Z1cO','upload/profile.jpg','08365527443','Adam','Provider','I am a good service provider','Approved'),
(82,'2023-05-15 14:08:47.931','2023-05-15 14:12:21.449',NULL,'yshwoou@gmail.com','$2a$10$m18JmedarHCzXa8uNytRcuyaqVITzfPF5sF9bBBFLdk1xITe1qMZ.','upload/1684159941445370106profile1.png','07536144283','Provider','Provider','update description','Pending'),
(83,'2023-05-15 14:28:39.486','2023-05-15 14:28:39.486',NULL,'sx1u22@soton.ac.uk','$2a$10$goVQnxqN7c9pAvaLj0Qr9OdrPkKGBU.KAtKyuuhUW/eoZQ..rnVk2','upload/profile.jpg','07899685746','IVY','Customer','','Pending');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
