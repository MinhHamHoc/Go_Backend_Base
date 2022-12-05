package api

import (
	"backendbase/database/repository"
	utilities "backendbase/ultilities"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

type UserHandler struct {
	UserRepository repository.UserRepository
}

func (h *UserHandler) Add(w http.ResponseWriter, r *http.Request) {
	var newUser bson.M
	if err := BindJSON(r, newUser); err != nil {
		fmt.Println(err)
		return
	}
	handler := &serviceUser.AddUserHandler{
		UserRepository: h.UserRepository,
	}
	err := handler.Handle(newUser)
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_ADD_FAILED, ResponseBody{
			Message: "unable to add user",
			Code:    HTTP_ERROR_CODE_ADD_FAILED,
		})
		return
	}

	WriteJSON(w, http.StatusOK, ResponseBody{
		Message: "add user successfully",
		Code:    http.StatusOK,
	})
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	start := query.Get("start")
	end := query.Get("end")
	limit := query.Get("limit")
	page := query.Get("page")
	Users, err := h.UserRepository.AllUser()
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get User from server",
		})
	}
	results, err := getByPage(start, end, limit, page, Users)
	if err != nil {
		WriteJSON(w, http.StatusNotFound, ResponseBody{
			Message: err.Error(),
			Code:    http.StatusNotFound,
		})
		return
	}
	WriteJSON(w, http.StatusOK, results)
}

// get by id, update, delete

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := utilities.GetQuery(r, "id")
	User, err := h.UserRepository.FindUserByID(id)
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get User by id : " + id,
		})
	}
	WriteJSON(w, http.StatusOK, User)
	return
}

func (h *UserHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	id, _ := utilities.GetQuery(r, "id")
	var updateContact bson.M
	if err := BindJSON(r, updateContact); err != nil {
		fmt.Println(err)
		return
	}
	if err := h.UserRepository.UpdateUserByID(id, updateContact); err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_UPDATE_FAILED, ResponseBody{
			Message: err,
			Code:    HTTP_ERROR_CODE_UPDATE_FAILED,
		})
		return
	}
	WriteJSON(w, http.StatusOK, ResponseBody{
		Message: "Update successfully",
	})
	return
}

func (h *UserHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	id, _ := utilities.GetQuery(r, "id")
	if err := h.UserRepository.RemoveUserByID(id); err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_DELETE_FAILED, ResponseBody{
			Message: err,
			Code:    HTTP_ERROR_CODE_DELETE_FAILED,
		})
		return
	}
	WriteJSON(w, http.StatusOK, ResponseBody{
		Message: "Update successfully",
	})
	return
}
