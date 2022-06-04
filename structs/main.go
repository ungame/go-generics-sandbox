package main

import (
	"fmt"
	"github.com/ungame/go-generics-sandbox/structs/fake"
	"github.com/ungame/go-generics-sandbox/structs/memory"
	"github.com/ungame/go-generics-sandbox/structs/models"
	"strings"
)

func main() {
	users := fake.GetUsers()
	comments := fake.GetComments(users)

	mu := memory.New[*models.User]()

	for _, user := range users {
		mu.Set(user.ID, user)
	}

	mc := memory.New[*models.Comment]()

	for _, comment := range comments {
		mc.Set(comment.Key(), comment)
	}

	title("USERS")
	userId := users[0].ID
	user := mu.Get(userId)
	fmt.Println(user)

	user = mu.Get("?")
	fmt.Println(user)

	title("COMMENTS")
	commentKey := comments[0].Key()
	comment := mc.Get(commentKey)
	fmt.Println(comment)

	comment = mc.Get("?")
	fmt.Println(comment)
}

func title(t string) {
	fmt.Printf("\n%s\n%s\n%s\n", bar(10), t, bar(10))
}

func bar(size int) string {
	return strings.Repeat("=", size)
}
