package models

type Response struct {
	StatusCode int
	Message    string
	Data       map[string]interface{}
}
