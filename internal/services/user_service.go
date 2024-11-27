package services

import (
	"blog/config"
	"blog/internal/models"
	"blog/internal/repositories"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

var ErrUNF error = errors.New("user not found")
var ErrIVC error = errors.New("invalid verification code")

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.UserRepo.GetByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := generateJWT(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func generateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	secret := config.AppConfig.JWT.SecretKey
	return token.SignedString([]byte(secret))
}

func (s *UserService) RegisterUser(user *models.User) error {
	exists, err := s.UserRepo.EmailExists(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email already exists")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPass)

	verificationCode, err := generateVerificationCode()
	if err != nil {
		return err
	}
	user.VerificationCode = verificationCode
	user.IsVerified = false

	if err := sendVerificationEmail(user.Email, verificationCode); err != nil {
		return errors.New("failed to send verification email" + err.Error())
	}

	if err := s.UserRepo.Create(user); err != nil {
		return err
	}

	return nil
}

func generateVerificationCode() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func sendVerificationEmail(email, code string) error {
	d := gomail.NewDialer(
		config.AppConfig.Email.SMTPServer,
		config.AppConfig.Email.SMTPPort,
		config.AppConfig.Email.Username,
		config.AppConfig.Email.Password,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	message := gomail.NewMessage()
	message.SetHeader("From", "thunderleo@mail.ru")
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Email Verification")
	message.SetBody("text/plain", "Your verification code is: "+code)

	return d.DialAndSend(message)
}

func (s *UserService) VerifyEmail(email, code string) error {
	user, err := s.UserRepo.GetByEmail(email)
	if err != nil {
		return ErrUNF
	}

	if user.VerificationCode != code {
		return ErrIVC
	}

	user.IsVerified = true
	user.VerificationCode = ""
	return s.UserRepo.Update(user)
}
