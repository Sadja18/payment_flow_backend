package models

import "time"

type InvoiceStatus string

const (
    StatusPaid    InvoiceStatus = "paid"
    StatusUnpaid  InvoiceStatus = "unpaid"
    StatusOverdue InvoiceStatus = "past due"
)

type Invoice struct {
    ID          uint           `gorm:"primaryKey"`
    InvoiceUUID string         `gorm:"uniqueIndex;not null"`
    PayerUserID uint           `gorm:"not null"`
    PayeeUserID uint           `gorm:"not null"`
    Amount      float64        `gorm:"not null"`
    Currency    string         `gorm:"not null"`
    InvoiceDate time.Time      `gorm:"not null"`
    DueDate     time.Time      `gorm:"not null"`
    Status      InvoiceStatus  `gorm:"default:'unpaid'"`
}
