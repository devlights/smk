package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

var url = "https://www.jma.go.jp/jp/yoho/332.html"

// CityIsEmptyError - パラメータ[city]が空の場合のエラー
type CityIsEmptyError struct{}

func (c CityIsEmptyError) Error() string {
	return "city is empty."
}

// GetTemperature - 指定された都市の気温を取得します.
// @params
//   city - 都市名
// @returns
//   temperature - 気温. 取得できなかった場合は空文字.
//   err - エラー
func GetTemperature(city string) (temperature string, err error) {
	if city == "" {
		return "", CityIsEmptyError{}
	}

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	temperature = ""

	doc.Find(".city").Each(func(i int, s *goquery.Selection) {
		if s.Text() == city {
			temperature = s.SiblingsFiltered(".max").Text()
		}
	})

	return temperature, nil
}

func main() {
	city := "神戸"
	temperature, err := GetTemperature(city)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(city, temperature)
}
