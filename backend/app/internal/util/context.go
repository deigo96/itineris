package util

import (
	"context"

	constant "github.com/deigo96/itineris/app/internal/const"
	"github.com/gin-gonic/gin"
)

type Context struct {
	ID   int
	Nip  string
	Role string
}

func (c Context) IsAdmin() bool {
	return c.Nip == constant.PPK.String()
}

func GetContext(c *gin.Context) Context {
	id, exist := c.Get("id")
	if !exist {
		return Context{}
	}
	nip, exist := c.Get("nip")
	if !exist {
		return Context{}
	}

	role, exist := c.Get("role")
	if !exist {
		return Context{}
	}

	newId := id.(float64)

	return Context{ID: int(newId), Nip: nip.(string), Role: role.(string)}
}

func GetContexID(c context.Context) int {
	return c.Value("id").(int)
}

func GetContextNIP(c context.Context) string {
	return c.Value("nip").(string)
}
