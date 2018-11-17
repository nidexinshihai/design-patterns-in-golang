package main

func main() {

	weatherData := new(WeatherData)

	currentWeather := CurrentWeather{weatherData}
	statisticWeather := StatisticWeather{weatherData, nil, nil}

	weatherData.RegisterObserver(&currentWeather)
	weatherData.RegisterObserver(&statisticWeather)

	weatherData.SetData(26, 60)
	weatherData.SetData(27, 70)
	weatherData.SetData(28, 80)

	weatherData.DeregisterObserver(&statisticWeather)

	weatherData.SetData(29, 90)
}
