package models

type Odd struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	BetId  uint   `json:"bet_id" sql:"type:integer REFERENCES bets(id)"`
	UserId string `json:"user_id" sql:"type:text REFERENCES users(id)"`
	Odds   string `json:"odds"`
}
