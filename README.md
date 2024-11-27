# Blog

This project is an educational blog with a **Go** backend, communicating via HTTP requests. All interactions are performed using `POST`, `GET`, and `DELETE` methods. The project includes functionality for user and post management, as well as email confirmation for registration.

---

## 📋 Features

### Key functionalities:
- **User Management:**
  - Registration with email confirmation via Mail.ru.
  - Password hashing using bcrypt.
- **Post Management:**
  - Create, edit, and delete posts.
  - Like functionality for posts.
- **Security:**
  - Protected API endpoints.
  - Input validation.

---

## 🛠 Technology Stack

### Backend:
- Language: Go
- ORM: GORM
- Database: PostgreSQL
- Libraries:
  - `gorilla/mux` for routing.
  - `golang.org/x/crypto/bcrypt` for password hashing.
  - `gomail` for email handling.
  - `viper` for environment variables.

---

## 🚀 Getting Started

### Prerequisites:
- Go 1.19+ installed.
- PostgreSQL database set up.

### Steps to Run the Project:
1. Clone the repository:
   ```bash
   git clone https://github.com/Oxeeee/go-blog.git
   cd go-blog
   ```

2. Set up environment variables:
   Edit a `./config/config.yaml` file with your database and email settings.

3. Build and run the application:
   ```bash
   go mod tidy
   go run cmd/main.go
   ```

4. Use your preferred HTTP client (e.g., Postman, cURL) to test the endpoints.

---

## 🛡 Security Considerations

- User passwords are hashed using **bcrypt**.
- Input validation is implemented to prevent SQL injection and XSS attacks.

---

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

---

## 🤝 Contributing

Contributions are welcome! Please open an issue or submit a pull request for any bugs, features, or suggestions.

---

Let me know if you’d like to refine this further!