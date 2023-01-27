package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Weather struct {
	Weather string
	Temp    int
	Rain    int
}

func main() {
	resp, err := http.Get("https://weather.tsukumijima.net/api/forecast/city/130010")
	if err != nil {
	} else {
		defer resp.Body.Close()
		byteArray, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(byteArray))
		// JSONをパースする
		var weather Weather
		if err := json.Unmarshal(byteArray, &weather); err != nil {
		} else {
			fmt.Printf("今日の天気は%sです。\n", weather.Weather)
			fmt.Printf("今日の気温は%d℃です。\n", weather.Temp)
			fmt.Printf("今日の降水確率は%d%%です。\n", weather.Rain)
		}
	}
}
