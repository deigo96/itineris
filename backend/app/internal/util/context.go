package util

import "context"

type Context struct {
	ID   int
	Nip  string
	Role string
}

func GetContext(c context.Context) Context {
	id := c.Value("id").(int)
	nip := c.Value("nip").(string)
	role := c.Value("role").(string)
	return Context{ID: id, Nip: nip, Role: role}
}

func GetContexID(c context.Context) int {
	return c.Value("id").(int)
}

func GetContextNIP(c context.Context) string {
	return c.Value("nip").(string)
}
