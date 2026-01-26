package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	ID    int    `json:"id"` // json:"" used for data cloumn
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	// dsn := "user:password@tcp(host:port)/database"
	dsn := "app_user:app_pass@tcp(127.0.0.1:3606)/demo?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect the database ---> ", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Database ping failed ---> ", err)
	}

	createTable := `
		CREATE TABLE IF NOT EXISTS users (
			id    INT AUTO_INCREMENT PRIMARY KEY,
			name  VARCHAR(100) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			age   INT
		);
	`

	if _, err := db.Exec(createTable); err != nil {
		log.Fatal("Failed to create db....", err)
	}

	e.POST("/users", func(c echo.Context) error {
		// 1. Initialize the User struct and bind the request body
		u := new(User) // Assuming 'User' is a struct like {ID int, Name string, Email string, Age int}
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"Error": "Invalid request body"})
		}

		// 2. Execute the INSERT query using Exec for MySQL
		result, err := db.Exec(
			"INSERT INTO users (name, email, age) VALUES (?, ?, ?)",
			u.Name, u.Email, u.Age,
		)

		if err != nil {
			// Handle database errors
			return c.JSON(http.StatusInternalServerError, map[string]string{"Error": "Database error: " + err.Error()})
		}

		// 3. Get the last inserted ID
		lastID, err := result.LastInsertId()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"Error": "Could not get last insert ID: " + err.Error()})
		}

		// 4. Assign the ID to the User struct
		u.ID = int(lastID) // Convert int64 to int

		// 5. Return the created user with a 201 status
		return c.JSON(http.StatusCreated, u)
	})

	e.GET("/users", func(c echo.Context) error {
		rows, err := db.Query("Select id,name,email,age from users")

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"Error": "Database error: " + err.Error()})
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			users = append(users, u)
		}

		return c.JSON(http.StatusOK, users)

	})

	e.GET("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		var u User

		err := db.QueryRow("select * from users where id=?", id).Scan(&u.ID, &u.Name, &u.Email, &u.Age)

		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"Error": "User not present"})
		}

		return c.JSON(http.StatusOK, u)

	})

	// put create new entry at same id not update existing
	e.PUT("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		u := new(User)

		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"Error": "Invalid request body"})
		}

		result, err := db.Exec("update users set name=?,email=?,age=? where id=?", u.Name, u.Email, u.Age, id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"Error": err.Error()})
		}

		rowsAffected, _ := result.RowsAffected()

		if rowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]string{"Error": "User not found"})
		}

		u.ID = id

		return c.JSON(http.StatusOK, u)
	})

	e.DELETE("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		result, err := db.Exec("delete from users where id=?", id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"Error": err.Error()})
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]string{"Error": "User not found"})
		}

		return c.JSON(http.StatusOK, map[string]string{"Ok": "Users deleted successfully"})
	})

	e.Logger.Fatal(e.Start(":8090"))
}
