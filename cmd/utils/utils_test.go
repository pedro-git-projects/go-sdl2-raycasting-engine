package utils_test

import (
	"testing"

	"github.com/pedro-git-projects/go-raycasting/cmd/utils"
)

func TestNormalization(t *testing.T) {
	neg := -1.0
	utils.NormalizeAngle(&neg)
	if neg < 0 {
		t.Errorf("Expected positive result but got %f", neg)
	}
}
