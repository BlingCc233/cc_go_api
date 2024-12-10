package models

type QCS struct {
	ID             uint   `gorm:"primarykey" json:"id"`
	Luck           string `json:"luck"`
	Description    string `json:"description"`
	Interpretation string `json:"interpretation"`
	Howsas         string `json:"howsas"`
	Advice         string `json:"advice"`
}
