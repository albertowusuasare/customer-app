package uuid

import "github.com/google/uuid"

func Gen() GenFunc {
	return func() string {
		return uuid.New().String()
	}
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
