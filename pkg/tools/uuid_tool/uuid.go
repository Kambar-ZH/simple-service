package uuid_tool

import (
	"github.com/google/uuid"
)

func GetRandomRequestID() string {
	return uuid.New().String()
}
