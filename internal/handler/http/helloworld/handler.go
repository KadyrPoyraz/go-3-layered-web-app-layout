package helloworldhandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	handlerinterface "github.com/KadyrPoyraz/httplayout/internal/handler/http/interface"
	"github.com/gorilla/mux"
)

type User struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}

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
	users := User{
		Name:     "Solomon",
		LastName: "Kchay",
	}
	b, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(b))
}
