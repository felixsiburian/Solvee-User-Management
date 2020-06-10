package controllers

import "Solvee-User-Management/api/middlewares"

func (s *Server) InitializeRoutes() {
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	s.Router.HandleFunc("/regencies", middlewares.SetMiddlewareJSON(s.GetRegencies)).Methods("GET")
	s.Router.HandleFunc("/regencies/{id}", middlewares.SetMiddlewareJSON(s.GetRegency)).Methods("GET")

	s.Router.HandleFunc("/provinces", middlewares.SetMiddlewareJSON(s.GetProvinces)).Methods("GET")
	s.Router.HandleFunc("provinces/{id}", middlewares.SetMiddlewareJSON(s.GetProvince)).Methods("GET")

	s.Router.HandleFunc("/villages", middlewares.SetMiddlewareJSON(s.GetVillages)).Methods("GET")
	s.Router.HandleFunc("/villages/{id}", middlewares.SetMiddlewareJSON(s.GetVillage)).Methods("GET")

	s.Router.HandleFunc("/districts", middlewares.SetMiddlewareJSON(s.GetDistricts)).Methods("GET")
	s.Router.HandleFunc("/districts/{id}", middlewares.SetMiddlewareJSON(s.GetDistrict)).Methods("GET")

	s.Router.HandleFunc("/categories", middlewares.SetMiddlewareJSON(s.GetCategory)).Methods("GET")

	s.Router.HandleFunc("/reports", middlewares.SetMiddlewareJSON(s.CreateReport)).Methods("POST")
	s.Router.HandleFunc("/reports", middlewares.SetMiddlewareJSON(s.GetReports)).Methods("GET")
	s.Router.HandleFunc("/reports/{id}", middlewares.SetMiddlewareJSON(s.GetAReport)).Methods("GET")
	s.Router.HandleFunc("/reports/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateReport))).Methods("PUT")
	s.Router.HandleFunc("/reports/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteReport)).Methods("DELETE")


}