package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	core "k8s.io/api/core/v1"
)

func TestKafkaReconcilerFns(t *testing.T) {

}

func TestConverterFuncs(t *testing.T) {
	t.Run("Test intMin", func(t *testing.T) {
		answer, _ := IntMin([]string{"4", "6", "7"})
		if answer != "4" {
			t.Errorf("Min function should have returned 4, returned %s", answer)
		}
	})
	t.Run("Test intMax", func(t *testing.T) {
		answer, _ := IntMax([]string{"4", "6", "7"})
		if answer != "7" {
			t.Errorf("Min function should have returned 7, returned %s", answer)
		}
	})
	t.Run("Test ListMerge", func(t *testing.T) {
		answer, _ := ListMerge([]string{"4,5,6", "6", "7,2"})
		if answer != "2,4,5,6,7" {
			t.Errorf("Min function should have returned 2,4,5,6,7 returned %s", answer)
		}
	})
}

func TestBase64Decode(t *testing.T) {
	s := core.Secret{
		Data: map[string][]byte{
			"key": []byte("bnVtYmVy"),
		},
	}
	decodedValue, _ := B64Decode(&s, "key")
	assert.Equal(t, decodedValue, "number", "should decode the right value")
}

func TestRandString(t *testing.T) {
	a := RandString(12)
	b := RandString(12)
	assert.NotEqual(t, a, b)
}
