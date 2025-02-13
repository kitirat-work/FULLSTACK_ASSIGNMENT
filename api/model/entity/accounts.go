package entity

type Accounts struct {
	AccountID     string `gorm:"column:account_id;type:varchar(50);primaryKey"`
	UserID        string `gorm:"column:user_id;type:varchar(50)"`
	Type          string `gorm:"column:type;type:varchar(50)"`
	Currency      string `gorm:"column:currency;type:varchar(10)"`
	AccountNumber string `gorm:"column:account_number;type:varchar(20)"`
	Issuer        string `gorm:"column:issuer;type:varchar(100)"`
	DummyCol3     string `gorm:"column:dummy_col_3;type:varchar(255)"`

	// has one AccountBalances
	AccountBalances AccountBalances `gorm:"foreignKey:AccountID;references:AccountID"`
	// has one AccountDetails
	AccountDetails AccountDetails `gorm:"foreignKey:AccountID;references:AccountID"`
	// has many AccountFlags
	AccountFlags []AccountFlags `gorm:"foreignKey:AccountID;references:AccountID"`
}

func (Accounts) TableName() string {
	return "accounts"
}

type AccountBalances struct {
	AccountID string  `gorm:"column:account_id;type:varchar(50);primaryKey"`
	UserID    string  `gorm:"column:user_id;type:varchar(50)"`
	Amount    float64 `gorm:"column:amount;type:decimal(15,2)"`
	DummyCol4 string  `gorm:"column:dummy_col_4;type:varchar(255)"`
}

func (AccountBalances) TableName() string {
	return "account_balances"
}

type AccountDetails struct {
	AccountID     string `gorm:"column:account_id;type:varchar(50);primaryKey"`
	UserID        string `gorm:"column:user_id;type:varchar(50)"`
	Color         string `gorm:"column:color;type:varchar(10)"`
	IsMainAccount bool   `gorm:"column:is_main_account;type:tinyint(1)"`
	Progress      int    `gorm:"column:progress;type:int"`
	DummyCol5     string `gorm:"column:dummy_col_5;type:varchar(255)"`
}

func (AccountDetails) TableName() string {
	return "account_details"
}

type AccountFlags struct {
	FlagID    int    `gorm:"column:flag_id;type:int;primaryKey;autoIncrement"`
	AccountID string `gorm:"column:account_id;type:varchar(50)"`
	UserID    string `gorm:"column:user_id;type:varchar(50)"`
	FlagType  string `gorm:"column:flag_type;type:varchar(50)"`
	FlagValue string `gorm:"column:flag_value;type:varchar(30)"`
	CreatedAt string `gorm:"column:created_at;type:timestamp"`
	UpdatedAt string `gorm:"column:updated_at;type:timestamp"`
}

func (AccountFlags) TableName() string {
	return "account_flags"
}
