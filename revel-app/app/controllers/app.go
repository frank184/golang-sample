package controllers

import (
	"strconv"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

type Response struct {
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors"`
	Meta   interface{} `json:"meta"`
}

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

const UserCount = 1e4

func (c App) Index() revel.Result {
	users := []User{}
	for i := 0; i < UserCount; i++ {
		user := User{
			"John" + strconv.Itoa(i),
			"Doe" + strconv.Itoa(i),
			"johndoe@mail.com" + strconv.Itoa(i),
		}
		users = append(users, user)
	}
	return setupResponse(c, users, nil, nil)
}

func setupResponse(c App, data interface{}, errors interface{}, meta interface{}) revel.Result {
	res := Response{data, errors, meta}
	c.Response.ContentType = "application/json; charset=utf-8"
	c.Response.Status = 200
	return c.RenderJSON(res)
}
