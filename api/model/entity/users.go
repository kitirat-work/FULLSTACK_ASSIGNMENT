package entity

type Users struct {
	UserID string `json:"userId" gorm:"column:user_id;type:varchar(50);primaryKey"`
	Name   string `json:"name" gorm:"column:name;type:varchar(100)"`

	// has many banners
	Banners []Banners `json:"banners" gorm:"foreignKey:UserID;references:UserID"`
	// has one user greetings
	UserGreetings UserGreetings `json:"userGreetings" gorm:"foreignKey:UserID;references:UserID"`
	// has many accounts
	Accounts []Accounts `json:"accounts" gorm:"foreignKey:UserID;references:UserID"`
	// has many debit cards
	DebitCards []DebitCards `json:"debitCards" gorm:"foreignKey:UserID;references:UserID"`
	// has many transactions
	Transactions []Transactions `json:"transactions" gorm:"foreignKey:UserID;references:UserID"`
}

func (Users) TableName() string {
	return "users"
}

type Banners struct {
	BannerID    string `json:"bannerId" gorm:"column:banner_id;type:varchar(50);primaryKey"`
	UserID      string `json:"userId" gorm:"column:user_id;type:varchar(50)"`
	Title       string `json:"title" gorm:"column:title;type:varchar(255)"`
	Description string `json:"description" gorm:"column:description;type:text"`
	Image       string `json:"image" gorm:"column:image;type:varchar(255)"`
}

func (Banners) TableName() string {
	return "banners"
}

type UserGreetings struct {
	UserID   string `json:"userId" gorm:"column:user_id;type:varchar(50);primaryKey"`
	Greeting string `json:"greeting" gorm:"column:greeting;type:text"`
}

func (UserGreetings) TableName() string {
	return "user_greetings"
}
