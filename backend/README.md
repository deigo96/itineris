# 🚀 Project Name

> A short description of your project.

## 📌 Prerequisites

Ensure you have the following installed on your system:

- **Go** (Version 1.x) → [Install Go](https://go.dev/dl/)
- **Database** (PostgreSQL, MySQL, etc.)
- **Environment Variables** (Configured in `.env` file)

---

## ⚙️ Installation

### **1️⃣ Install Go**
Make sure Go is installed. You can check by running:

```sh
go version
```

If Go is not installed, download it [here](https://go.dev/dl/).

---

### **2️⃣ Create a `.env` File**
Inside the root directory, create a `.env` file and add your environment variables:

```ini
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

> ⚠️ **Note:** Make sure to update these values according to your setup.

---

### **3️⃣ Configure the Database & Services**
- Start your database PostgreSQL
- Ensure the database user has the required permissions.

---

### **4️⃣ Import the SQL Schema**
Run the following command to initialize your database:

```sh
psql -U your_user -d your_database -f schema.sql
```

---

### **5️⃣ Install Dependencies**
Run:

```sh
go mod tidy
```

This will fetch and install all necessary dependencies.

---

### **6️⃣ Run the Service**
To start the service, execute:

```sh
go run main.go
```

or build and run:

```sh
go build -o app && ./app
```

The service should now be running on `http://localhost:8080` 🎉.

---

## 📄 API Endpoints

| Method | Endpoint       | Description        |
|--------|--------------|------------------|
| `GET`  | `/ping`      | Health check    |
| `POST` | `/login`     | User login      |
| `POST` | `/logout`    | User logout     |

---

## 🛠 Troubleshooting

- **Error: `no such file or directory: .env`**  
  👉 Make sure the `.env` file is created and correctly configured.

- **Error: `database connection failed`**  
  👉 Ensure your database is running and credentials are correct.

---

## 🤝 Contributing

Contributions are welcome! Feel free to submit issues and pull requests.

---

## 📜 License

This project is licensed under the **MIT License**.

---

### 🎯 **Happy Coding!** 🚀

