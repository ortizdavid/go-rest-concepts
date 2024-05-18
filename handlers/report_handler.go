package handlers

import (
	"fmt"
	"net/http"
	"github.com/ortizdavid/go-nopain/docgen"
	"github.com/ortizdavid/go-nopain/httputils"
	"github.com/ortizdavid/go-rest-concepts/models"
)

type ReportHandler struct {
}


func (re ReportHandler) Routes(router *http.ServeMux) {
	router.HandleFunc("GET /api/reports", re.reports)
}


func (ReportHandler) reports(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("param")
	var report models.ReportModel
	var tableReport models.TableReport
	pdfGen :=  docgen.NewHtmlPdfGenerator()

	switch param {
	case "users":
		tableReport = report.GetAllUsers()
	case "tasks":
		tableReport = report.GetAllTasks()
	default:
		httputils.WriteJsonError(w, "Invalid report type", http.StatusBadRequest)
		return
	}
	//-----------------------
	templateFile :=  param +".html"
	title := "Report: " +tableReport.Title
	fileName := title +".pdf"
	data := map[string]any{
		"Title": title,
		"AppName": "Task Management App",
		"Rows": tableReport.Rows,
		"Count": tableReport.Count,
	}
	//----------- Render PDF
	pdfBytes, err := pdfGen.GeneratePDF(fmt.Sprintf("./templates/reports/%s", templateFile), data)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = pdfGen.SetOutput(w, pdfBytes, fileName)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
	}
}