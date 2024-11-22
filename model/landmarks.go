package model

type Landmark struct {
	ID        uint    `gorm:"column:idx;primaryKey"`
	Name      string  `gorm:"column:name;size:255"`
	Detail    string  `gorm:"column:detail;size:2000"`
	Url       string  `gorm:"column:url;size:1000"`
	CountryID uint    `gorm:"column:country"`
	Country   Country `gorm:"foreignkey:CountryID"`
}

func (l *Landmark) TableName() string {
	return "landmark"
}
