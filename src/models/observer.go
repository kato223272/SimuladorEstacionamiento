package models

type Observer interface {
	Update(data interface{})
}
