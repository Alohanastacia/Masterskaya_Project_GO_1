package processors

import (
	"complaint_service/internal/config"
	"complaint_service/internal/models"
	"complaint_service/internal/repository"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
)

const (
	tokenTTL   = time.Hour * 12
	expiration = 3600
)

type tokenClaims struct {
	jwt.StandardClaims
	User_UUID uuid.UUID `json:"user_UUID"`
}

type Authorization interface {
	CreateUser(user models.UserSignUp) (int, error)
	GetToken(username, password string) (string, error)
}

type AuthService struct {
	repo         repository.Authorization
	SessionCache repository.SessionCache
}

// NewAuthService является конструктором структуры AuthService. Принимает на вход переменную типа repository.Authorization и возвращает AuthService
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo:         repo,
		SessionCache: *repository.NewSessionCache()}
}

/*
CreateUser проверяет на корректность полученные от пользователя данные и вызывает функцию
repo.CreateUser для создания пользователя. Принимает на вход структуру User,
возвращает id типа int и ошибку типа error
*/
func (s *AuthService) CreateUser(user models.UserSignUp) (int, error) {
	user.UserUUID = uuid.NewV4()
	if len(user.Password) == 0 || len(user.UserName) == 0 {
		return 0, fmt.Errorf("имя пользователя или пароль не могут быть пустыми")
	}
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

/*
GenerateToken генерирует JWT токен пользователя. Принимает на вход структуру username и password,
возвращает JWT токен типа string и ошибку.
*/
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	configs, err := config.LoadEnv()
	if err != nil {
		fmt.Println(err)
	}
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.UserUUID,
	})

	return token.SignedString([]byte(configs.JwtSigningKey))
}

/*
GetToken получает JWT токен из функции GetToken и сохраняет в кеш. Принимает на вход структуру username и password,
возвращает JWT токен типа string и ошибку.
*/
func (s *AuthService) GetToken(username, password string) (string, error) {
	if len(password) == 0 || len(username) == 0 {
		return "", fmt.Errorf("имя пользователя или пароль не могут быть пустыми")
	}
	token, err := s.GenerateToken(username, password)
	if err != nil {
		return "", err
	}

	password = generatePasswordHash(password)

	value, err := json.Marshal(&models.UserSessions{
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
	})

	if err != nil {
		return "", err
	}

	err = s.SessionCache.Set(token, value, int32(expiration))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) ParseToken(token string) (uuid.UUID, error) {
	result, err := s.SessionCache.Get(token)
	if err != nil {
		return uuid.Nil, err
	}

	if result != nil {
		userId, err := ParseJWT(token)
		if err != nil {
			return uuid.Nil, err
		}
		return userId, nil
	}

	return uuid.Nil, fmt.Errorf("срок действия сессии истёк")
}

func ParseJWT(accessToken string) (uuid.UUID, error) {
	configs, err := config.LoadEnv()
	if err != nil {
		fmt.Println(err)
	}
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(configs.JwtSigningKey), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return uuid.Nil, fmt.Errorf("token claims are not of type *tokenClaims")
	}

	return claims.User_UUID, nil
}

/*
generatePasswordHash создаёт хеш пороля. Принимает на вход переменную password типа string возвращает хешированный пароль типа string
*/

func generatePasswordHash(password string) string {
	configs, err := config.LoadEnv()
	if err != nil {
		fmt.Println(err)
	}
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(configs.JwtSalt)))
}
