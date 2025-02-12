package entity

/* CREATE TABLE `users` (
  `user_id` varchar(50) NOT NULL,
  `name` varchar(100) DEFAULT NULL,
  `dummy_col_1` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci; */

type Users struct {
	UserID    string `gorm:"column:user_id;type:varchar(50);primaryKey"`
	Name      string `gorm:"column:name;type:varchar(100)"`
	DummyCol1 string `gorm:"column:dummy_col_1;type:varchar(255)"`

	// has one Banner
	Banner *Banners `gorm:"foreignKey:UserID"`
}
