package handler

import (
	"net/http"
	"project-app-inventory-restapi-golang-fathoni/service"
	"project-app-inventory-restapi-golang-fathoni/utils"
)

type ReportHandler struct {
	ReportService service.ReportService
	Config utils.Configuration
}

func NewReportHandler(reportService service.ReportService, config utils.Configuration) ReportHandler {
	return ReportHandler{
		ReportService: reportService,
		Config: config,
	}
}

// get list reports
func (reportHandler *ReportHandler) GetListReports(w http.ResponseWriter, r *http.Request) {
	// Get data report form service all report
	reports, err := reportHandler.ReportService.GetListReports()
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to fetch reports: "+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success get data report", reports)

}

// get list reports
func (reportHandler *ReportHandler) GetListMinStocks(w http.ResponseWriter, r *http.Request) {
	// Get data report form service all report
	reports, err := reportHandler.ReportService.GetListMinStocks()
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to fetch minimum stock: "+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success get data minimum stock product", reports)

}