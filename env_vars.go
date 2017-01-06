package util

import (
	"fmt"
	"log"
	"os"
)

func EnvVarString(s *string, name string) {
	*s = os.Getenv(name)
	if *s == "" {
		log.Fatal(fmt.Sprintf("%s not set", name))
	}
}
