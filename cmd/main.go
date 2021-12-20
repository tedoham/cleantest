package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tedoham/cleantest/internal/repository"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	client := repository.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	// create a post
	createdPost, err := client.Course.CreateOne(
		repository.Course.Title.Set("Hi from Prisma!"),
		repository.Course.Description.Set("description here"),
		repository.Course.Price.Set(34.0),
	).Exec(ctx)

	if err != nil {
		return err
	}

	result, _ := json.MarshalIndent(createdPost, "", "  ")
	fmt.Printf("created post: %s\n", result)

	// find a single post
	post, err := client.Course.FindUnique(
		repository.Course.ID.Equals(createdPost.ID),
	).Exec(ctx)
	if err != nil {
		return err
	}

	result, _ = json.MarshalIndent(post, "", "  ")
	fmt.Printf("post: %s\n", result)

	// for optional/nullable values, you need to check the function and create two return values
	// `desc` is a string, and `ok` is a bool whether the record is null or not. If it's null,
	// `ok` is false, and `desc` will default to Go's default values; in this case an empty string (""). Otherwise,
	// `ok` is true and `desc` will be "my description".
	// desc, ok := post.Description()
	// if !ok {
	// 	return fmt.Errorf("post's description is null")
	// }

	// fmt.Printf("The posts's description is: %s\n", desc)

	return nil
}
