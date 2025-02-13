package entity

type Users struct {
	UserID string `gorm:"column:user_id;type:varchar(50);primaryKey"`
	Name   string `gorm:"column:name;type:varchar(100)"`

	// has many banners
	Banners []Banners `gorm:"foreignKey:UserID;references:UserID"`
	// has one user greetings
	UserGreetings UserGreetings `gorm:"foreignKey:UserID;references:UserID"`
}

func (Users) TableName() string {
	return "users"
}

type Banners struct {
	BannerID    string `gorm:"column:banner_id;type:varchar(50);primaryKey"`
	UserID      string `gorm:"column:user_id;type:varchar(50)"`
	Title       string `gorm:"column:title;type:varchar(255)"`
	Description string `gorm:"column:description;type:text"`
	Image       string `gorm:"column:image;type:varchar(255)"`
}

func (Banners) TableName() string {
	return "banners"
}

type UserGreetings struct {
	UserID   string `gorm:"column:user_id;type:varchar(50);primaryKey"`
	Greeting string `gorm:"column:greeting;type:text"`
}

func (UserGreetings) TableName() string {
	return "user_greetings"
}
