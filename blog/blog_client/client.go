package main

import (
	"context"
	"fmt"
	"github.com/grpc-project/blog/blogpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Blog Client")
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Sang Jinsu",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}
	res, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{
		Blog: blog,
	})
	if err != nil {
		log.Fatalf("Unexcepted error: %v\n", err)
		return
	}
	fmt.Printf("Blog has been created: %v\n", res.GetBlog())
	blogId := res.GetBlog().GetId()

	fmt.Println("Reading the blog")
	_, readErr := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: "asdfasdf",
	})
	if readErr != nil {
		fmt.Printf("Error happend while reading %v\n", readErr)
	}

	readBlogReq := &blogpb.ReadBlogRequest{
		BlogId: blogId,
	}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		log.Fatalf("Error happend while reading %v\n", readBlogErr)
	}

	fmt.Printf("Blog was read %v\n", readBlogRes)

}
