package user

//
//import (
//	"crud/entity"
//	"github.com/stretchr/testify/assert"
//	"testing"
//	"time"
//)
//
//func TestUseCaseUser_CreateUser(t *testing.T) {
//	createUser := useCaseUser{}
//	user := UserParam{
//		Name:     "john",
//		Email:    "john@example.com",
//		Password: "123",
//	}
//	callUser := entity.User{
//		ID:        0,
//		Name:      user.Name,
//		Email:     user.Email,
//		Password:  user.Password,
//		CreatedAt: time.Time{},
//		UpdatedAt: time.Time{},
//	}
//	expected := entity.User{
//		Name:     "john",
//		Email:    "john@example.com",
//		Password: "123",
//	}
//
//	result, err := createUser.CreateUser(user)
//
//	assert.Nil(t, err)
//	assert.Equal(t, expected, result)
//}
