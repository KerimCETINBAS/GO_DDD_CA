package models

type Entity[TId comparable] struct {
	Id TId
}

func (e *Entity[TId]) IsEqual(d Entity[TId]) bool {
	return e.Id == d.Id
}
