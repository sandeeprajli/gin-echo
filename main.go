package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	c := gin.New()
	c.POST("/test", handle)
	c.Run(":8001")
}

func handle(c *gin.Context) {
	val, err := c.GetRawData()
	if err != nil {
		c.Error(err)
		return
	}
	var obj map[string]interface{}
	err = json.Unmarshal(val, &obj)
	if err != nil {
		c.Error(err)
		return
	}
	sum := 0.0
	for _, value := range obj {
		sum += findSum(value)
	}
	fmt.Println(sum)
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
