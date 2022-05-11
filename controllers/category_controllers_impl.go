package controllers

import (
	"net/http"
	"restapi-golang/helper"
	"restapi-golang/models/web"
	"restapi-golang/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllersImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryControllers(categoryService service.CategoryService) CategoryControllers {
	return &CategoryControllersImpl{
		CategoryService: categoryService,
	}
}

// create
func (controller *CategoryControllersImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

// update
func (controllers *CategoryControllersImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controllers.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

// delete
func (controllers *CategoryControllersImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controllers.CategoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}

	helper.WriteResponseBody(writer, webResponse)
}

// find by id
func (controllers *CategoryControllersImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controllers.CategoryService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

// find all
func (controllers *CategoryControllersImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponse := controllers.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}
