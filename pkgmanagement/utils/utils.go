package utils

import "github.com/jinzhu/gorm"

type Output interface {
	// Write takes in group of points to be written to the Output
	Write(points *[]*PKGContent) error
}

type Input interface {
	// Gather takes in an accumulator and adds the inputInfo to the Input
	Gather() (interface{}, error)
}

type PKGContent struct {
	gorm.Model
	Name     string `gorm:"unique;not null"`
	Parent   string
	Synopsis string
	Href     string
}
