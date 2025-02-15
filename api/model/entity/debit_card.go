package entity

type DebitCards struct {
	CardID string `json:"card_id" gorm:"column:card_id;type:varchar(50);primaryKey"`
	UserID string `json:"user_id" gorm:"column:user_id;type:varchar(50)"`
	Name   string `json:"name" gorm:"column:name;type:varchar(100)"`

	// has one DebitCardDesign
	DebitCardDesign DebitCardDesign `json:"debit_card_design" gorm:"foreignKey:CardID;references:CardID"`
	// has one DebitCardDetails
	DebitCardDetails DebitCardDetails `json:"debit_card_details" gorm:"foreignKey:CardID;references:CardID"`
	// has one DebitCardStatus
	DebitCardStatus DebitCardStatus `json:"debit_card_status" gorm:"foreignKey:CardID;references:CardID"`
}

func (DebitCards) TableName() string {
	return "debit_cards"
}

type DebitCardDesign struct {
	CardID      string `json:"card_id" gorm:"column:card_id;type:varchar(50);primaryKey"`
	UserID      string `json:"user_id" gorm:"column:user_id;type:varchar(50)"`
	Color       string `json:"color" gorm:"column:color;type:varchar(10)"`
	BorderColor string `json:"border_color" gorm:"column:border_color;type:varchar(10)"`
}

func (DebitCardDesign) TableName() string {
	return "debit_card_design"
}

type DebitCardDetails struct {
	CardID string `json:"card_id" gorm:"column:card_id;type:varchar(50);primaryKey"`
	UserID string `json:"user_id" gorm:"column:user_id;type:varchar(50)"`
	Issuer string `json:"issuer" gorm:"column:issuer;type:varchar(100)"`
	Number string `json:"number" gorm:"column:number;type:varchar(25)"`
}

func (DebitCardDetails) TableName() string {
	return "debit_card_details"
}

type DebitCardStatus struct {
	CardID string `json:"card_id" gorm:"column:card_id;type:varchar(50);primaryKey"`
	UserID string `json:"user_id" gorm:"column:user_id;type:varchar(50)"`
	Status string `json:"status" gorm:"column:status;type:varchar(20)"`
}
