package service

import (
	"github.com/google/uuid"
)

func GenerateHash(data string) string {
	return uuid.New().String()[:8]
}
