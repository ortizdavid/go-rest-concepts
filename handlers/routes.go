package handlers

import "net/http"

func RegisterRoutes(router* http.ServeMux) {
	TaskHandler{}.Routes(router)
	UserHandler{}.Routes(router)
	ReportHandler{}.Routes(router)
}