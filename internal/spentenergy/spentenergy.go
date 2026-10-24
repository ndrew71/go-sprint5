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
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("неверные данные прогулки")
	}

	v := MeanSpeed(steps,height, duration)
	cal := weight*v*duration.Minutes() / minInH
	return cal * walkingCaloriesCoefficient,nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("неверные данные тренировки")
	}
	v := MeanSpeed(steps, height,duration)
	cal := weight * v * duration.Minutes() / minInH
	return cal, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}

	d := Distance(steps,height)
	return d / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	step := height * stepLengthCoefficient
  return float64(steps) * step / mInKm
}
