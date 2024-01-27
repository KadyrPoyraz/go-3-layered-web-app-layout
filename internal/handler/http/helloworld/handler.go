package helloworldhandler

import (
	"fmt"
	"net/http"

	handlerinterface "github.com/KadyrPoyraz/httplayout/internal/handler/http/interface"
	"github.com/gorilla/mux"
)

type helloWorldHandler struct{}

func New() handlerinterface.Handler {
	return &helloWorldHandler{}
}

func (h *helloWorldHandler) Fill(router *mux.Router) {
    base := "/hello"
    r := router.PathPrefix(base).Subrouter()

    r.HandleFunc("/", h.handleHello)
}

func (h *helloWorldHandler) handleHello(w http.ResponseWriter, r *http.Request) {
   // helloRepo := helloworld.NewPostgresqlRepo(db)
   //
   //  helloService := helloworldService.New(helloRepo)
   //
   //  ctx := context.Background()
   //  helloService.SayHello(ctx)
    fmt.Fprintf(w, "Hello from handler layear!")
}
