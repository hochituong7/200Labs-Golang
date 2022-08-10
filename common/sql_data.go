package common

import "time"

type SQLModel struct {
	Id        int        `json:"id" gorm:"column:id"`
	Status    int        `json:"status" gorm:"column:status"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (sqlModel *SQLModel) PrepareForInsert() {
	now := time.Now().UTC()
	sqlModel.Id = 0
	sqlModel.Status = 1
	sqlModel.CreatedAt = &now
	sqlModel.UpdatedAt = &now
}