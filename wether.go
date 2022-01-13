package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var place string = "Noida"

func WeatherFunc() {

	window := fyne.CurrentApp().NewWindow("Weather")
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(500, 400))
	img, _ := fyne.LoadResourceFromPath("./img/weather.png")
	window.SetIcon(img)
	window.SetPadded(false)

	img2 := canvas.NewImageFromFile("./img/weather.png")
	img2.Resize(fyne.NewSize(500, 400))

	res, err := http.Get(fmt.Sprint("http://api.openweathermap.org/data/2.5/weather?q=", place, "&appid=03e1063ba06ab6430556eb0e2fc3997e"))
	if err != nil {
		fmt.Print(err)
	}
	selectRegion := widget.NewSelectEntry([]string{"peshawar", "multan", "gujrat", "kabul", "dehli"})
	selectRegion.OnSubmitted = func(s string) {
		place = s
		locationChange(window)

	}
	selectRegion.Resize(fyne.NewSize(200, 10))
	selectRegion.PlaceHolder = place
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
	}

	weathers, err := UnmarshalWeather(body)
	if err != nil {
		fmt.Print(err)
	}

	temp := canvas.NewText(fmt.Sprint(int32((weathers.Main.Temp)-273.15)), color.Black)
	temp.TextSize = 90
	temp.TextStyle = fyne.TextStyle{Bold: true}

	label_Celsius := canvas.NewText(fmt.Sprint(weathers.Sys.Country, ", ", weathers.Name, ", °C"), color.Black)
	log.Println(weathers.Wind.Speed)
	label_Wind_Speed := canvas.NewText(fmt.Sprint("Wind Speed ", int64(weathers.Wind.Deg), "°"), color.Black)

	window.SetContent(container.NewBorder(nil, nil, nil, nil, img2, container.NewVBox(selectRegion, container.NewCenter(temp),
		container.NewCenter(label_Celsius),
		container.NewCenter(label_Wind_Speed),
	)))
	window.Show()
}

func locationChange(window fyne.Window) {

	img2 := canvas.NewImageFromFile("./img/weather.png")

	img2.Resize(fyne.NewSize(500, 400))

	res, err := http.Get(fmt.Sprint("http://api.openweathermap.org/data/2.5/weather?q=", place, "&appid=03e1063ba06ab6430556eb0e2fc3997e"))
	if err != nil {
		fmt.Print(err)
	}

	selectRegion := widget.NewSelectEntry([]string{"New York", "Berlin", "Dubai", "kabul", "Dehli"})
	selectRegion.OnSubmitted = func(s string) {
		place = s
		locationChange(window)

	}
	selectRegion.Resize(fyne.NewSize(200, 10))
	selectRegion.PlaceHolder = place

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
	}

	weathers, err := UnmarshalWeather(body)
	if err != nil {
		fmt.Print(err)
	}

	temp := canvas.NewText(fmt.Sprint(int32((weathers.Main.Temp)-273.15)), color.Black)
	temp.TextSize = 90
	temp.TextStyle = fyne.TextStyle{Bold: true}

	label_Celsius := canvas.NewText(fmt.Sprint(weathers.Sys.Country, ", ", weathers.Name, ", °C"), color.Black)
	log.Println(weathers.Wind.Speed)
	label_Wind_Speed := canvas.NewText(fmt.Sprint("Wind Speed ", int64(weathers.Wind.Deg), "°"), color.Black)

	window.SetContent(container.NewBorder(nil, nil, nil, nil, img2, container.NewVBox(selectRegion, container.NewCenter(temp),
		container.NewCenter(label_Celsius),
		container.NewCenter(label_Wind_Speed),
	)))
}

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`
	Weather    []WeatherElement `json:"weather"`
	Base       string           `json:"base"`
	Main       Main             `json:"main"`
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`
	Rain       Rain             `json:"rain"`
	Clouds     Clouds           `json:"clouds"`
	Dt         int64            `json:"dt"`
	Sys        Sys              `json:"sys"`
	Timezone   int64            `json:"timezone"`
	ID         int64            `json:"id"`
	Name       string           `json:"name"`
	Cod        int64            `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Rain struct {
	The1H float64 `json:"1h"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type WeatherElement struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed int64 `json:"speed"`
	Deg   int64 `json:"deg"`
}
