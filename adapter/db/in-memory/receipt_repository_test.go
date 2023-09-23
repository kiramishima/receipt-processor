package in_memory

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReceiptRepository_SaveReceiptPoints(t *testing.T) {
	repo := NewReceiptRepository()

	t.Run("OK", func(t *testing.T) {
		id, err := repo.SaveReceiptPoints(6)
		assert.NoError(t, err)
		assert.NotEmpty(t, id)
	})

	t.Run("Multiple Times", func(t *testing.T) {

		var points = []int16{0, 10, 12, 20, 120}
		for _, point := range points {
			id, err := repo.SaveReceiptPoints(point)
			assert.NoError(t, err)
			assert.NotEmpty(t, id)
		}

		assert.Equal(t, len(repo.records), len(points)+1) // +1 because last execution add 1 record
	})
}

func TestReceiptRepository_FindReceiptById(t *testing.T) {
	repo := NewReceiptRepository()

	// Generate some entries
	var points = []int16{0, 10, 12, 20, 120}
	var id string
	for _, point := range points {
		id, _ = repo.SaveReceiptPoints(point)
	}

	t.Run("OK", func(t *testing.T) {
		item, err := repo.FindReceiptById(id)
		assert.NoError(t, err)
		assert.NotEmpty(t, item)
	})

	t.Run("Not exists", func(t *testing.T) {
		var uid = uuid.New().String()
		item, err := repo.FindReceiptById(uid)
		assert.Error(t, err)
		assert.Nil(t, item)
	})
}
