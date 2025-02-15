package entity

type Transactions struct {
	TransactionID string `json:"transactionId" gorm:"column:transaction_id;type:varchar(50);primaryKey"`
	UserID        string `json:"userId" gorm:"column:user_id;type:varchar(50)"`
	Name          string `json:"name" gorm:"column:name;type:varchar(100)"`
	Image         string `json:"image" gorm:"column:image;type:varchar(255)"`
	IsBank        bool   `json:"isBank" gorm:"column:isBank;type:tinyint(1)"`
}

func (Transactions) TableName() string {
	return "transactions"
}
