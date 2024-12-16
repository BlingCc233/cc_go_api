package models

type WhosSpy struct {
	ID    uint   `gorm:"primarykey" json:"id"`
	Word1 string `json:"word1"`
	Word2 string `json:"word2"`
}

func (WhosSpy) TableName() string {
	return "whos_spy"
}
