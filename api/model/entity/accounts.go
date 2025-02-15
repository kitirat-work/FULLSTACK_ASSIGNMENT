package entity

type Accounts struct {
	AccountID     string `json:"accountId" gorm:"column:account_id;type:varchar(50);primaryKey"`
	UserID        string `json:"userId" gorm:"column:user_id;type:varchar(50)"`
	Type          string `json:"type" gorm:"column:type;type:varchar(50)"`
	Currency      string `json:"currency" gorm:"column:currency;type:varchar(10)"`
	AccountNumber string `json:"accountNumber" gorm:"column:account_number;type:varchar(20)"`
	Issuer        string `json:"issuer" gorm:"column:issuer;type:varchar(100)"`
	DummyCol3     string `json:"dummyCol3" gorm:"column:dummy_col_3;type:varchar(255)"`

	// has one AccountBalances
	AccountBalances AccountBalances `json:"accountBalances" gorm:"foreignKey:AccountID;references:AccountID"`
	// has one AccountDetails
	AccountDetails AccountDetails `json:"accountDetails" gorm:"foreignKey:AccountID;references:AccountID"`
	// has many AccountFlags
	AccountFlags []AccountFlags `json:"accountFlags" gorm:"foreignKey:AccountID;references:AccountID"`
}

func (Accounts) TableName() string {
	return "accounts"
}

type AccountBalances struct {
	AccountID string  `json:"accountId" gorm:"column:account_id;type:varchar(50);primaryKey"`
	UserID    string  `json:"userId" gorm:"column:user_id;type:varchar(50)"`
	Amount    float64 `json:"amount" gorm:"column:amount;type:decimal(15,2)"`
	DummyCol4 string  `json:"dummyCol4" gorm:"column:dummy_col_4;type:varchar(255)"`
}

func (AccountBalances) TableName() string {
	return "account_balances"
}

type AccountDetails struct {
	AccountID     string `json:"accountId" gorm:"column:account_id;type:varchar(50);primaryKey"`
	UserID        string `json:"userId" gorm:"column:user_id;type:varchar(50)"`
	Color         string `json:"color" gorm:"column:color;type:varchar(10)"`
	IsMainAccount bool   `json:"isMainAccount" gorm:"column:is_main_account;type:tinyint(1)"`
	Progress      int    `json:"progress" gorm:"column:progress;type:int"`
	DummyCol5     string `json:"dummyCol5" gorm:"column:dummy_col_5;type:varchar(255)"`
}

func (AccountDetails) TableName() string {
	return "account_details"
}

type AccountFlags struct {
	FlagID    int    `json:"flagId" gorm:"column:flag_id;type:int;primaryKey;autoIncrement"`
	AccountID string `json:"accountId" gorm:"column:account_id;type:varchar(50)"`
	UserID    string `json:"userId" gorm:"column:user_id;type:varchar(50)"`
	FlagType  string `json:"flagType" gorm:"column:flag_type;type:varchar(50)"`
	FlagValue string `json:"flagValue" gorm:"column:flag_value;type:varchar(30)"`
	CreatedAt string `json:"createdAt" gorm:"column:created_at;type:timestamp"`
	UpdatedAt string `json:"updatedAt" gorm:"column:updated_at;type:timestamp"`
}

func (AccountFlags) TableName() string {
	return "account_flags"
}
