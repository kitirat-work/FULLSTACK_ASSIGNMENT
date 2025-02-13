package entity

type Transactions struct {
	TransactionID string `gorm:"column:transaction_id;type:varchar(50);primaryKey"`
	UserID        string `gorm:"column:user_id;type:varchar(50)"`
	Name          string `gorm:"column:name;type:varchar(100)"`
	Image         string `gorm:"column:image;type:varchar(255)"`
	IsBank        bool   `gorm:"column:isBank;type:tinyint(1)"`
}

func (Transactions) TableName() string {
	return "transactions"
}
