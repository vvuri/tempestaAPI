package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testAPI/internal/handlers"
)

// как подксказка
var _ handlers.Handler = &handler{}

const (
	usersURL = "/api/users"
	userURL  = "/api/users/:uuid"
)

type handler struct {
}

func NewHandler() handlers.Handler { // возвращаем интерфейс
	return &handler{} // на самом деле структуру
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpgradeUser)
	router.PATCH(userURL, h.PartiallyUpgradeUser)
	router.DELETE(usersURL, h.DeleteUser)
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	w.Write([]byte("API Get User"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	w.Write([]byte("API Create User"))
}

func (h *handler) UpgradeUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	w.Write([]byte("API Update User"))
}

func (h *handler) PartiallyUpgradeUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	w.Write([]byte("API Partially update user"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	w.Write([]byte("API Delete User"))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	w.Write([]byte("API List Users"))
}
