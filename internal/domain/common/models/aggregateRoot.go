package models

import "time"

type AggregateRoot[TId any] struct {
	id        TId
	createdAt time.Time
	updatedAt time.Time
}

func NewAggregateRoot[TId any](
	id TId,
) AggregateRoot[TId] {
	return AggregateRoot[TId]{id: id}
}

func (a *AggregateRoot[TId]) CreatedAt() time.Time {
	return a.createdAt
}

func (a *AggregateRoot[TId]) UpdatedAt() time.Time {
	return a.updatedAt
}

func (a *AggregateRoot[TId]) SetUpdatedAt(t time.Time) {
	a.updatedAt = t
}

func (a *AggregateRoot[TId]) Id() *TId {
	return &a.id
}
