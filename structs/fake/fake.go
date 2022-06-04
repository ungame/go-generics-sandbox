package fake

import (
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/ungame/go-generics-sandbox/structs/models"
	"math/rand"
	"time"
)

func NewUser() *models.User {
	return &models.User{
		ID:        uuid.NewString(),
		Email:     faker.Email(),
		Password:  faker.Password(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewComment(userID string) *models.Comment {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return &models.Comment{
		ID:        r.Int63(),
		UserID:    userID,
		Text:      faker.Paragraph(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func GetUsers() []*models.User {
	return []*models.User{
		NewUser(),
		NewUser(),
		NewUser(),
		NewUser(),
		NewUser(),
	}
}

func GetComments(users []*models.User) []*models.Comment {
	comments := make([]*models.Comment, 0, len(users))
	for _, user := range users {
		comments = append(comments, NewComment(user.ID))
	}
	return comments
}
