package models

type Block struct {
  ID uint `gorm:"primaryKey"`
  Name string `gorm:"unique;not null"`
  Description string
  CourseID uint `gorm:"constraint:OnDelete:Cascade"`
  Course Course
}
