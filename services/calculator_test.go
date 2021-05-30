package services_test

import (
	"testing"

	"github.com/hyperyuri/test-go/services"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	// NATIVE
	if services.Sum(2, 2) != 4 {
		t.Error("2 + 2 must be 4")
	}

	// TESTIFY
	require.Equal(t, services.Sum(2, 2), 4)
}

func TestMult(t *testing.T) {
	// NATIVE
	if services.Sum(4, 4) != 4 {
		t.Error("2 + 2 must be 4")
	}

	// TESTIFY
	require.Equal(t, 4, services.Multiply(2, 2))
}
