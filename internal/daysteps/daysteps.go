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
//Отформатировал функцию, надеюсь правильно
func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
    s := strings.Split(data, ",")
	if len(s) != 2 {
	//Исправил вывод ошибки
	    return 0, 0, fmt.Errorf("incorrect data format: %q", data) 
	}  
    steps, err := strconv.Atoi(s[0])
	if err != nil {
	//Исправил вывод ошибки	
		return 0, 0, fmt.Errorf("conversation error: %w", err)
	}
	if steps <= 0 {
	//Исправил вывод ошибки	
	    return 0, 0, fmt.Errorf("steps are set incorrectly %d", steps) 
	} 
	t, err := time.ParseDuration(s[1])
	if err != nil {
	//Исправил вывод ошибки	
		return 0, 0, fmt.Errorf("time are set incorrectly: %w", err)
	}
	if t <= 0 {
	//Исправил вывод ошибки	
        return 0, 0, fmt.Errorf("time are set incorrectly: %q", s[1])
	}	   
    return steps, t, nil
}
//Отформатировал функцию, надеюсь правильно
func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	var calories float64
    steps, t, err := parsePackage(data)
	if err != nil {
        fmt.Println(err)
        return "" 
	}
	if steps <= 0 {
        return ""
	}
	if t <= 0 {
        return ""
	}	   
    distance := float64(steps) * stepLength / mInKm
    calories, err = spentcalories.WalkingSpentCalories(steps, weight, height, t)
	if err != nil {
	//Исправил вывод ошибки	
	    return fmt.Sprintf("incorrect calories counting: %w", err)
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)
}
