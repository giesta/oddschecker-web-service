package models

type Odd struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	BetId  uint   `json:"betId"`
	UserId string `json:"userId"`
	Odds   string `json:"odds"`
}
