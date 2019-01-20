package main

import (
	"context"
	"fmt"
	"log"

	"github.com/FernandoDevBh/grpc-go-course/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	id, _ := doCreateBlog(c)

	doReadBlog(c, id)
}

func doCreateBlog(c blogpb.BlogServiceClient) (string, error) {
	fmt.Println("Creating a blog")
	blogRequest := &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			AuthorId: "Fernando",
			Title:    "My first blog",
			Content:  "Content for my first blog",
		},
	}
	createBlogRes, err := c.CreateBlog(context.Background(), blogRequest)
	if err != nil {
		log.Fatalf("Unexpected Error: %v", err)
		return "", err
	}
	fmt.Printf("Blog has been created: %v", createBlogRes)

	return createBlogRes.Blog.GetId(), nil
}

func doReadBlog(c blogpb.BlogServiceClient, id string) {
	fmt.Println("Reading a blog")
	_, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: "555233",
	})

	if err != nil {
		fmt.Printf("Erro happended while reading: %v\n", err)
	}

	res, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: id,
	})

	if err != nil {
		fmt.Printf("Erro happended while reading: %v\n", err)
	}

	fmt.Printf("Blog was read: %v\n", res.GetBlog())
}
