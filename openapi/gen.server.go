// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// get doc json
	// (GET /api/v1/dm.json)
	GetDocJSON(ctx echo.Context) error
	// get doc html
	// (GET /api/v1/docs)
	GetDocHTML(ctx echo.Context) error
	// get data source list
	// (GET /api/v1/sources)
	DMAPIGetSourceList(ctx echo.Context) error
	// create new data source
	// (POST /api/v1/sources)
	DMAPICreateSource(ctx echo.Context) error
	// delete a data source
	// (DELETE /api/v1/sources/{source-name})
	DMAPIDeleteSource(ctx echo.Context, sourceName string) error
	// enable relay log function for the data source
	// (PATCH /api/v1/sources/{source-name}/start-relay)
	DMAPIStartRelay(ctx echo.Context, sourceName string) error
	// get the current status of the data source
	// (GET /api/v1/sources/{source-name}/status)
	DMAPIGetSourceStatus(ctx echo.Context, sourceName string) error
	// disable relay log function for the data source
	// (PATCH /api/v1/sources/{source-name}/stop-relay)
	DMAPIStopRelay(ctx echo.Context, sourceName string) error
	// get task list
	// (GET /api/v1/tasks)
	DMAPIGetTaskList(ctx echo.Context) error
	// create and start task
	// (POST /api/v1/tasks)
	DMAPIStartTask(ctx echo.Context) error
	// delete and stop task
	// (DELETE /api/v1/tasks/{task-name})
	DMAPIDeleteTask(ctx echo.Context, taskName string, params DMAPIDeleteTaskParams) error
	// get task status
	// (GET /api/v1/tasks/{task-name}/status)
	DMAPIGetTaskStatus(ctx echo.Context, taskName string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetDocJSON converts echo context to params.
func (w *ServerInterfaceWrapper) GetDocJSON(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDocJSON(ctx)
	return err
}

// GetDocHTML converts echo context to params.
func (w *ServerInterfaceWrapper) GetDocHTML(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDocHTML(ctx)
	return err
}

// DMAPIGetSourceList converts echo context to params.
func (w *ServerInterfaceWrapper) DMAPIGetSourceList(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DMAPIGetSourceList(ctx)
	return err
}

// DMAPICreateSource converts echo context to params.
func (w *ServerInterfaceWrapper) DMAPICreateSource(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DMAPICreateSource(ctx)
	return err
}

// DMAPIDeleteSource converts echo context to params.
func (w *ServerInterfaceWrapper) DMAPIDeleteSource(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "source-name" -------------
	var sourceName string

	err = runtime.BindStyledParameterWithLocation("simple", false, "source-name", runtime.ParamLocationPath, ctx.Param("source-name"), &sourceName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter source-name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DMAPIDeleteSource(ctx, sourceName)
	return err
}

// DMAPIStartRelay converts echo context to params.
func (w *ServerInterfaceWrapper) DMAPIStartRelay(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "source-name" -------------
	var sourceName string

	err = runtime.BindStyledParameterWithLocation("simple", false, "source-name", runtime.ParamLocationPath, ctx.Param("source-name"), &sourceName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter source-name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DMAPIStartRelay(ctx, sourceName)
	return err
}

// DMAPIGetSourceStatus converts echo context to params.
func (w *ServerInterfaceWrapper) DMAPIGetSourceStatus(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "source-name" -------------
	var sourceName string

	err = runtime.BindStyledParameterWithLocation("simple", false, "source-name", runtime.ParamLocationPath, ctx.Param("source-name"), &sourceName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter source-name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DMAPIGetSourceStatus(ctx, sourceName)
	return err
}

// DMAPIStopRelay converts echo context to params.
func (w *ServerInterfaceWrapper) DMAPIStopRelay(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "source-name" -------------
	var sourceName string

	err = runtime.BindStyledParameterWithLocation("simple", false, "source-name", runtime.ParamLocationPath, ctx.Param("source-name"), &sourceName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter source-name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DMAPIStopRelay(ctx, sourceName)
	return err
}

// DMAPIGetTaskList converts echo context to params.
func (w *ServerInterfaceWrapper) DMAPIGetTaskList(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DMAPIGetTaskList(ctx)
	return err
}

// DMAPIStartTask converts echo context to params.
func (w *ServerInterfaceWrapper) DMAPIStartTask(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DMAPIStartTask(ctx)
	return err
}

// DMAPIDeleteTask converts echo context to params.
func (w *ServerInterfaceWrapper) DMAPIDeleteTask(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "task-name" -------------
	var taskName string

	err = runtime.BindStyledParameterWithLocation("simple", false, "task-name", runtime.ParamLocationPath, ctx.Param("task-name"), &taskName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter task-name: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params DMAPIDeleteTaskParams
	// ------------- Optional query parameter "source-name-list" -------------

	err = runtime.BindQueryParameter("form", true, false, "source-name-list", ctx.QueryParams(), &params.SourceNameList)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter source-name-list: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DMAPIDeleteTask(ctx, taskName, params)
	return err
}

// DMAPIGetTaskStatus converts echo context to params.
func (w *ServerInterfaceWrapper) DMAPIGetTaskStatus(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "task-name" -------------
	var taskName string

	err = runtime.BindStyledParameterWithLocation("simple", false, "task-name", runtime.ParamLocationPath, ctx.Param("task-name"), &taskName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter task-name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DMAPIGetTaskStatus(ctx, taskName)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/api/v1/dm.json", wrapper.GetDocJSON)
	router.GET(baseURL+"/api/v1/docs", wrapper.GetDocHTML)
	router.GET(baseURL+"/api/v1/sources", wrapper.DMAPIGetSourceList)
	router.POST(baseURL+"/api/v1/sources", wrapper.DMAPICreateSource)
	router.DELETE(baseURL+"/api/v1/sources/:source-name", wrapper.DMAPIDeleteSource)
	router.PATCH(baseURL+"/api/v1/sources/:source-name/start-relay", wrapper.DMAPIStartRelay)
	router.GET(baseURL+"/api/v1/sources/:source-name/status", wrapper.DMAPIGetSourceStatus)
	router.PATCH(baseURL+"/api/v1/sources/:source-name/stop-relay", wrapper.DMAPIStopRelay)
	router.GET(baseURL+"/api/v1/tasks", wrapper.DMAPIGetTaskList)
	router.POST(baseURL+"/api/v1/tasks", wrapper.DMAPIStartTask)
	router.DELETE(baseURL+"/api/v1/tasks/:task-name", wrapper.DMAPIDeleteTask)
	router.GET(baseURL+"/api/v1/tasks/:task-name/status", wrapper.DMAPIGetTaskStatus)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"O7bZ3/WoGH90uDn8JwAA//8aBbz8AEgAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
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

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
