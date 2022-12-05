package api

import (
	"backendbase/database/repository"
	utilities "backendbase/ultilities"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

type AccountHandler struct {
	AccountRepository repository.AccountRepository
}

func (h *AccountHandler) Add(w http.ResponseWriter, r *http.Request) {
	var newAccount bson.M
	if err := BindJSON(r, newAccount); err != nil {
		fmt.Println(err)
		return
	}
	handler := &serviceAccount.AddAccountHandler{
		AccountRepository: h.AccountRepository,
	}
	err := handler.Handle(newAccount)
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_ADD_FAILED, ResponseBody{
			Message: "unable to add user",
			Code:    HTTP_ERROR_CODE_ADD_FAILED,
		})
		return
	}

	WriteJSON(w, http.StatusOK, ResponseBody{
		Message: "add account successfully",
		Code:    http.StatusOK,
	})
}

func (h *AccountHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.AccountRepository.AllAccount()
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get account from server",
		})
	}
	WriteJSON(w, http.StatusOK, accounts)
}

// get by id, update, delete

func (h *AccountHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := utilities.GetQuery(r, "id")
	account, err := h.AccountRepository.FindAccountByIDByID(id)
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get account by id : " + id,
		})
	}
	WriteJSON(w, http.StatusOK, account)
	return
}

func (h *AccountHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	id, _ := utilities.GetQuery(r, "id")
	var updateContact bson.M
	if err := BindJSON(r, updateContact); err != nil {
		fmt.Println(err)
		return
	}
	if err := h.AccountRepository.UpdateAccountByID(id, updateContact); err != nil {
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

func (h *AccountHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	id, _ := utilities.GetQuery(r, "id")
	if err := h.AccountRepository.RemoveAccountByIDByID(id); err != nil {
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
