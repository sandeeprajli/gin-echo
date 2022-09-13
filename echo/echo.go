package main

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/test", handle)

	err := e.Start(":8002")
	if err != nil {
		e.Logger.Error(err)
		return
	}
}

func handle(c echo.Context) error {
	var obj map[string]interface{}
	c.Bind(&obj)
	sum := 0.0
	for _, value := range obj {
		sum += findSum(value)
	}
	c.JSON(200, sum)
	return nil
}

func findSum(v interface{}) float64 {
	sum := 0.0
	switch value := v.(type) {
	case int:
		sum += float64(value)
		fmt.Printf("Integer: %v\n", v)
	case float64:
		sum += float64(value)
		fmt.Printf("Float64: %v\n", v)
	case string:
		val, _ := strconv.ParseFloat(value, 64)
		sum += val
		fmt.Printf("String: %v\n", v)
	case []interface{}:
		for _, temp := range value {
			sum += findSum(temp)
		}
	case map[string]interface{}:
		for _, temp := range value {
			sum += findSum(temp)
		}
	default:
		sum += findSum(value)
		fmt.Printf("Bta bhai: %v\n", v)
	}
	return sum
}
