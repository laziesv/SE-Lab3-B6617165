package entity

import (

	"gorm.io/gorm"
)

type Proposal struct {
	gorm.Model
	Title         string         `gorm:"type:varchar(255);not null" json:"title"`              // ชื่อข้อเสนอ
	Description   string         `gorm:"type:text" json:"description"`                         // รายละเอียดข้อเสนอ
	ProposerID    uint           `json:"proposer_id"`                                          // ผู้เสนอ (FK ไปยัง User)
	Status        string         `gorm:"type:varchar(50);default:'pending'" json:"status"`     // สถานะ เช่น pending, approved, rejected
	Budget        float64        `gorm:"type:decimal(10,2)" json:"budget"`                     // งบประมาณที่เสนอ
	AttachmentURL string         `gorm:"type:varchar(255)" json:"attachment_url"`              // ไฟล์แนบ (เช่น เอกสาร PDF)
}