package models

import (
	"time"
)

type EducationStage struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:100;not null"`
	Description string `gorm:"size:300"`
}

type LeadSource struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100;not null"`
}

type UserRole struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100;not null"`
}

type User struct {
	ID          uint   `gorm:"primaryKey"`
	FirstName   string `gorm:"size:30;not null"`
	LastName    string `gorm:"size:30;not null"`
	Patronymic  string `gorm:"size:30;"`
	Description string `gorm:"size:300"`

	RoleID           *uint
	Role             UserRole `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	EducationStageID *uint
	EducationStage   EducationStage `gorm:"foreignKey:EducationStageID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	LeadSourceID     *uint
	LeadSource       LeadSource `gorm:"foreignKey:LeadSourceID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	Events []Event `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	Students []*User `gorm:"many2many:mentors_students;joinForeignKey:MentorID;joinReferences:StudentID"`
	Mentors  []*User `gorm:"many2many:mentors_students;joinForeignKey:StudentID;joinReferences:MentorID"`
}

type MentorStudent struct {
	MentorID  uint `gorm:"primaryKey"`
	StudentID uint `gorm:"primaryKey"`

	Mentor  User `gorm:"foreignKey:MentorID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Student User `gorm:"foreignKey:StudentID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

type Event struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null"`
	Date        time.Time `gorm:"not null"`
	Description string    `gorm:"size:300;not null"`
}
