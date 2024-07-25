package helper

import "github.com/julienschmidt/httprouter"

type wrap func(httprouter.Handle) httprouter.Handle

func WrapMiddelware(handle httprouter.Handle, middlewares ...wrap) httprouter.Handle {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handle = middlewares[i](handle)
	}
	return handle
}
