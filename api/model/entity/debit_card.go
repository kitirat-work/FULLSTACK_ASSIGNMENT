package entity

type DebitCards struct {
	CardID string `gorm:"column:card_id;type:varchar(50);primaryKey"`
	UserID string `gorm:"column:user_id;type:varchar(50)"`
	Name   string `gorm:"column:name;type:varchar(100)"`

	// has one DebitCardDesign
	DebitCardDesign DebitCardDesign `gorm:"foreignKey:CardID;references:CardID"`
	// has one DebitCardDetails
	DebitCardDetails DebitCardDetails `gorm:"foreignKey:CardID;references:CardID"`
	// has one DebitCardStatus
	DebitCardStatus DebitCardStatus `gorm:"foreignKey:CardID;references:CardID"`
}

func (DebitCards) TableName() string {
	return "debit_cards"
}

type DebitCardDesign struct {
	CardID      string `gorm:"column:card_id;type:varchar(50);primaryKey"`
	UserID      string `gorm:"column:user_id;type:varchar(50)"`
	Color       string `gorm:"column:color;type:varchar(10)"`
	BorderColor string `gorm:"column:border_color;type:varchar(10)"`
}

func (DebitCardDesign) TableName() string {
	return "debit_card_design"
}

type DebitCardDetails struct {
	CardID string `gorm:"column:card_id;type:varchar(50);primaryKey"`
	UserID string `gorm:"column:user_id;type:varchar(50)"`
	Issuer string `gorm:"column:issuer;type:varchar(100)"`
	Number string `gorm:"column:number;type:varchar(25)"`
}

func (DebitCardDetails) TableName() string {
	return "debit_card_details"
}

type DebitCardStatus struct {
	CardID string `gorm:"column:card_id;type:varchar(50);primaryKey"`
	UserID string `gorm:"column:user_id;type:varchar(50)"`
	Status string `gorm:"column:status;type:varchar(20)"`
}
