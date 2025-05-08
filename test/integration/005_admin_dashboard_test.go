package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"api/internal/adapter/http"
	"github.com/stretchr/testify/assert"
	"net/http/cookiejar"
)

func TestAdminDashboardAccess(t *testing.T) {
	handler := &http.AdminHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()

	jar, _ := cookiejar.New(nil)

	client := &http.Client{Jar: jar}

	// Helper: set cookie for admin or member
	setRoleCookie := func(role string) {
		cookie := &http.Cookie{
			Name:  "session_token",
			Value: makeFakeJWTWithRole(role),
			Path:  "/",
		}
		u := server.URL + "/admin"
		req, _ := http.NewRequest("GET", u, nil)
		client.Jar.SetCookies(req.URL, []*http.Cookie{cookie})
	}

	t.Run("admin user can access /admin", func(t *testing.T) {
		setRoleCookie("admin")
		resp, err := client.Get(server.URL)
		assert.NoError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("member user is redirected", func(t *testing.T) {
		setRoleCookie("member")
		resp, err := client.Get(server.URL)
		assert.NoError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, http.StatusSeeOther, resp.StatusCode)
		location, _ := resp.Location()
		assert.Contains(t, location.String(), "/dashboard")
	})

	t.Run("/admin loads for admin", func(t *testing.T) {
		setRoleCookie("admin")
		resp, err := client.Get(server.URL)
		assert.NoError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

// makeFakeJWTWithRole returns a dummy JWT with the given "role" claim and valid signature.
func makeFakeJWTWithRole(role string) string {
	// For test only, alg: none, payload: {"role":role}
	// JWT: header.payload.signature (base64url)
	// Header: {"alg":"none","typ":"JWT"}
	header := "eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0"
	payload := `{"role":"` + role + `"}`
	// base64url encode payload
	payloadEnc := base64URLEncode([]byte(payload))
	// No signature for alg:none
	return header + "." + payloadEnc + "."
}

// base64URLEncode encodes bytes to base64url without padding.
func base64URLEncode(src []byte) string {
	const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	n := len(src)
	var dst []byte
	for i := 0; i < n; i += 3 {
		var b [3]byte
		copy(b[:], src[i:])
		v := uint(b[0])<<16 | uint(b[1])<<8 | uint(b[2])
		dst = append(dst, encodeStd[(v>>18)&0x3F])
		dst = append(dst, encodeStd[(v>>12)&0x3F])
		if i+1 < n {
			dst = append(dst, encodeStd[(v>>6)&0x3F])
		}
		if i+2 < n {
			dst = append(dst, encodeStd[v&0x3F])
		}
	}
	return string(dst)
}