package util

import (
	"os"
	"testing"
)

func TestEnvVarString(t *testing.T) {
	var stage string
	os.Setenv("TEST", "thevar")
	EnvVarString(&stage, "TEST")
	if stage != "thevar" {
		t.Fatal(stage)
	}
}
