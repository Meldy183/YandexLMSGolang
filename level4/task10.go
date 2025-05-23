package main

import "reflect"

type Vehicle interface {
	CalculateTravelTime(distance float64) float64
}

type Car struct {
	Type     (string)
	Speed    (float64)
	FuelType (string)
}

func (c Car) CalculateTravelTime(distance float64) float64 {
	return distance / c.Speed
}
func (c Motorcycle) CalculateTravelTime(distance float64) float64 {
	return distance / c.Speed
}

type Motorcycle struct {
	Type     (string)
	Speed    (float64)
	FuelType (string)
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	var ans = make(map[string]float64)
	for _, vehicle := range vehicles {
		ans[reflect.TypeOf(vehicle).String()] = vehicle.CalculateTravelTime(distance)
	}
	return ans
}
