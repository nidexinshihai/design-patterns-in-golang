package main

import (
	"fmt"
)

type WeatherObserver interface {
	Update() error
}

type CurrentWeather struct {
	subject WeatherSubject
}

func (cw *CurrentWeather) Update() error {
	cw.display()
	return nil
}

func (cw *CurrentWeather) display() {
	fmt.Printf("\nCurrent temperature: %.1f C, ", cw.subject.GetTemperature())
	fmt.Printf("Current Humidity: %.1f%%\n", cw.subject.GetHumidity())
}

type StatisticWeather struct {
	subject      WeatherSubject
	temperatures []float32
	humidities   []float32
}

const (
	STATISTIC_PERIOD = 10
)

func (sw *StatisticWeather) Update() error {
	if len(sw.temperatures) >= STATISTIC_PERIOD {
		sw.temperatures = sw.temperatures[1:]
	}
	sw.temperatures = append(sw.temperatures, sw.subject.GetTemperature())

	if len(sw.humidities) >= STATISTIC_PERIOD {
		sw.humidities = sw.humidities[1:]
	}
	sw.humidities = append(sw.humidities, sw.subject.GetHumidity())

	sw.display()
	return nil
}

func (sw *StatisticWeather) display() {
	var sumTemperature float32
	var sumHumidity float32
	for _, temperature := range sw.temperatures {
		sumTemperature += temperature
	}

	for _, humidity := range sw.humidities {
		sumHumidity += humidity
	}

	fmt.Printf("\nAverage temperature: %.1f C, ", sumTemperature/float32(len(sw.temperatures)))
	fmt.Printf("Average Humidity: %.1f%%\n", sumHumidity/float32(len(sw.humidities)))
}
