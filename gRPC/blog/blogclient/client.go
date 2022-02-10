package main

import (
	"context"
	"fmt"
	"gRPC/blog/blogpb"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	fmt.Println("Hello im a blog client")

	tls := false

	opts := grpc.WithInsecure()

	if tls {
		certFile := "D:/Udemy/go/gRPC/ssl/ca.crt" // CA trust certificate

		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificates : %v", sslErr)
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	ls, err := grpc.Dial("localhost:5000", opts)
	if err != nil {
		log.Fatalf("Failed to dial to localhost:5000 with error %v", err)
	}

	// makes sure the client connection is closed at the end
	defer ls.Close()
	c := blogpb.NewBlogServiceClient(ls)

	// Create Blog

	blog := &blogpb.Blog{
		AuthorId: "Stephanie",
		Title:    "My seventh Blog",
		Content:  "Content of the seventh blog post",
	}

	// fmt.Printf("Created client %f", c)
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error %v", err)
	}
	fmt.Println("Blog has been created: ", createBlogRes)
	blogId := createBlogRes.GetBlog().GetId()

	fmt.Println("the blog id is ", blogId)
	// Read Blog

	fmt.Println("Reading the blog")
	// _, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "I129379ds"})

	// if err2 != nil {
	// 	fmt.Println("Error while reading the data ", err2)
	// }

	res, err3 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: blogId})

	if err3 != nil {
		fmt.Println("Error while reading the data ", err3)
	}

	fmt.Println("The blog was read", res)

	// update blog

	blog2 := &blogpb.Blog{
		Id:       blogId,
		AuthorId: "Hitesh Kumar",
		Title:    "My first Blog Deleted",
		Content:  "Content of the first blog post",
	}

	res1, err4 := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: blog2})

	if err4 != nil {
		fmt.Println("Error while reading the data ", err4)
	}

	fmt.Println("Blog was read", res1)

	// delete blog

	deleteRes, err5 := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogId})

	if err5 != nil {
		fmt.Println("Error while reading the data ", err5)
	}

	fmt.Println("Blog was deleted", deleteRes)

	// List blogs

	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatalf("error while calling ListBlog RPC: %v", err)
	}
	i := 0
	for {
		i += 1
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetBlog())
	}
	fmt.Println(i)

}
