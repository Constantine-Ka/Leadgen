package config

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {

	_, err := os.ReadFile("./../../config.yaml")
	if err != nil {
		t.Error("Файла конфигурации нет", err)
	}

}
