package spentcalories

import (
	"time"
	"strings"
	"strconv"
	"log"
	"fmt"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	        s := strings.Split(data, ",")
			   if len(s) != 3 {
				  return 0, "", 0, fmt.Errorf("Неверные входные данные: %q", data)
			   }
            steps, err := strconv.Atoi(s[0])  
               if err != nil {
				  return 0, "", 0, fmt.Errorf("Не задано количество шагов: %q", s[0])
			   } 
            if steps <= 0 {
                  return 0, "", 0, fmt.Errorf("Неверное количество шагов: %d", steps)
			}
			t, err := time.ParseDuration(s[2])
			    if err != nil {
					return 0, "", 0, fmt.Errorf("Время задано неверно: %q", s[2])
				}
			if t <=0 {
					return 0, "", 0, fmt.Errorf("Время задано неверно: %q", s[2])
			}	
			return steps, s[1], t, nil	 			   
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	        stepLength := height * stepLengthCoefficient
			distance := float64(steps) * stepLength / mInKm
			return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	        if duration <= 0 {
				return 0
			}
			d := distance(steps, height)
			speed := d / duration.Hours()
			return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	       var calories float64
           steps, kind, t, err := parseTraining(data)
		       if err != nil {
				    log.Println(err)
					return "", err
			   }
		       dist := distance(steps, height)
			   spd := meanSpeed(steps, height, t)
		   switch kind {
		       case "Ходьба":
    	 		     calories, err = WalkingSpentCalories(steps, weight, height, t)
				case "Бег":
					 calories, err = RunningSpentCalories(steps, weight, height, t)
                default:
					 return "", fmt.Errorf("неизвестный тип тренировки")
		   }
		       if err != nil {
				     log.Println(err)
					 return "", err
			   }
		   return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", kind, t.Hours(), dist, spd, calories), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	         if steps <= 0 {
                  return 0, fmt.Errorf("Неверное количество шагов: %d", steps)
			 }
			 if duration <= 0 {
					return 0, fmt.Errorf("Время задано неверно: %v", duration)
				}
			if weight <= 0 {
					return 0, fmt.Errorf("Вес задан неверно: %.2f", weight)
				}	
			if height <= 0 {
					return 0, fmt.Errorf("Рост задан неверно: %.2f", height)
				}
			spd := meanSpeed(steps, height, duration)
			durationInMinutes := duration.Minutes()		
			return (weight * spd * durationInMinutes) / minInH, nil	

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	         if steps <= 0 {
                  return 0, fmt.Errorf("Неверное количество шагов: %d", steps)
			 }
			 if duration <= 0 {
					return 0, fmt.Errorf("Время задано неверно: %v", duration)
				}
			if weight <= 0 {
					return 0, fmt.Errorf("Вес задан неверно: %.2f", weight)
				}	
			if height <= 0 {
					return 0, fmt.Errorf("Рост задан неверно: %.2f", height)
				}
			spd := meanSpeed(steps, height, duration)
			durationInMinutes := duration.Minutes()		
			return (weight * spd * durationInMinutes) / minInH * walkingCaloriesCoefficient, nil	
}
