package setup

import (
	"fmt"
	"os"
)

type KEY struct {
	api string
}

func key(key string) (string, error) {
	if val, ok := os.LookupEnv(key); ok {
		return val, nil
	}

	return "", fmt.Errorf("os.LookupEnv: API key not assigned: %s\n", key)
}

func NewK() *KEY {
	key, err := key("POLYGON_ACCESS_KEY")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	k := &KEY{key}

	return k
}

func (k *KEY) Key() string {
	return k.api
}
