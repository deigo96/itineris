# Leave Request API

## 📌 Table of Contents
- [📚 Introduction](#introduction)
- [🛠️ Installation](#installation)
- [⚙️ Environment Setup](#environment-setup)
- [📂 Database & Service Configuration](#database--service-configuration)
- [📝 API Endpoints](#api-endpoints)
  - [Create Leave Request](#create-leave-request)
  - [Get Leave Balance](#get-leave-balance)
- [🚀 Running the Service](#running-the-service)

---

## 📚 Introduction
This is a **simple Leave Request API** that allows employees to request leave, check leave balances, and manage approvals.

---

## 🛠️ Installation
1. Install **Go** (latest version) → [Download Go](https://go.dev/doc/install)
2. Clone this repository:
   ```sh
   git clone https://github.com/deigo96/itineris.git
   ```
3. Navigate to the project directory:
   ```sh
   cd itineris/backend
   ```
4. Install dependencies:
   ```sh
   go mod tidy
   ```

---

## ⚙️ Environment Setup
1. Create a `.env` file:
   ```sh
   touch .env
   ```
2. Add your environment variables:
   ```env
    SERVICE_NAME=""
    SERVICE_HOST=""
    SERVICE_PORT=""
    SECRET_KEY=""


    # database
    DB_HOST=""
    DB_PORT=""
    DB_USER=""
    DB_NAME=""
    DB_PASSWORD=""
   ```

---

## 📂 Database & Service Configuration
1. Import the database schema:
   ```sh
   psql -U postgres -d leave_db -f schema.sql
   ```
2. Configure the service in `config.yaml`.

---

## 📝 API Endpoints

### `url = {{BASE_URL}}/api/v1`

### 📌 Login
- **Endpoint:** `POST {{url}}/auth/login`
- **Description:** Login to get the token.
- **Headers:**
  ```json
    {
        "Content-Type": "application/json"
    }
  ```
- **Request Body:**
  ```json
    {
        "nip": "1234567", //numeric, required
        "password": "12345678" //required
    }
  ```
- **Response:**
  ```json
    {
        "code": "2000",
        "message": "success",
        "data": {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDA3NTc1NTMsImlkIjoyLCJuaXAiOiIxMjM0NTY3Iiwicm9sZSI6IlN0YWZmIn0.17lfvpQ1p-nCG_LRfCBclI8sNJ5DAMsafB7RwrwVuyM",
            "token_type": "Bearer"
        }
    }
  ```

### 📌 Get Employee
- **Endpoint:** `GET {{url}}/employees`
- **Description:** Get current employee.
- **Headers:**
  ```json
    {
        "Authorization": "Bearer <token>",
        "Content-Type": "application/json"
    }
  ```

- **Response:**
  ```json
    {
        "code": "2000",
        "message": "success",
        "data": {
            "id": 2,
            "name": "siahaan",
            "nip": "1234567",
            "leave_balance": 12,
            "role": "Staff",
            "total_pending_request": 0,
            "created_at": "2025-02-27 16:27:54.401362 +0700 +07",
            "created_by": "postgres",
            "updated_at": "2025-02-27 16:27:54.401362 +0700 +07",
            "updated_by": "postgres"
        }
    }
  ```

### 📌 Get Leave Type
- **Endpoint:** `GET {{url}}/employees/leave-type`
- **Description:** Get leave types.
- **Headers:**
  ```json
    {
        "Authorization": "Bearer <token>",
        "Content-Type": "application/json"
    }
  ```

- **Response:**
  ```json
    {
        "code": "2000",
        "message": "success",
        "data": [
            {
                "id": 1,
                "type_name": "Cuti Tahunan"
            },
            {
                "id": 2,
                "type_name": "Cuti Besar"
            }
        ]
    }
  ```

### 📌 Store Leave Request
- **Endpoint:** `POST {{url}}/leave-requests`
- **Description:** Submit leave request.
- **Headers:**
  ```json
    {
        "Authorization": "Bearer <token>",
        "Content-Type": "application/json"
    }
  ```
- **Request Body:**
  ```json
    {
        "start_date": "2025-02-28", //required
        "end_date": "2025-03-01", //required
        "reason": "sakit", //required
        "leave_type": 1 //required
    }
  ```
- **Response:**
  ```json
    {
        "code": "2000",
        "message": "success",
        "data": null
    }
  ```

### 📌 Get Histories Leave Request
- **Endpoint:** `GET /leave-requests`
- **Description:** Retrieves all leave requests.
- **Headers:**
  ```json
    {
        "Authorization": "Bearer <token>"
    }
  ```
- **Response:**
  ```json
    {
        "code": "2000",
        "message": "success",
        "data": [
            {
                "id": 13,
                "employee_id": 2,
                "status": "pending",
                "start_date": "2025-02-28",
                "end_date": "2025-03-01",
                "reason": "sakit",
                "rejection_note": "",
                "total_days": 2,
                "leave_type": "Cuti Tahunan",
                "created_at": "2025-02-27 22:57:01"
            },
            {
                "id": 12,
                "employee_id": 2,
                "status": "rejected",
                "start_date": "2025-02-28",
                "end_date": "2025-03-01",
                "reason": "sakit",
                "rejection_note": "",
                "total_days": 2,
                "leave_type": "Cuti Tahunan",
                "created_at": "2025-02-27 16:28:37"
            }
        ]
    }
  ```



### 📌 Get 1 History Leave Request
- **Endpoint:** `GET /leave-requests/{id}`
- **Description:** Retrieves leave request by id.
- **Headers:**
  ```json
    {
        "Authorization": "Bearer <token>"
    }
  ```
- **Response:**
  ```json
    {
        "code": "2000",
        "message": "success",
        "data": {
            "id": 12,
            "employee_id": 2,
            "status": "rejected",
            "start_date": "2025-02-28",
            "end_date": "2025-03-01",
            "reason": "sakit",
            "rejection_note": "",
            "total_days": 2,
            "leave_type": "Cuti Tahunan",
            "created_at": "2025-02-27 16:28:37"
        }
    }
  ```

### 📌 Approval Action
- **Endpoint:** `POST {{url}}/leave-requests/action`
- **Description:** Action to approve or reject request.
- **Notes:**
    - Approve = 1
    - Reject = 2
- **Headers:**
  ```json
    {
        "Authorization": "Bearer <token>",
        "Content-Type": "application/json"
    }
  ```
- **Request Body:**
  ```json
    {
        "id": 9, // required
        "status": 2, // required
        "rejection_note": "" // optional
    }
  ```
- **Response:**
  ```json
    {
        "code": "2000",
        "message": "success",
        "data": null
    }
  ```

---

## 🚀 Running the Service
Run the application using:
```sh
go run main.go
```
Then open the API at:
```
http://localhost:8080
```

