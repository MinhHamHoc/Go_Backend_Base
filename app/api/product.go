package api

import (
	utilities "backendbase/ultilities"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

type ProductHandler struct {
	ProductRepository repository.ProductRepository
}

func (h *ProductHandler) Add(w http.ResponseWriter, r *http.Request) {
	var newProduct bson.M
	if err := BindJSON(r, newProduct); err != nil {
		fmt.Println(err)
		return
	}
	handler := &serviceProduct.AddProductHandler{
		ProductRepository: h.ProductRepository,
	}
	err := handler.Handle(newProduct)
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_ADD_FAILED, ResponseBody{
			Message: "unable to add product",
			Code:    HTTP_ERROR_CODE_ADD_FAILED,
		})
		return
	}

	WriteJSON(w, http.StatusOK, ResponseBody{
		Message: "add Product successfully",
		Code:    http.StatusOK,
	})
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	start, _ := utilities.GetQuery(r, "start")
	end, _ := utilities.GetQuery(r, "end")
	limit, _ := utilities.GetQuery(r, "limit")
	page, _ := utilities.GetQuery(r, "page")
	Products, err := h.ProductRepository.All()
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get Product from server",
		})
	}
	results, err := getByPage(start, end, limit, page, Products)
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

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := utilities.GetQuery(r, "id")
	Product, err := h.ProductRepository.GetByID(id)
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get Product by id : " + id,
		})
	}
	WriteJSON(w, http.StatusOK, Product)
	return
}

func (h *ProductHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	id, _ := utilities.GetQuery(r, "id")
	var updateContact bson.M
	if err := BindJSON(r, updateContact); err != nil {
		fmt.Println(err)
		return
	}
	if err := h.ProductRepository.UpdateByID(id, updateContact); err != nil {
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

func (h *ProductHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	id, _ := utilities.GetQuery(r, "id")
	if err := h.ProductRepository.DeleteByID(id); err != nil {
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
