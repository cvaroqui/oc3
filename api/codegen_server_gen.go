// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /feed/daemon/ping)
	PostFeedDaemonPing(ctx echo.Context) error

	// (POST /feed/daemon/status)
	PostFeedDaemonStatus(ctx echo.Context) error

	// (POST /feed/object/config)
	PostFeedObjectConfig(ctx echo.Context) error

	// (POST /feed/system)
	PostFeedSystem(ctx echo.Context) error

	// (GET /public/openapi)
	GetSwagger(ctx echo.Context) error

	// (GET /version)
	GetVersion(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostFeedDaemonPing converts echo context to params.
func (w *ServerInterfaceWrapper) PostFeedDaemonPing(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{})

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostFeedDaemonPing(ctx)
	return err
}

// PostFeedDaemonStatus converts echo context to params.
func (w *ServerInterfaceWrapper) PostFeedDaemonStatus(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{})

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostFeedDaemonStatus(ctx)
	return err
}

// PostFeedObjectConfig converts echo context to params.
func (w *ServerInterfaceWrapper) PostFeedObjectConfig(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{})

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostFeedObjectConfig(ctx)
	return err
}

// PostFeedSystem converts echo context to params.
func (w *ServerInterfaceWrapper) PostFeedSystem(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{})

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostFeedSystem(ctx)
	return err
}

// GetSwagger converts echo context to params.
func (w *ServerInterfaceWrapper) GetSwagger(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetSwagger(ctx)
	return err
}

// GetVersion converts echo context to params.
func (w *ServerInterfaceWrapper) GetVersion(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetVersion(ctx)
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

	router.POST(baseURL+"/feed/daemon/ping", wrapper.PostFeedDaemonPing)
	router.POST(baseURL+"/feed/daemon/status", wrapper.PostFeedDaemonStatus)
	router.POST(baseURL+"/feed/object/config", wrapper.PostFeedObjectConfig)
	router.POST(baseURL+"/feed/system", wrapper.PostFeedSystem)
	router.GET(baseURL+"/public/openapi", wrapper.GetSwagger)
	router.GET(baseURL+"/version", wrapper.GetVersion)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xYW3PbthL+Kxic85DMMJJ8OefM6DzlUjdpO7UnctoH26OBwBWJlASQxdK24uF/7wAg",
	"JdKi5NhNMn3okyTucnfx7bcX6I5LU1qjQZPj0zuO4KzRDsKP48nEf0ijCTT5r8LaQklByujxR2e0f+Zk",
	"DqXw3/6NsORT/q/xxuY4St34DM2igJLXdZ3wFJxEZb0ZPuWvRMrew6cKHPE64ceTg+/h9YMWFeUG1WdI",
	"o9uj7+H2xOBCpSlo7/M/3wfgd5oAtSjYDPAakP2AaND7PwFI3wgojT5TOnspJViC9FEhWTQWkFRkjFl8",
	"BEnzG0W5qWgujV6qzAv6ARXKETNLFtWZFiU4RrkghvCpUgiOnZ3OztnYyKPxEiAdR81xYzDhiqB024Y7",
	"BnnCaWWBT7kjVDrz520eCESx4vXmQXxtCLo0oMMcCaocI1WCI1Fax25UUbAFMIQlgsshZUuDTJsUguEm",
	"Kz7C02D89RqKPmLCWv+xFao0ZdkkYEuWop0HT/uEwfoapgegSDjo60G9ZQG381LcdoRKE2SAG6nSe6Qk",
	"MAMaViiNVmQQ0jmCMxVKmEtT6R3aBmUOjlDQ8MmtoHxQgOKmQ8WlwVIQn/LFigZZ4qSx8Dj0yFhTmGw1",
	"oOz9R1KnfHoRg7zaYl7Cz4yjTT3OAuG26SJzobN7uR0sAYNMaUdCyy+ohISngkQn+E1YFuFamcrNK5sK",
	"gnQuqAeif/jC18WQl6e8cw3oVOwy/YNRDsxY0O5aMlko0MR81Kx9IXkA+XDEZI3gxtNgNppuupWAFEio",
	"Yju6lyyvSqFfIIhULApgcGsLoUPbZM6CVEslGRlGuXLMSFkhgpbg2yDlcKlt9Di61IOcXPOh7/Y8B/b2",
	"/Pys7VDSpMCeXbw/ef2/w6ODq4TNQIYQ/vucZaDBF0/KFqvo06DKlGYujgXfwIajY0PBdSqTFBUwhInL",
	"DVJyHxpXlaXA1T3jzNsdMfaO2Ozt6Ydf3lzqX0/PWcwXW6Ipu4GR2R1mwuDWj7JL7Y9kK7TGgfNKhZGi",
	"UJ9jVp7BKBslrHJKZ/5VIUldA2tG36XWkBlSQff/zAGwAViPRsfPB1N2n3yRNutEtpgNcc8K+YfIYGBW",
	"oBxucKHSi+KRlRam5JA5FxvljqZxt7dm98PQDOZNyYYjNZaj4yFE3MrRUDF2kFp3w307Uqv/BatA71B9",
	"rx0B3IrSeurzyWgyOniQBrtbjj8lyAoVrWY+2uhqIZySL6s418Ip/Dvh6cZXTmR9wAsQCNhqx18nLRN+",
	"+v2cJx0TQXrfRh24tDQhkbGm+brlmqIASQaZsKqTwyk/GE1Gx2FAW9BeOOVHo8lowuNIDgeJW1zcpsbW",
	"Q+NhNY62u8b7uFExuFWOfGnu3MHajSvUn89QqNV3KZ/em6Z+u+VJ/3JzODncRZe13njHglwn/HByvB16",
	"qZzbjrgNM2mW2gaFRqocWzOkywI+vejl/+KqTu56Ob64qj2PROY8tUTm18Urb6IH9mZy7Id7OOYHoZ21",
	"/QzjDe6VSVdf704z5KruFxVhBfVwbvet8vv296S9+e4nh1fa3Fcf0j3oXDIf0j3qXA7363qlr8qa/kVr",
	"J2s+hM3Ob5kSIXyLa2d8r4p02UOe3qXo25Cn5+KppGl36P7BttnzdZPQmXeD6HsUWVRiqXLSXAOu9qA9",
	"i/a+Dc5NsE9FuD2GX+X/KcotPthqUSg5Xo/WO97cpvuZ/hFodiOyLMz0e6A/7j+mB/8SOf35bwNwC1gE",
	"qUGss541UPXjR6AKtd9iOlfHLTR/W4v+Epr7Cqf1vhPjJ2EhhRULVaiwrV7VkYf+fucCDSss+JSPjTzi",
	"9VX9ZwAAAP//P3UWmf8VAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
