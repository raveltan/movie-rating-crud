# Movie Rater

A movie rater CRUD Application written in GO

> The frontend portion of the app is not yet optimized for production, use at your own discretion!

## Setup

Before running the webserver, make sure that you have the created user john or you can follow the following steps:
As an admin, run:

```sql
create user 'john'@'localhost' IDENTIFIED BY 'John@1234';
grant all privileges on *.* to 'john'@'localhost';
```

Then run the migration sql from data.sql or run the following:

```sql
DROP TABLE IF EXISTS `movie`;

CREATE TABLE `movie` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `rating` decimal(5,1) NOT NULL DEFAULT '0.0',
  `voter` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `review`;

CREATE TABLE `review` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(60) NOT NULL,
  `review` varchar(200) NOT NULL,
  `rating` int NOT NULL,
  `movie` int NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`movie`) REFERENCES movie(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(60) NOT NULL,
  `password` varchar(60) NOT NULL,
  `firstName` varchar(60) NOT NULL,
  `lastName` varchar(60) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

```

Run each the executable according the system.
