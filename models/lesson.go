package models

type Lesson struct {
  ID  uint `gorm:"primaryKey"`
  Name string `gorm:"unique;not null"`
  Description string
  Materials string
  BlockID uint `gorm:"constraint:OnDelete:Cascade"`
  Block Block
}
