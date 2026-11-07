package daysteps

import (
	"fmt"
	"time"
	"strings"
	"strconv"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
           s := strings.Split(data, ",")
		   if len(s) != 2 {
			     return 0, 0, fmt.Errorf("Неверный формат входных данных: %q", data) 
		   }  
           steps, err := strconv.Atoi(s[0])
		      if err != nil {
				return 0, 0, fmt.Errorf("Не задано количество шагов: %q", s[0])
		 	  }
		   if steps <= 0 {
			    return 0, 0, fmt.Errorf("Неверное количество шагов %d", steps) 
		   } 
		   t, err := time.ParseDuration(s[1])
		      if err != nil {
				return 0, 0, fmt.Errorf("Время задано неверно: %q", s[1])
			  } 
          return steps, t, nil

}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
              steps, t, err := parsePackage(data)
			       if err != 0 {
                      fmt.Println(err)
                      return "" 
					}
			  if steps <= 0 {
                      return ""
			  }
             distance := float64(steps) * stepLength / mInKm
             calories := spentcalories.WalkingSpentCalories(steps, weight, height, t)
			 return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f.\nВы сожгли %.2f.\n", steps, distance, calories)
}
