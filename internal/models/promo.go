package models

import "gorm.io/gorm"

type PromoBucket struct {
	gorm.Model

	BucketName   string
	BucketKeys   int
	BucketAuthor string

	Type int
}

type PromoCode struct {
	gorm.Model

	BucketID uint

	Code           string `gorm:"unique"`
	Creator        string
	Activations    int
	MaxActivations int
	Time           int
}

type PromoActivations struct {
	gorm.Model

	CodeID    uint
	Activator string
}
