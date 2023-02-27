package main

type RequestParams struct {
	Uuid       string `validate:"uuid4"`
	SourceName string `validate:"min=3"`
}
