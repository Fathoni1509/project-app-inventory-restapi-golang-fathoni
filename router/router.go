package router

import (
	"net/http"
	"project-app-inventory-restapi-golang-fathoni/handler"
	"project-app-inventory-restapi-golang-fathoni/service"

	mCostume "project-app-inventory-restapi-golang-fathoni/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	// "github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func NewRouter(handler handler.Handler, service service.Service, log *zap.Logger) *chi.Mux {
	r := chi.NewRouter()

	// middleware
	mw := mCostume.NewMiddlewareCustome(service, log)

	// r.Mount("/api/v1", Apiv1(handler, mw))
	// r.Mount("/api/v2", Apiv2(handler))
	r.Mount("/api/v2", Apiv2(handler, mw))

	// //menu
	// r.Route("/user", func(r chi.Router) {
	// 	r.Use(middleware.AuthMiddleware)
	// 	r.Get("/assignments", handler.AssignmentHandler.List)
	// 	r.Get("/success-submit", handler.AssignmentHandler.SuccessSubmit)
	// 	r.Post("/submit-assignment", handler.AssignmentHandler.SubmitAssignment)
	// 	r.Get("/grade", handler.HandlerMenu.GradeView)
	// 	r.Get("/logout", handler.HandlerAuth.LogoutView)
	// })
	// r.Get("/page401", handler.HandlerMenu.PageUnauthorized)

	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public/*", http.StripPrefix("/public/", fs))

	return r
}

/*
func Apiv1(handler handler.Handler, mw mCostume.MiddlewareCostume) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(mw.Logging)
	//authentication
	r.Post("/login", handler.HandlerAuth.Login)
	// r.Post("/logout", handler.HandlerAuth.Logout)

	r.Route("/assignment", func(r chi.Router) {
		r.Get("/", handler.AssignmentHandler.List)
		r.Post("/", handler.AssignmentHandler.Create)
		r.Route("/{assignment_id}", func(r chi.Router) {
			r.Get("/", handler.AssignmentHandler.GetByID)
			r.Put("/", handler.AssignmentHandler.Update)
			r.Delete("/", handler.AssignmentHandler.Delete)
		})
	})

	return r
}*/

func Apiv2(handler handler.Handler, mw mCostume.MiddlewareCostume) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(mw.Logging)
	// CRUD Category
	r.Route("/category", func(r chi.Router) {
		r.Get("/", handler.CategoryHandler.GetListCategories)
		r.Post("/", handler.CategoryHandler.AddCategory)
		r.Route("/{category_id}", func(r chi.Router) {
			r.Get("/", handler.CategoryHandler.GetListCategoryByID)
			r.Put("/", handler.CategoryHandler.UpdateCategory)
			r.Delete("/", handler.CategoryHandler.DeleteCategory)
		})
	})
	// CRUD Warehouse
	r.Route("/warehouse", func(r chi.Router) {
		r.Get("/", handler.WarehouseHandler.GetListWarehouses)
		r.Post("/", handler.WarehouseHandler.AddWarehouse)
		r.Route("/{warehouse_id}", func(r chi.Router) {
			r.Get("/", handler.WarehouseHandler.GetListWarehouseByID)
			r.Put("/", handler.WarehouseHandler.UpdateWarehouse)
			r.Delete("/", handler.WarehouseHandler.DeleteWarehouse)
		})
	})
	// CRUD Shelve
	r.Route("/shelve", func(r chi.Router) {
		r.Get("/", handler.ShelveHandler.GetListShelves)
		r.Post("/", handler.ShelveHandler.AddShelve)
		r.Route("/{shelve_id}", func(r chi.Router) {
			r.Get("/", handler.ShelveHandler.GetListShelveByID)
			r.Put("/", handler.ShelveHandler.UpdateShelve)
			r.Delete("/", handler.ShelveHandler.DeleteShelve)
		})
	})
	// CRUD User
	r.Route("/user", func(r chi.Router) {
		r.Get("/", handler.UserHandler.GetListUsers)
		r.Post("/", handler.UserHandler.AddUser)
		r.Route("/{user_id}", func(r chi.Router) {
			r.Get("/", handler.UserHandler.GetListUserByID)
			r.Put("/", handler.UserHandler.UpdateUser)
			r.Delete("/", handler.UserHandler.DeleteUser)
		})
	})
	// CRUD Product
	r.Route("/product", func(r chi.Router) {
		r.Get("/", handler.ProductHandler.GetListProducts)
		r.Post("/", handler.ProductHandler.AddProduct)
		r.Route("/{product_id}", func(r chi.Router) {
			r.Get("/", handler.ProductHandler.GetListProductByID)
			r.Put("/", handler.ProductHandler.UpdateProduct)
			r.Delete("/", handler.ProductHandler.DeleteProduct)
		})
	})
	// CRUD Sale
	r.Route("/sale", func(r chi.Router) {
		r.Get("/", handler.SaleHandler.GetListSales)
		r.Post("/", handler.SaleHandler.AddSale)
		r.Route("/{sale_id}", func(r chi.Router) {
			r.Get("/", handler.SaleHandler.GetListSaleByID)
			r.Put("/", handler.SaleHandler.UpdateSale)
			r.Delete("/", handler.SaleHandler.DeleteSale)
		})
	})
	// read report
	r.Get("/report", handler.ReportHandler.GetListReports)
	// check stock
	r.Get("/minstock", handler.ReportHandler.GetListMinStocks)

	return r
}
