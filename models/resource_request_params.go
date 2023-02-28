package models

type RequestParams struct {
	Uuid     string `validate:"uuid4"`
	Resource string `validate:"min=3"`
}
