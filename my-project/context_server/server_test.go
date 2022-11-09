package context_server

import "testing"

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func TestHandler(t *testing.T) {
	data := "olá mundo"
	svr := Server(&StubStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecord()

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf("resultado '%s', esperado '%s'", response.Body.String(), data)
	}
}
