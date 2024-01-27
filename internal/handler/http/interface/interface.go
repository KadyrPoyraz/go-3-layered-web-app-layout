package handlerinterface

import (
	"github.com/gorilla/mux"
)

type Handler interface {
	Fill(*mux.Router)
}
