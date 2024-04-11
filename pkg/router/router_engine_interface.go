package router

import (
	"net/http"
)

type RouterEngine interface {
	Get(pattern string, handler http.HandlerFunc)
	Post(pattern string, handler http.HandlerFunc)
	Put(pattern string, handler http.HandlerFunc)
	Patch(pattern string, handler http.HandlerFunc)
	Delete(pattern string, handler http.HandlerFunc)
	Listen(port string)
}
