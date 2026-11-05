package serverhandlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

// The uuid of a node already registered is its password: an anonymous
// caller, who knows only its nodename and app, does not get it.
func TestMayResendNodeUUID(t *testing.T) {
	for mode, want := range map[string]bool{
		AuthModeAnonRegister: false,
		AuthModeNode:         true,
		AuthModeUser:         true,
	} {
		c := echo.New().NewContext(httptest.NewRequest(http.MethodPost, "/api/auth/node", nil), httptest.NewRecorder())
		c.Set(XAuthMode, mode)
		if got := mayResendNodeUUID(c); got != want {
			t.Errorf("%s: got %v, want %v", mode, got, want)
		}
	}
}
