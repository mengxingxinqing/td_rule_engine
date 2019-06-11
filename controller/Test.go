package controller

import (
	"fmt"
	"github.com/caibirdme/yql"
	"github.com/labstack/echo"
	"net/http"
)

// User
type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

func Test(c echo.Context) error {
	rawYQL := `name='deen' and age>=23 and (hobby in ('soccer', 'swim') or score>90))`
	result, _ := yql.Match(rawYQL, map[string]interface{}{
		"name":  "deen",
		"age":   int64(23),
		"hobby": "basketball",
		"score": int64(100),
	})
	fmt.Println(result)
	rawYQL = `score ∩ (7,1,9,5,3)`
	result, _ = yql.Match(rawYQL, map[string]interface{}{
		"score": []int64{3, 100, 200},
	})
	fmt.Println(result)
	rawYQL = `score in (7,1,9,5,3)`
	result, _ = yql.Match(rawYQL, map[string]interface{}{
		"score": []int64{3, 5, 2},
	})
	fmt.Println(result)
	rawYQL = `score.sum() > 10`
	result, _ = yql.Match(rawYQL, map[string]interface{}{
		"score": []int{1, 2, 3, 4, 5},
	})
	fmt.Println(result)
	return nil
}

func Test2(c echo.Context) error {
	u := &User{
		Name:  "Jon",
		Email: "jon@labstack.com",
	}
	return c.JSON(http.StatusOK, u)
}