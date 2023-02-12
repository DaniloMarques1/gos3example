package model

type User struct {
	Name      string `bson:"name"`
	Email     string `bson:"email"`
	Phone     string `bson:"phone"`
	Bio       string `bson:"bio"`
	ResumeUrl string `bson:"resume_url"`
}
