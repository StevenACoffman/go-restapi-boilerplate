// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi"
	"net/http"
	"strings"
)

type ServerInterface interface {
	// create user (POST /api/v1/user)
	CreateUser(w http.ResponseWriter, r *http.Request)
	// get user (GET /api/v1/user/{userID})
	ReadUser(w http.ResponseWriter, r *http.Request)
}

// CreateUser operation middleware
func CreateUserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ReadUser operation middleware
func ReadUserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var err error

		// ------------- Path parameter "userID" -------------
		var userID string

		err = runtime.BindStyledParameter("simple", false, "userID", chi.URLParam(r, "userID"), &userID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter userID: %s", err), http.StatusBadRequest)
			return
		}

		ctx = context.WithValue(ctx, "userID", userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerFromMux(si, chi.NewRouter())
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r *chi.Mux) http.Handler {
	r.Group(func(r chi.Router) {
		r.Use(CreateUserCtx)
		r.Post("/api/v1/user", si.CreateUser)
	})
	r.Group(func(r chi.Router) {
		r.Use(ReadUserCtx)
		r.Get("/api/v1/user/{userID}", si.ReadUser)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xTTY/UMAz9K8hwrCYzgDjkCOxhhBDSij2hFQqpp+tV84HjjjSq+t+R03YZlkHiAIhL",
	"mzi2n/2ePYJPIaeIUQrYEYq/w+Dq8Yr5GktOsaBeM6eMLIT1MZROf3LKCBaKMMUOpgYYvw5Y5DO1F56n",
	"ZrWkL/foRQOOu81NQf4ZYc5wSBycgAWK8uolPCSgKNgha4ahIF/GW96iC/hb1aiJ4iGpc08el87neHi/",
	"/6gphaTXq6aGBo7IhVIEC7vNdrNVj5Qxukxg4UU1NZCd3NWmjMtkjjszrC2nIvrXxp1QivsWLLxhdII3",
	"M8DC6OvUntTTpygYa5DLuSdfw8x90RpW/fT0jPEAFp6a7wKbRV2zkl47VgBibMEKD1gNs+q14ufb7R+D",
	"PZ+oCt1i8UxZZgI/vINqO7ihl38FesWcuA5DGUJwfAILvvL/ZFFYXFfAfoLK2K16nqtoRv3u305aR4cX",
	"1LxG1z5o+ZeY/UHQ/5TVDuUXlOqCsAsoyGoegTSHLg006/bNLMPjcW3O6n2837fTNH0LAAD//6sA4e/g",
	"BAAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}