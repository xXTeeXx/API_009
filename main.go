package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/xXTeeXx/go-gorm/db"
	"github.com/xXTeeXx/go-gorm/models"

	"github.com/gin-contrib/cors"
	"time"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Read database configuration from environment variables
	dbType := os.Getenv("DB_TYPE")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Connect to the database
	database, err := db.ConnectDatabase(dbType, dbUser, dbPassword, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// AutoMigrate the database
	err = database.AutoMigrate(&models.Student{}, &models.Subject{}, &models.User{}, &models.Teacher{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Create repositories for each model
	studentRepo := models.NewStudentRepository(database)
	subjectRepo := models.NewSubjectRepository(database)
	teacherRepo := models.NewTeacherRepository(database)

	// Initialize Gin router
	r := gin.Default()

	// Set CORS (Cross-Origin Resource Sharing)
	r.Use(cors.New(cors.Config{
		// 3000 is the port where frontend react runs
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	
	// Student routes
	r.GET("/students", studentRepo.GetStudents)
	r.POST("/students", studentRepo.CreateStudent)
	r.GET("/students/:id", studentRepo.GetStudent)
	r.PUT("/students/:id", studentRepo.UpdateStudent)
	r.DELETE("/students/:id", studentRepo.DeleteStudent)

	// Subject routes
	r.GET("/subjects", subjectRepo.GetSubjects)
	r.POST("/subjects", subjectRepo.CreateSubject)
	r.GET("/subjects/:id", subjectRepo.GetSubject)
	r.PUT("/subjects/:id", subjectRepo.UpdateSubject)
	r.DELETE("/subjects/:id", subjectRepo.DeleteSubject)

	// Teacher routes
	r.GET("/teachers", teacherRepo.GetTeachers)
	r.POST("/teachers", teacherRepo.CreateTeacher)
	r.GET("/teachers/:id", teacherRepo.GetTeacher)
	r.PUT("/teachers/:id", teacherRepo.UpdateTeacher)
	r.DELETE("/teachers/:id", teacherRepo.DeleteTeacher)

	// Create userRepo variable to use UserRepository
	userRepo := models.NewUserRepository(database)

	// Route to get all users
	r.GET("/users", userRepo.GetUsers)

	// Route to create a user
	r.POST("/users", userRepo.PostUser)

	// Route to get a user by email
	r.GET("/users/:email", userRepo.GetUser)

	// Route to update a user by email
	r.PUT("/users/:email", userRepo.UpdateUser)

	// Route to delete a user by email
	r.DELETE("/users/:email", userRepo.DeleteUser)

	// Route to login
	r.POST("/users/login", userRepo.Login)

	// 404 route
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	// Run the server
	if err := r.Run(":5000"); err != nil {
		log.Fatalf("Server is not running: %v", err)
	}
}
