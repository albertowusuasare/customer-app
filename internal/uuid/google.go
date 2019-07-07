package uuid

import "github.com/google/uuid"

// Gen returns a google implementation of uuid generation
func Gen() GenFunc {
	return func() string {
		return uuid.New().String()
	}
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
