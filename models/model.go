package models

import "time"

type Member struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Handicap  float32   `json:"handicap"`
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TeeTime struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	MemberID  uint      `json:"memberId"`
	Time      time.Time `json:"time"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
