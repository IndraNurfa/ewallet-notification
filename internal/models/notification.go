package models

import "time"

type NotificationTemplate struct {
	ID           int
	TemplateName string `gorm:"column:template_name;type:varchar(255);uniqueIndex"`
	Subject      string `gorm:"column:subject;type:varchar(255)"`
	Body         string `gorm:"column:body;type:varchar(255)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (*NotificationTemplate) TableName() string {
	return "notification_templates"
}

type NotificationHistory struct {
	ID           int
	Recipient    string `gorm:"column:recipient;type:varchar(255)"`
	TemplateID   int    `gorm:"column:template_id;type:int"`
	Status       string `gorm:"column:status;type:varchar(10)"`
	ErrorMessage string `gorm:"column:error_message;type:varchar(255)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (*NotificationHistory) TableName() string {
	return "notification_histories"
}
