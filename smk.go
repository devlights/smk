package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var url = "https://www.data.jma.go.jp/obd/stats/data/mdrr/tem_rct/alltable/mxtemsadext00.html"

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

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	temperature = ""

	doc.Find("tr.mtx").Each(func(i int, s *goquery.Selection) {
		cityElem := s.ChildrenFiltered("td:nth-child(3)")
		temperatureElem := s.ChildrenFiltered("td:nth-child(4)")

		if strings.Contains(cityElem.Text(), city) {
			temperature = temperatureElem.Text()
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

	if temperature == "" {
		fmt.Println("見つかりませんでした...")
		return
	}

	fmt.Println(city, temperature)
}
