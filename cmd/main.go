package main

import (
	"amberovsky/teapicker/api/employees"
	"amberovsky/teapicker/api/members"
	"amberovsky/teapicker/api/pick"
	"amberovsky/teapicker/api/teams"
	employeeInternal "amberovsky/teapicker/internal/employee"
	memberInternal "amberovsky/teapicker/internal/member"
	pickInternal "amberovsky/teapicker/internal/pick"
	teamInternal "amberovsky/teapicker/internal/team"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	memberRepository := memberInternal.NewInMemoryRepository()

	employeesHandler := employees.NewHandler(employees.NewService(employeeInternal.NewInMemoryRepository(), memberRepository))
	teamsHandler := teams.NewHandler(teams.NewService(teamInternal.NewInMemoryRepository(), memberRepository))
	membersHandler := members.NewHandler(members.NewService(memberRepository))
	pickHandler := pick.NewHandler(pick.NewService(pickInternal.NewRandomStrategy(memberRepository).Invoke))

	r := mux.NewRouter()

	r.HandleFunc("/team", teamsHandler.GetTeams).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/team", teamsHandler.Add).Methods(http.MethodPost)
	r.HandleFunc("/team/{uuid}", teamsHandler.Delete).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/employee", employeesHandler.GetEmployees).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/employee", employeesHandler.Add).Methods(http.MethodPost)
	r.HandleFunc("/employee/{uuid}", employeesHandler.Delete).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/member", membersHandler.GetMembers).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/member", membersHandler.Add).Methods(http.MethodPost)
	r.HandleFunc("/member/{uuid}", membersHandler.Get).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/member/{uuid}", membersHandler.Delete).Methods(http.MethodDelete)

	r.HandleFunc("/pick/{uuid}", pickHandler.Pick).Methods(http.MethodPost, http.MethodOptions)

	r.Use(mux.CORSMethodMiddleware(r))

	_ = http.ListenAndServe(":8080", r)
}
