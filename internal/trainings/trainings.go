package trainings

import (
	"strings"
	"time"
	"fmt"
	"strconv"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
	personaldata.Personal 
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		return fmt.Errorf("в строке недостаточно данных")
	}

	t.Steps, err = strconv.Atoi(parts[0])
	if err != nil {
		return  fmt.Errorf("ошибка преобразования: %w", err)
	}

	if t.Steps <= 0 {
		return fmt.Errorf("количество шагов должно быть больше 0")
	}

	t.TrainingType = parts[1]
	t.Duration, err = time.ParseDuration(parts[2]) 

	if err != nil {
		return fmt.Errorf("ошибка преобразования: %w", err)
	}

	if t.Duration <= 0 {
		return fmt.Errorf("продолжительность должна быть больше 0")
	}

	return nil

}

func (t Training) ActionInfo() (string, error) {

	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	var (
		calories float64
		err error
	)

	switch t.TrainingType {
		case "Ходьба":
			calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
			if err != nil {
				return "", err
			}
		case "Бег":
			calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration) 
			if err != nil {
				return "", err
			}
		default:
			return "", fmt.Errorf("неизвестный тип тренировки")
	}

	result := 	fmt.Sprintf(
		"Тип тренировки: %s\n" +
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n",
		 t.TrainingType, t.Duration.Hours(), distance, avgSpeed, calories,
	)

	return result, nil
}
