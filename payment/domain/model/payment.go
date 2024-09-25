package model

type Payment struct {
	ID            int64  `gorm:"primary_key;not_null;auto_increment"`
	PaymentName   string `json:"payment_name"`
	PaymentSid    string `json:"payment_sid"`
	PaymentStatus bool   `json:"payment_status"`
	PaymentImage  string `json:"payment_image"`
}
