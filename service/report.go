package service

import (
	"project-app-inventory-restapi-golang-fathoni/dto"
	"project-app-inventory-restapi-golang-fathoni/repository"
)

type ReportService interface {
	GetListReports() ([]dto.ReportResponse, error)
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