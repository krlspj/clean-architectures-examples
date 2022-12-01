// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package ports

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /trainings)
	GetTrainings(w http.ResponseWriter, r *http.Request)

	// (POST /trainings)
	CreateTraining(w http.ResponseWriter, r *http.Request)

	// (DELETE /trainings/{trainingUUID})
	CancelTraining(w http.ResponseWriter, r *http.Request, trainingUUID string)

	// (PUT /trainings/{trainingUUID}/approve-reschedule)
	ApproveRescheduleTraining(w http.ResponseWriter, r *http.Request, trainingUUID string)

	// (PUT /trainings/{trainingUUID}/reject-reschedule)
	RejectRescheduleTraining(w http.ResponseWriter, r *http.Request, trainingUUID string)

	// (PUT /trainings/{trainingUUID}/request-reschedule)
	RequestRescheduleTraining(w http.ResponseWriter, r *http.Request, trainingUUID string)

	// (PUT /trainings/{trainingUUID}/reschedule)
	RescheduleTraining(w http.ResponseWriter, r *http.Request, trainingUUID string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// GetTrainings operation middleware
func (siw *ServerInterfaceWrapper) GetTrainings(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetTrainings(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// CreateTraining operation middleware
func (siw *ServerInterfaceWrapper) CreateTraining(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateTraining(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// CancelTraining operation middleware
func (siw *ServerInterfaceWrapper) CancelTraining(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "trainingUUID" -------------
	var trainingUUID string

	err = runtime.BindStyledParameter("simple", false, "trainingUUID", chi.URLParam(r, "trainingUUID"), &trainingUUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter trainingUUID: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CancelTraining(w, r, trainingUUID)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ApproveRescheduleTraining operation middleware
func (siw *ServerInterfaceWrapper) ApproveRescheduleTraining(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "trainingUUID" -------------
	var trainingUUID string

	err = runtime.BindStyledParameter("simple", false, "trainingUUID", chi.URLParam(r, "trainingUUID"), &trainingUUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter trainingUUID: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ApproveRescheduleTraining(w, r, trainingUUID)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// RejectRescheduleTraining operation middleware
func (siw *ServerInterfaceWrapper) RejectRescheduleTraining(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "trainingUUID" -------------
	var trainingUUID string

	err = runtime.BindStyledParameter("simple", false, "trainingUUID", chi.URLParam(r, "trainingUUID"), &trainingUUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter trainingUUID: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RejectRescheduleTraining(w, r, trainingUUID)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// RequestRescheduleTraining operation middleware
func (siw *ServerInterfaceWrapper) RequestRescheduleTraining(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "trainingUUID" -------------
	var trainingUUID string

	err = runtime.BindStyledParameter("simple", false, "trainingUUID", chi.URLParam(r, "trainingUUID"), &trainingUUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter trainingUUID: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RequestRescheduleTraining(w, r, trainingUUID)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// RescheduleTraining operation middleware
func (siw *ServerInterfaceWrapper) RescheduleTraining(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "trainingUUID" -------------
	var trainingUUID string

	err = runtime.BindStyledParameter("simple", false, "trainingUUID", chi.URLParam(r, "trainingUUID"), &trainingUUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter trainingUUID: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RescheduleTraining(w, r, trainingUUID)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL     string
	BaseRouter  chi.Router
	Middlewares []MiddlewareFunc
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/trainings", wrapper.GetTrainings)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/trainings", wrapper.CreateTraining)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/trainings/{trainingUUID}", wrapper.CancelTraining)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/trainings/{trainingUUID}/approve-reschedule", wrapper.ApproveRescheduleTraining)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/trainings/{trainingUUID}/reject-reschedule", wrapper.RejectRescheduleTraining)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/trainings/{trainingUUID}/request-reschedule", wrapper.RequestRescheduleTraining)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/trainings/{trainingUUID}/reschedule", wrapper.RescheduleTraining)
	})

	return r
}
