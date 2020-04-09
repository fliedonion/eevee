// Code generated by eevee. DO NOT EDIT!

package model

import (
	"api/entity"
	"api/model"
)

func DefaultField() *model.Field {
	value := model.NewField(&entity.Field{
		Difficulty: uint64(0x0),
		ID:         uint64(0x0),
		Level:      uint64(0x0),
		LocationX:  uint64(0x0),
		LocationY:  uint64(0x0),
		Name:       "",
		ObjectNum:  uint64(0x0),
	}, nil)
	return value
}

func DefaultsField() *model.Fields {
	values := model.NewFields(entity.Fields{})
	values.Add(DefaultField())
	return values
}