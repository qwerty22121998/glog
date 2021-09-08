package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestGenerateSlug(t *testing.T) {
	assert.Equal(t, "hello-world", GenerateSlug("Hello world"))
}
func TestGenerateSlugVietnamese(t *testing.T) {
	assert.Equal(t, "ho-khanh-vu", GenerateSlug("Hồ Khánh Vũ"))
}
