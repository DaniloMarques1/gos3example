package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/danilomarques1/gos3example/api/model"
	"github.com/danilomarques1/gos3example/api/repository"
	"github.com/danilomarques1/gos3example/api/service"
)

type CreateUserDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Bio   string `json:"bio"`
	Phone string `json:"phone"`
}

type UserController struct {
	repository repository.UserRepository
	s3Service  *service.S3Service
}

func NewUserController() (*UserController, error) {
	repository, err := repository.NewMongoUserRepository()
	if err != nil {
		return nil, err
	}
	s3Service := service.NewS3()
	return &UserController{repository, s3Service}, nil
}

const MaxFileSize = 100000000

func (u UserController) Save(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(MaxFileSize); err != nil {
		// TODO
		log.Printf("Error parsing file %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	file, fileHandler, err := r.FormFile("resume")
	if err != nil {
		// TODO
		log.Printf("Error getting file %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	log.Printf("File name %v\n", fileHandler.Filename)

	// TODO add struct validation
	userDto := &CreateUserDto{}
	userDto.Name = r.PostForm.Get("name")
	userDto.Bio = r.PostForm.Get("bio")
	userDto.Email = r.PostForm.Get("email")
	userDto.Phone = r.PostForm.Get("phone")

	if err := u.s3Service.PutObject(fileHandler.Filename, file); err != nil {
		// TODO
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error sending to s3 %v\n", err)
		return
	}

	user := &model.User{
		Name:      userDto.Name,
		Email:     userDto.Email,
		Bio:       userDto.Bio,
		Phone:     userDto.Phone,
		ResumeUrl: fmt.Sprintf("http://localhost:4566/%v/%v", os.Getenv("S3_BUCKET_NAME"), fileHandler.Filename),
	}

	if err := u.repository.Save(user); err != nil {
		// TODO
		log.Printf("Error saving to mongo %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (u UserController) FindByEmail(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	email := queryParams.Get("email") // TODO: validate
	user, err := u.repository.FindByEmail(email)
	if err != nil {
		// TODO
		log.Printf("Error fetching user %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}
