package entity

import (
	"time"
)

type AccountBalance struct {
	AccountID string  `gorm:"primaryKey;size:50"`
	UserID    string  `gorm:"size:50"`
	Amount    float64 `gorm:"type:decimal(15,2)"`
	DummyCol4 string  `gorm:"size:255"`
}

type AccountDetail struct {
	AccountID     string `gorm:"primaryKey;size:50"`
	UserID        string `gorm:"size:50"`
	Color         string `gorm:"size:10"`
	IsMainAccount bool
	Progress      int
	DummyCol5     string `gorm:"size:255"`
}

type AccountFlag struct {
	FlagID    int       `gorm:"primaryKey;autoIncrement"`
	AccountID string    `gorm:"size:50"`
	UserID    string    `gorm:"size:50"`
	FlagType  string    `gorm:"size:50"`
	FlagValue string    `gorm:"size:30"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Account struct {
	AccountID     string `gorm:"primaryKey;size:50"`
	UserID        string `gorm:"size:50"`
	Type          string `gorm:"size:50"`
	Currency      string `gorm:"size:10"`
	AccountNumber string `gorm:"size:20"`
	Issuer        string `gorm:"size:100"`
	DummyCol3     string `gorm:"size:255"`
}

type Banner struct {
	BannerID    string `gorm:"primaryKey;size:50"`
	UserID      string `gorm:"size:50"`
	Title       string `gorm:"size:255"`
	Description string `gorm:"type:text"`
	Image       string `gorm:"size:255"`
	DummyCol11  string `gorm:"size:255"`
}

type DebitCardDesign struct {
	CardID      string `gorm:"primaryKey;size:50"`
	UserID      string `gorm:"size:50"`
	Color       string `gorm:"size:10"`
	BorderColor string `gorm:"size:10"`
	DummyCol9   string `gorm:"size:255"`
}

type DebitCardDetail struct {
	CardID     string `gorm:"primaryKey;size:50"`
	UserID     string `gorm:"size:50"`
	Issuer     string `gorm:"size:100"`
	Number     string `gorm:"size:25"`
	DummyCol10 string `gorm:"size:255"`
}

type DebitCardStatus struct {
	CardID    string `gorm:"primaryKey;size:50"`
	UserID    string `gorm:"size:50"`
	Status    string `gorm:"size:20"`
	DummyCol8 string `gorm:"size:255"`
}

type DebitCard struct {
	CardID    string `gorm:"primaryKey;size:50"`
	UserID    string `gorm:"size:50"`
	Name      string `gorm:"size:100"`
	DummyCol7 string `gorm:"size:255"`
}

type Transaction struct {
	TransactionID string `gorm:"primaryKey;size:50"`
	UserID        string `gorm:"size:50"`
	Name          string `gorm:"size:100"`
	Image         string `gorm:"size:255"`
	IsBank        bool
	DummyCol6     string `gorm:"size:255"`
}

type UserGreeting struct {
	UserID    string `gorm:"primaryKey;size:50"`
	Greeting  string `gorm:"type:text"`
	DummyCol2 string `gorm:"size:255"`
}

type User struct {
	UserID    string `gorm:"primaryKey;size:50"`
	Name      string `gorm:"size:100"`
	DummyCol1 string `gorm:"size:255"`
}
