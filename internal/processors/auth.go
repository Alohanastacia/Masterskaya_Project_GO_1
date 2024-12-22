package processors

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/repository"
	"crypto/sha256"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

const salt = "afdafadfadfadf"

type Authorization interface {
	CreateUser(user entity.User) (int, error)
}

type AuthService struct {
	repo repository.Authorization
}

// Функция NewAuthService является конструктором структуры AuthService. Принимает на вход переменную типа repository.Authorization и возвращает AuthService.
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

/*
	Функция CreateUser проверяет на корректность полученные от пользователя данные и вызывает функцию repo.CreateUser для создания пользователя. Принимает на вход структуру User,

возвращает id типа int и ошибку типа error.
*/
func (s *AuthService) CreateUser(user entity.User) (int, error) {
	user.User_UUID = uuid.NewV4()
	if len(user.Password) == 0 || len(user.Username) == 0 {
		return 0, fmt.Errorf("имя пользователя или пароль не могут быть пустыми")
	}
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

/*
Функция generatePasswordHash создаёт хеш пороля. Принимает на вход переменную password типа string возвращает хешированный пароль типа string.
*/
func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
