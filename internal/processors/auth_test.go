package processors

import (
	"complaint_service/internal/config"
	"complaint_service/internal/models"
	"complaint_service/internal/repository"
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v2/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	cfg := config.NewConfig()
	fmt.Printf(cfg.Env)

	db, err := repository.NewPostgresDB()

	if err != nil {
		log.Error("Create database error: %v ", err)
	}
	repo := repository.CreateComplaintsRepository(db)
	service := CreateComplaintsProcessor(repo)

	db.Query("DELETE FROM users WHERE username = 'TestUser1';")

	testUser1 := models.UserSignUp{
		UserName: "TestUser1",
		Password: "PasswordTestUser1",
		Role:     models.User,
	}

	result, err := service.CreateUser(testUser1)

	require.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.NotNil(t, result)

	testUser2 := models.UserSignUp{
		UserName: "",
		Password: "PasswordTestUser2",
		Role:     models.User,
	}

	result, err = service.CreateUser(testUser2)
	assert.Error(t, err)

	result, err = service.CreateUser(testUser1)
	assert.Error(t, err)

	testUser3 := models.UserSignUp{
		UserName: "TestUser1",
		Password: "",
		Role:     models.User,
	}

	result, err = service.CreateUser(testUser3)
	assert.Error(t, err)

	testUser4 := models.UserSignUp{
		UserName: "TestUser4",
		Password: "PasswordTestUser4",
	}

	result, err = service.CreateUser(testUser4)

	require.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.NotNil(t, result)

}

func TestGetToken(t *testing.T) {
	cfg := config.NewConfig()
	fmt.Printf(cfg.Env)

	db, err := repository.NewPostgresDB()

	if err != nil {
		log.Error("Create database error: %v ", err)
	}
	repo := repository.CreateComplaintsRepository(db)
	service := CreateComplaintsProcessor(repo)

	testUser1 := models.UserSignUp{
		UserName: "TestUser1",
		Password: "PasswordTestUser1",
	}

	result, err := service.GetToken(testUser1.UserName, testUser1.Password)

	require.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.NotNil(t, result)

	testUser2 := models.UserSignUp{
		UserName: "TestUser2",
		Password: "PasswordTestUser2",
	}

	result, err = service.GetToken(testUser2.UserName, testUser2.Password)
	assert.Error(t, err)

	testUser3 := models.UserSignUp{
		UserName: "",
		Password: "PasswordTestUser2",
	}

	result, err = service.GetToken(testUser3.UserName, testUser3.Password)
	assert.Error(t, err)

	testUser4 := models.UserSignUp{
		UserName: "TestUser4",
		Password: "",
	}

	result, err = service.GetToken(testUser4.UserName, testUser4.Password)
	assert.Error(t, err)
}
