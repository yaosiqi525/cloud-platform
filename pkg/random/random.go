package random

import (
	"github.com/google/uuid"
)

func GetRandomString() string {
	return uuid.NewString()
}
