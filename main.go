package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type Member struct {
    Name      string `json:"name"`
    Email     string `json:"email"`
    StartDate string `json:"start_date"`
}

var memberships = map[string]Member{} // In-memory store

func setJSONResponseHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
func isValidDate(dateStr string) bool {
    // Define the expected date format
    layout := "2006-01-02" // Go's standard format for parsing dates
    _, err := time.Parse(layout, dateStr)
    return err == nil // If err is nil, the date is valid
}
func rootRouteHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the Go HTTP Server!")
}

func registerMembership(w http.ResponseWriter, r *http.Request) {
	setJSONResponseHeader(w)
    var member Member
    if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
        http.Error(w, `{"message": "Invalid request body"}`, http.StatusBadRequest)
        return
    }
    if !isValidDate(member.StartDate) {
        http.Error(w, `{"message": "Invalid start_date format. Expected YYYY-MM-DD"}`, http.StatusBadRequest)
        return
    }
    _, exists := memberships[member.Email]
    if exists {
        http.Error(w, `{"message": "Membership already exists with this email"}`, http.StatusBadRequest)
        return
    }
    memberships[member.Email] = member
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Membership registered successfully",
        "membership_details": member,
    })
}

func viewMembershipDetails(w http.ResponseWriter, r *http.Request) {
	setJSONResponseHeader(w)
    email := r.URL.Query().Get("email")
    member, exists := memberships[email]
    if !exists {
        http.Error(w, `{"message": "Membership not found"}`, http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Membership retrieved successfully",
        "membership_details": member,
    })
}

func viewAllMemberships(w http.ResponseWriter, r *http.Request) {
	setJSONResponseHeader(w)
    type MemberResponse struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }
	if len(memberships) == 0 {
        json.NewEncoder(w).Encode(map[string]string{
            "message": "No memberships found",
        })
        return
    }
    var members []MemberResponse
    for _, member := range memberships {
        members = append(members, MemberResponse{
            Name:  member.Name,
            Email: member.Email,
        })
    }
    json.NewEncoder(w).Encode(members)
}

func cancelMembership(w http.ResponseWriter, r *http.Request) {
	setJSONResponseHeader(w)
    email := r.URL.Query().Get("email")
    if _, exists := memberships[email]; !exists {
        http.Error(w, `{"message": "Membership not found"}`, http.StatusNotFound)
        return
    }
    delete(memberships, email)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Membership canceled successfully",
    })
}

func modifyMembershipStartDate(w http.ResponseWriter, r *http.Request) {
	setJSONResponseHeader(w)
    email := r.URL.Query().Get("email")
    var data struct {
        NewStartDate string `json:"start_date"`
    }
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, `{"message": "Invalid request body"}`, http.StatusBadRequest)
        return
    }
    member, exists := memberships[email]
    if !exists {
        http.Error(w, `{"message": "Membership not found"}`, http.StatusNotFound)
        return
    }
    if !isValidDate(data.NewStartDate) {
        http.Error(w, `{"message": "Invalid start_date format. Expected YYYY-MM-DD"}`, http.StatusBadRequest)
        return
    }
    member.StartDate = data.NewStartDate
    memberships[email] = member
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Membership start date updated successfully",
        "membership_details": member,
    })
}

func handler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "POST":
        registerMembership(w, r)
    case "GET":
        email := r.URL.Query().Get("email")
        if email != "" {
            viewMembershipDetails(w, r)
        } else {
			viewAllMemberships(w, r)
        }
    case "DELETE":
        cancelMembership(w, r)
    case "PATCH":
        modifyMembershipStartDate(w, r)
    default:
        
        http.Error(w, `{"message": "Method not allowed"}`, http.StatusMethodNotAllowed)
    }
}

func main() {
    http.HandleFunc("/", rootRouteHandler)
    http.HandleFunc("/memberships", handler)
	http.HandleFunc("/memberships/", handler)

    fmt.Println("Server is running on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server: ", err)
    }
}
