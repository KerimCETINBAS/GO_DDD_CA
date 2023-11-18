package models_test

import (
	"testing"
	"time"

	. "gihub.com/kerimcetinbas/go_ddd_ca/domain/common/models"
	"github.com/stretchr/testify/assert"
)

func TestAggregateRoot(t *testing.T) {
	a := NewAggregateRoot[string]("id")
	t.Run("should create aggrageRoot", func(t *testing.T) {

		assert.NotNil(t, a)
	})

	t.Run("should get properties", func(t *testing.T) {
		assert.Equal(t, *a.Id(), "id")
		assert.True(t, time.Now().After(a.CreatedAt()))
	})

	t.Run("UpdatedAt must be mutable", func(t *testing.T) {

		updatedBefore := a.UpdatedAt()

		a.SetUpdatedAt(time.Now())

		assert.NotEqual(t, updatedBefore, a.UpdatedAt())
	})
}
