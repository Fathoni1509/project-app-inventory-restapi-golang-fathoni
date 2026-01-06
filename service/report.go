package service

import (
	"project-app-inventory-restapi-golang-fathoni/dto"
	"project-app-inventory-restapi-golang-fathoni/repository"
)

type ReportService interface {
	GetListReports() ([]dto.ReportResponse, error)
	GetListMinStocks() ([]dto.MinStockResponse, error)
}

type reportService struct {
	Repo repository.Repository
}

func NewReportService(repo repository.Repository) ReportService {
	return &reportService{Repo: repo}
}

// service get list reports
func (rep *reportService) GetListReports() ([]dto.ReportResponse, error) {
	return rep.Repo.ReportRepo.GetListReports()
}

// service get list min stock
func (rep *reportService) GetListMinStocks() ([]dto.MinStockResponse, error) {
	return rep.Repo.ReportRepo.GetListMinStocks()
}