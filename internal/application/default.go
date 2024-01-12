package application

import (
	"first_api/internal/handler"
	"first_api/internal/loader"
	"first_api/internal/repository"
	"first_api/internal/service"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type DefaultHTTP struct {
	addr string
}

func NewDefaultHTTP(addr string) *DefaultHTTP {
	return &DefaultHTTP{
		addr: addr,
	}
}

func (h *DefaultHTTP) Run() (err error) {
	// loader
	loader := loader.NewDataLoaded()
	loadData, err := loader.LoadData()
	if err != nil {
		fmt.Println(err)
		return
	}
	// repository
	rp := repository.NewProductMap(loadData, 0)
	// service
	sv := service.NewProductDefault(rp)
	// handler
	hd := handler.NewDefaultProducts(sv)
	// router
	rt := chi.NewRouter()

	// endpoints
	rt.Get("/products", hd.GetProducts())
	// run http service
	err = http.ListenAndServe(h.addr, rt)
	return
}
