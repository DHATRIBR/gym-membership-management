# Gym Membership Management System

A lightweight HTTP server built with Go to manage gym memberships. This project provides basic functionalities for registering, retrieving, and updating memberships, as well as handling cancellations.

---

## Features

- **Register Membership**: Add new members to the gym with their name, email, and membership start date.
- **View Membership Details**: Retrieve specific membership information using the member's email.
- **View All Memberships**: List all active members with their name and email.
- **Modify Membership Start Date**: Update the start date of a membership.
- **Cancel Membership**: Delete a member from the system.

---

## API Endpoints

| Method | Endpoint                        | Description                           |
|--------|---------------------------------|---------------------------------------|
| GET    | `/`                             | Welcome message for the server.       |
| POST   | `/memberships`                  | Register a new membership.            |
| GET    | `/memberships?email={email}`    | Retrieve membership details.          |
| GET    | `/memberships`                  | List all active memberships.          |
| PATCH  | `/memberships?email={email}`    | Modify membership start date.         |
| DELETE | `/memberships?email={email}`    | Cancel a membership.                  |

---

## **Request & Response Format**

### **1. GET `/`**
**Description**: Welcome route of the HTTP server.

**Request**:  
No request body needed.

**Response**:  
HTTP Status: `200 OK`  
```plaintext
Welcome to the Go HTTP Server!
```

---

### **2. POST `/memberships`**
**Description**: Register a new membership.

**Request**:
```json
{
    "name": "Name Value",
    "email": "email@example.com",
    "start_date": "2025-03-31"
}
```

**Response (Success)**:  
HTTP Status: `200 OK`  
```json
{
    "message": "Membership registered successfully",
    "membership_details": {
        "name": "Name Value",
        "email": "email@example.com",
        "start_date": "2025-03-31"
    }
}
```

**Response (Invalid Start Date)**:  
HTTP Status: `400 Bad Request`  
```json
{
    "message": "Invalid start_date format. Expected YYYY-MM-DD"
}
```

**Response (Invalid Request Body)**:  
HTTP Status: `400 Bad Request`  
```json
{
    "message": "Invalid request body"
}
```

---

### **3. GET `/memberships?email={email}`**
**Description**: Retrieve membership details by email.

**Request**:  
No request body needed. Pass the email as a query parameter.

Example:
```
http://localhost:8080/memberships?email=email@example.com
```

**Response (Success)**:  
HTTP Status: `200 OK`  
```json
{
    "message": "Membership retrieved successfully",
    "membership_details": {
        "name": "Name Value",
        "email": "email@example.com",
        "start_date": "2025-03-31"
    }
}
```

**Response (Membership Not Found)**:  
HTTP Status: `404 Not Found`  
```json
{
    "message": "Membership not found"
}
```

**Response (Missing Email)**:  
HTTP Status: `400 Bad Request`  
```json
{
    "message": "Email query parameter is required"
}
```

---

### **4. GET `/memberships`**
**Description**: List all active memberships.

**Request**:  
No request body needed.

**Response (Memberships Exist)**:  
HTTP Status: `200 OK`  
```json
[
    {
        "name": "Name Value",
        "email": "email@example.com"
    },
    {
        "name": "Name Another",
        "email": "jane@example.com"
    }
]
```

**Response (No Memberships Found)**:  
HTTP Status: `200 OK`  
```json
{
    "message": "No memberships found"
}
```

---

### **5. DELETE `/memberships?email={email}`**
**Description**: Cancel a membership by email.

**Request**:  
No request body needed. Pass the email as a query parameter.

Example:
```
http://localhost:8080/memberships?email=email@example.com
```

**Response (Success)**:  
HTTP Status: `200 OK`  
```json
{
    "message": "Membership canceled successfully"
}
```

**Response (Membership Not Found)**:  
HTTP Status: `404 Not Found`  
```json
{
    "message": "Membership not found"
}
```

**Response (Missing Email)**:  
HTTP Status: `400 Bad Request`  
```json
{
    "message": "Email query parameter is required"
}
```

---

### **6. PATCH `/memberships?email={email}`**
**Description**: Update the membership start date.

**Request**:  
Pass the email as a query parameter. Include the new start date in the request body.

Example:
Query:
```
http://localhost:8080/memberships?email=email@example.com
```

Request Body:
```json
{
    "start_date": "2025-04-31"
}
```

**Response (Success)**:  
HTTP Status: `200 OK`  
```json
{
    "message": "Membership start date updated successfully",
    "membership_details": {
        "name": "Name Value",
        "email": "email@example.com",
        "start_date": "2025-04-31"
    }
}
```

**Response (Invalid Request Body)**:  
HTTP Status: `400 Bad Request`  
```json
{
    "message": "Invalid request body"
}
```

**Response (Membership Not Found)**:  
HTTP Status: `404 Not Found`  
```json
{
    "message": "Membership not found"
}
```

**Response (Missing Email)**:  
HTTP Status: `400 Bad Request`  
```json
{
    "message": "Email query parameter is required"
}
```

---

### **Notes**
- **Date Validation**: Dates must follow the `YYYY-MM-DD` format for all APIs that handle dates.
- **Error Responses**: All error responses are consistent and include a `message` field with a detailed description.

---

### **Error Responses**
Consistent error responses are returned with HTTP error codes:
```json
{
    "message": "Error description here"
}
```

---

## Getting Started

### Prerequisites
- **Go** installed on your system (version 1.16 or higher recommended).

### Setup
1. Clone this repository:
   ```bash
   git clone https://github.com/your-username/gym-membership-management.git
   cd gym-membership-management
   ```

2. Run the server:
   ```bash
   go run main.go
   ```

3. Access the server at [http://localhost:8080](http://localhost:8080).

---

## Testing
Run unit tests for the application:
```bash
go test -v
```

---
