package model

type Country struct {
	ID   uint   `gorm:"column:idx;primaryKey"`
	Name string `gorm:"column:name;size:255"`
}

func (c *Country) TableName() string {
	return "country"
}
