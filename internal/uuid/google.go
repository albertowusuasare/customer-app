package uuid

import "github.com/google/uuid"

func Gen() GenFunc {
	return func() string {
		return uuid.New().String()
	}
}
