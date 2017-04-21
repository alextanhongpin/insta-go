package authsvc

import (
	// "fmt"
	"net/http"
	"net/http/httptest"
	// "net/url"
	// "strings"
	"testing"
)

// func TestLogin(t *testing.T) {
//  assert := assert.New(t)
//  t.Skip("Unable to test this for now")

//  tests := []struct {
//    description        string
//    method             string
//    email              string
//    password           string
//    url                string
//    expectedStatusCode int
//    expectedBody       string
//    isForm             bool
//  }{
//    {
//      description:        "request login",
//      method:             "GET",
//      email:              "john.doe@mail.com",
//      password:           "123456",
//      url:                "/login",
//      expectedStatusCode: 200,
//      expectedBody:       `[{"name": "test"}]`,
//    },
//  }

//  for _, tc := range tests {
//    form := url.Values{}
//    if tc.email != "" {
//      form.Add("email", email)
//    }
//    if tc.password != "" {
//      form.Add("password", password)
//    }
//    req, err := http.NewRequest(tc.method, tc.url, strings.NewReader(form.Encode()))

//    w := httptest.NewRecorder()
//    service := Service{}

//    endpoint.Login(service)(w, req)

//    assert.Equal(tc.expectedStatusCode, w.Code, tc.description)
//    assert.Equal(tc.expectedBody, w.Body.String(), tc.description)
//  }
// }

// func TestSuccessfulLogin(t *testing.T) {
//  t.Skip("Skipping login")
//  assert := assert.New(t)

//  form := url.Values{}
//  form.Add("email", "john.doe@mail.com")
//  form.Add("password", "123456")

//  req, err := http.NewRequest("POST", "/login", strings.NewReader(form.Encode()))

//  w := httptest.NewRecorder()
//  endpoint.Login(Service{})(w, req)

//  // assert.Equal(tc.expectedStatusCode, w.Code, tc.description)
//  // assert.Equal(tc.expectedBody, w.Body.String(), tc.description)

// }

func TestMockEndpoint(t *testing.T) {
	// Output: Testing the mock endpoint
	endpoint := Endpoint{}
	service := &Service{common.GetDatabaseContext()}

	req, _ := http.NewRequest("GET", "/mock", nil)

	w := httptest.NewRecorder()
	endpoint.Mock(service)(w, req, nil)

	expectedBody := `{"message":"Hello world"}`
	actualBody := w.Body.String()

	// assert.Equal(302, w.Code, "Status code is valid")
	if actualBody != expectedBody {
		t.Fatalf("Expected %s but got %s", expectedBody, actualBody)
	}
}
