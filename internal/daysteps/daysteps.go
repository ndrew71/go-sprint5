package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps int
  Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	s := strings.Split(datastring, ",")
	if len(s) != 2 {
		return fmt.Errorf("нужно два значения")
	}

	n,err := strconv.Atoi(s[0])
	if err != nil {
		return err
	}
	if n <= 0 {
		return fmt.Errorf("шагов должно быть больше нуля")
	}

	d, err := time.ParseDuration(s[1])
	if err != nil {
		return err
	}
	if d <= 0 {
		return fmt.Errorf("время должно быть больше нуля")
	}
	ds.Steps = n
  ds.Duration = d
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	d := spentenergy.Distance(ds.Steps,ds.Height)
	cal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight,ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps,d,cal), nil
}
