package entity

/* CREATE TABLE `banners` (
  `banner_id` varchar(50) NOT NULL,
  `user_id` varchar(50) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `description` text,
  `image` varchar(255) DEFAULT NULL,
  `dummy_col_11` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`banner_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci; */

type Banners struct {
	BannerID    string `gorm:"column:banner_id;type:varchar(50);primaryKey"`
	UserID      string `gorm:"column:user_id;type:varchar(50)"`
	Title       string `gorm:"column:title;type:varchar(255)"`
	Description string `gorm:"column:description;type:text"`
	Image       string `gorm:"column:image;type:varchar(255)"`
	DummyCol11  string `gorm:"column:dummy_col_11;type:varchar(255)"`
}
