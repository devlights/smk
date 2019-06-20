package main

import (
	"testing"
)

func TestGetTemperatureSuccess(t *testing.T) {
	// Arrange
	cities := []string{
		"神戸",
	}

	for _, city := range cities {
		// Act
		temperature, err := GetTemperature(city)
		if err != nil {
			t.Error(err)
		}

		// Assert
		if temperature == "" {
			t.Errorf("temperature is empty [%s]\n", city)
		}
	}
}

func TestGetTemperatureCityIsEmpty(t *testing.T) {
	// Arrange
	city := ""

	// Act
	_, err := GetTemperature(city)

	// Assert
	if err == nil {
		t.Error("都市を空で指定したのにエラーが発生していない")
	}

	if err.Error() == "" {
		t.Error("エラーメッセージが取得できない")
	}
}

func TestGetTemperatureNotExistsCity(t *testing.T) {
	// Arrange
	cities := []string{
		"青森",
	}

	for _, city := range cities {
		// Act
		temperature, err := GetTemperature(city)
		if err != nil {
			t.Error(err)
		}

		// Assert
		if temperature != "" {
			t.Errorf("存在しない都市の気温が取れている [%s]\n", city)
		}
	}
}
