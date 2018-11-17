package main

import (
	"fmt"
)

type WeatherSubject interface {
	RegisterObserver(observer WeatherObserver) error
	DeregisterObserver(observer WeatherObserver) error

	GetTemperature() float32
	GetHumidity() float32
}

type WeatherData struct {
	temperature float32
	humidity    float32
	observers   []WeatherObserver
}

func (data *WeatherData) RegisterObserver(observer WeatherObserver) error {

	data.observers = append(data.observers, observer)
	return nil
}

func (data *WeatherData) DeregisterObserver(observer WeatherObserver) error {

	for idx, o := range data.observers {
		if o == observer {
			data.observers = append(data.observers[:idx], data.observers[idx+1:]...)
			return nil
		}
	}

	return fmt.Errorf("observer not found")
}

func (data *WeatherData) GetTemperature() float32 {
	return data.temperature
}

func (data *WeatherData) SetTemperature(temperature float32) error {

	data.temperature = temperature
	return nil
}

func (data *WeatherData) GetHumidity() float32 {
	return data.humidity
}

func (data *WeatherData) SetHumidity(humidity float32) error {

	data.humidity = humidity
	return nil
}

func (data *WeatherData) SetData(temperature float32, humidity float32) error {

	if err := data.SetTemperature(temperature); err != nil {
		return err
	}

	if err := data.SetHumidity(humidity); err != nil {
		return err
	}

	return data.notifyObservers()
}

func (data *WeatherData) notifyObservers() error {

	for _, o := range data.observers {
		err := o.Update()
		if err != nil {
			return nil
		}
	}

	return nil
}
