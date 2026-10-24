package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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
	s := strings.Split(datastring,",")
	if len(s) != 3 {
		return fmt.Errorf("нужно три значения")
	}

	n, err := strconv.Atoi(s[0])
	if err != nil {
		return err
	}
	if n <= 0 {
		return fmt.Errorf("шагов должно быть больше нуля")
	}
	d,err := time.ParseDuration(s[2])
	if err != nil {
		return err
	}
	if d <= 0 {
		return fmt.Errorf("время должно быть больше нуля")
	}

	t.Steps = n
	t.TrainingType = s[1]
  t.Duration = d
	return nil
}

func (t Training) ActionInfo() (string, error) {
	d := spentenergy.Distance(t.Steps,t.Height)
	v := spentenergy.MeanSpeed(t.Steps, t.Height,t.Duration)

	var cal float64
	var err error
	switch t.TrainingType {
	case "Бег":
		cal, err = spentenergy.RunningSpentCalories(t.Steps,t.Weight, t.Height,t.Duration)
	case "Ходьба":
		cal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight,t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	if err != nil {
		return "",err
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType,t.Duration.Hours(), d,v, cal),nil
}
