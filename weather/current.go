package weather

import (
	"encoding/json"
	"fmt"
)

type CurrentWeatherData struct {
	Coord      Coordinates `json:"coord"`
	Weather    []Weather   `json:"weather"`
	Base       string      `json:"base"`
	Main       Main        `json:"main"`
	Visibility int         `json:"visibility"`
	Wind       Wind        `json:"wind"`
	Clouds     Clouds      `json:"clouds"`
	Dt         int         `json:"dt"`
	Sys        Sys         `json:"sys"`
	Timezone   int         `json:"timezone"`
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Cod        int         `json:"cod"`
	Key        string
	*Settings
}

type RequiredData struct {
	Name        string
	Zipcode     int
	Longitude   float64 `json:"lon"`
	Latitude    float64 `json:"lat"`
	Visibility  int
	Temperature float64
	Pressure    float64
	Humidity    int
	Country     string
}

func NewCurrent(key string ) (*CurrentWeatherData, error) {
	c := &CurrentWeatherData{
		Settings: NewSettings(),
	}
	c.Key = key
	return c, nil
}

func (w *CurrentWeatherData) CurrentByName(location string) (RequiredData, error) {
	url := fmt.Sprintf(fmt.Sprintf(baseURL, "appid=%s&q=%s"), w.Key, location)
	response, err := w.client.Get(url)
	if err != nil {
		return RequiredData{}, err
	}
	defer response.Body.Close()
	var currentObj CurrentWeatherData
	if err := json.NewDecoder(response.Body).Decode(&currentObj); err != nil {
		return RequiredData{}, err
	}
	data := RequiredData{
		Name:        currentObj.Name,
		Zipcode:     0,
		Longitude:   currentObj.Coord.Longitude,
		Latitude:    currentObj.Coord.Latitude,
		Visibility:  currentObj.Visibility,
		Temperature: currentObj.Main.Temp,
		Pressure:    currentObj.Main.Pressure,
		Humidity:    currentObj.Main.Humidity,
		Country:     currentObj.Sys.Country,
	}
	return data, nil
}
