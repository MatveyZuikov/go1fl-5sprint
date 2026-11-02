package daysteps

import (
	"fmt"
	"time"
	"strings"
	"strconv"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	personaldata.Personal 
	Steps int
	Duration time.Duration
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	
	parts := strings.Split(datastring, ",")

	if len(parts) != 2 {
		return fmt.Errorf("в строке недостаточно данных")
	}

	ds.Steps, err = strconv.Atoi(parts[0])
	if err != nil {
		return  fmt.Errorf("ошибка преобразования: %w", err)
	}

	if ds.Steps <= 0 {
		return fmt.Errorf("количество шагов должно быть больше 0")
	}

	ds.Duration, err = time.ParseDuration(parts[1]) 

	if err != nil {
		return fmt.Errorf("ошибка преобразования: %w", err)
	}

	if ds.Duration <= 0 {
		return fmt.Errorf("продолжительность должна быть больше 0")
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
		if err != nil {
			return "", err
		}
	
	return fmt.Sprintf(
		"Количество шагов: %d.\n" +
		"Дистанция составила %.2f км.\n" +
		"Вы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories,
	) , nil
}
