package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше 0")
	}

	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше 0")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше 0")
	}

	avgSpeed := MeanSpeed(steps, height, duration)
	durationInMins := duration.Minutes()
	result := (weight * avgSpeed * durationInMins)/ float64(minInH)

	return result * walkingCaloriesCoefficient, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше 0")
	}

	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше 0")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше 0")
	}

	avgSpeed := MeanSpeed(steps, height, duration)
	durationInMins := duration.Minutes()

	 return (weight * avgSpeed * durationInMins)/ float64(minInH), nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {

	if steps < 0 {
		return 0
	}
	
	if duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)

	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	
	stepLength := height * stepLengthCoefficient

	return  stepLength * float64(steps) / float64(mInKm) 
}
