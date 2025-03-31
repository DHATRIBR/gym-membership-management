package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

// Helper function to create HTTP requests and record responses
func performRequest(handlerFunc http.HandlerFunc, method, path string, body []byte) *httptest.ResponseRecorder {
    req := httptest.NewRequest(method, path, bytes.NewBuffer(body))
    rr := httptest.NewRecorder()
    handlerFunc(rr, req)
    return rr
}

// Helper function to assert response status codes
func assertStatus(t *testing.T, rr *httptest.ResponseRecorder, expected int) {
    if rr.Code != expected {
        t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, expected)
    }
}

// Helper function to decode JSON responses
func decodeResponseBody(t *testing.T, body []byte, target interface{}) {
    err := json.Unmarshal(body, target)
    if err != nil {
        t.Errorf("failed to decode response body: %v", err)
    }
}

func TestRootRouteHandler(t *testing.T) {
    rr := performRequest(rootRouteHandler, "GET", "/", nil)

    assertStatus(t, rr, http.StatusOK)

    expected := "Welcome to the Go HTTP Server!\n"
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}

func TestRegisterMembership(t *testing.T) {
    member := Member{Name: "John Doe", Email: "john@example.com", StartDate: "2025-04-12"}
    body, _ := json.Marshal(member)

    rr := performRequest(registerMembership, "POST", "/memberships", body)

    assertStatus(t, rr, http.StatusOK)

    var response map[string]interface{}
    decodeResponseBody(t, rr.Body.Bytes(), &response)

    if response["message"] != "Membership registered successfully" {
        t.Errorf("unexpected response message: got %v", response["message"])
    }
}

func TestViewMembershipDetails(t *testing.T) {
    memberships["john@example.com"] = Member{Name: "John Doe", Email: "john@example.com", StartDate: "2025-04-12"}

    rr := performRequest(viewMembershipDetails, "GET", "/memberships?email=john@example.com", nil)

    assertStatus(t, rr, http.StatusOK)

    var response map[string]interface{}
    decodeResponseBody(t, rr.Body.Bytes(), &response)

    if response["message"] != "Membership retrieved successfully" {
        t.Errorf("unexpected response message: got %v", response["message"])
    }
}

func TestViewAllMemberships(t *testing.T) {
    memberships["john@example.com"] = Member{Name: "John Doe", Email: "john@example.com", StartDate: "2025-04-12"}
    memberships["jane@example.com"] = Member{Name: "Jane Smith", Email: "jane@example.com", StartDate: "2025-02-25"}

    rr := performRequest(viewAllMemberships, "GET", "/memberships", nil)

    assertStatus(t, rr, http.StatusOK)

    var members []map[string]string
    decodeResponseBody(t, rr.Body.Bytes(), &members)

    if len(members) != 2 {
        t.Errorf("unexpected number of memberships: got %d want %d", len(members), 2)
    }
}

func TestCancelMembership(t *testing.T) {
    memberships["john@example.com"] = Member{Name: "John Doe", Email: "john@example.com", StartDate: "2025-04-12"}

    rr := performRequest(cancelMembership, "DELETE", "/memberships?email=john@example.com", nil)

    assertStatus(t, rr, http.StatusOK)

    var response map[string]string
    decodeResponseBody(t, rr.Body.Bytes(), &response)

    if response["message"] != "Membership canceled successfully" {
        t.Errorf("unexpected response message: got %v", response["message"])
    }

    if _, exists := memberships["john@example.com"]; exists {
        t.Error("membership was not canceled")
    }
}

func TestModifyMembershipStartDate(t *testing.T) {
    memberships["john@example.com"] = Member{Name: "John Doe", Email: "john@example.com", StartDate: "2025-04-12"}
    body := `{"start_date": "2025-03-31"}`

    rr := performRequest(modifyMembershipStartDate, "PATCH", "/memberships?email=john@example.com", []byte(body))

    assertStatus(t, rr, http.StatusOK)

    var response map[string]interface{}
    decodeResponseBody(t, rr.Body.Bytes(), &response)

    if response["message"] != "Membership start date updated successfully" {
        t.Errorf("unexpected response message: got %v", response["message"])
    }
}
