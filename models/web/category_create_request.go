package web

type CategoryCreateRequest struct {
	Name string `validate:"required,max=150,min=1" json:"name"`
}
