package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	ID           uuid.UUID      `json:"id" gorm:"type:varchar(36);primary_key"`
	FullName     string         `json:"full_name" gorm:"size:255"`
	Email        string         `json:"email" gorm:"size:255;unique"`
	PhoneNumber  string         `json:"phone_number" gorm:"size:255;unique"`
	Password     string         `json:"password" gorm:"size:255"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Department struct {
	ID             uuid.UUID      `json:"id" gorm:"type:varchar(36);primary_key"`
	DepartmentName string         `json:"department_name" gorm:"type:varchar(255)"`
	CreatedAt      time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Courses        []Course       `json:"courses" gorm:"foreignKey:DepartmentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ImageID        *uuid.UUID     `json:"image_id" gorm:"type:varchar(36);default:NULL"`
	Image          *Image         `gorm:"foreignKey:ImageID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Course struct {
	ID           uuid.UUID      `json:"id" gorm:"type:varchar(36);primary_key"`
	Level        string         `json:"level" gorm:"size:255"`
	CourseName   string         `json:"course_name" gorm:"type:varchar(255)"`
	KcseGrade    string         `json:"kcse_grade" gorm:"type:varchar(255)"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	DepartmentID uuid.UUID      `json:"department_id" gorm:"type:varchar(36);default:NULL"`
	Department   Department     `json:"-" gorm:"foreignKey:DepartmentID;references:ID"`
}

type Job struct {
	ID              uuid.UUID      `json:"id" gorm:"type:varchar(36);primary_key"`
	Title           string         `json:"title" gorm:"size:255"`
	Role            string         `json:"role" gorm:"type:varchar(255)"`
	ApplicationLink string         `json:"application_link" gorm:"type:varchar(1024)"`
	CreatedAt       time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Sponsor struct {
	ID              uuid.UUID      `json:"id" gorm:"type:varchar(36);primary_key"`
	Name            string         `json:"name" gorm:"type:varchar(255)"`
	ApplicationLink string         `json:"application_link" gorm:"type:varchar(1024)"`
	Description     string         `json:"description" gorm:"type:text"`
	CreatedAt       time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Image struct {
	ID        uuid.UUID      `json:"id" gorm:"type:varchar(36);primary_key"`
	Path      string         `json:"path" gorm:"type:varchar(255)"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
