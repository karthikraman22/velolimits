package domain

import (
	"time"
)

type TxnEntry struct {
	id         uint      `gorm:"primaryKey,autoIncrement:true"`
	entity_id  string    `gorm:"index"`
	created_at time.Time `gorm:"index"`
	value      float64
}
