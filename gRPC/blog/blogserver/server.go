package main

import (
	"context"
	"fmt"
	"gRPC/blog/blogpb"
	"log"
	"net"
	"os"
	"os/signal"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type exampleserver struct {
}

var collection *mongo.Collection
var ctx = context.TODO()

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func (*exampleserver) UpdateBlog(ctx context.Context, in *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	fmt.Println("update blog requst")
	blog := in.GetBlog()

	oid, err := primitive.ObjectIDFromHex(blog.GetId())

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument recieved cannot parse ID")
	}

	data := &blogItem{}
	filter := bson.M{"_id": oid}
	filteredData := collection.FindOne(context.Background(), filter)

	if err := filteredData.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find the specified  blog with the requested ID %v", err))
	}
	data.AuthorID = blog.GetAuthorId()
	data.Content = blog.GetContent()
	data.Title = blog.GetTitle()

	_, updateErr := collection.ReplaceOne(context.Background(), filter, data)

	if updateErr != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot update object in mongodb %v", updateErr))
	}
	return &blogpb.UpdateBlogResponse{
		Blog: DatatoBlogpb(data),
	}, nil

}

func DatatoBlogpb(data *blogItem) *blogpb.Blog {
	return &blogpb.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorID,
		Content:  data.Content,
		Title:    data.Title,
	}
}

func (*exampleserver) ReadBlog(ctx context.Context, in *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	blogId := in.GetBlogId()

	fmt.Println("Recieved a request to read blog", blogId)

	oid, err := primitive.ObjectIDFromHex("61d086534a1fb8f4e45612e3")
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument recieved cannot parse ID")
	}

	// create an empty struct
	data := &blogItem{}
	// filter := bson.NewDocument(bson.Ec.ObjectID("_id", oid))

	// fmt.Println("Object Id ", oid)

	filteredData := collection.FindOne(context.Background(), bson.M{"ID": oid})

	if err := filteredData.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find the specified  blog with the requested ID %v", err))
	}

	return &blogpb.ReadBlogResponse{
		Blog: &blogpb.Blog{
			Id:       data.ID.Hex(),
			AuthorId: data.AuthorID,
			Content:  data.Content,
			Title:    data.Title,
		},
	}, nil

}

func (*exampleserver) CreateBlog(ctx context.Context, in *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {

	blog := in.GetBlog()

	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}

	fmt.Println("Recived req to create a blog for Author with id ", data.AuthorID)

	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error while inserting data %v", err))
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("cannot convert data %v", err))
	}

	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       oid.Hex(),
			AuthorId: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		},
	}, nil

}

func (*exampleserver) DeleteBlog(ctx context.Context, in *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	fmt.Println("delete blog request")
	blogId := in.GetBlogId()

	fmt.Println("Recieved a request to read blog", blogId)

	oid, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument recieved cannot parse ID")
	}

	filter := bson.M{"_id": oid}
	res, err2 := collection.DeleteOne(context.Background(), filter)
	if err2 != nil {
		return nil, status.Errorf(codes.Internal, "Couldnt delete the post")
	}
	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Cannot find the document under the blog")
	}

	return &blogpb.DeleteBlogResponse{BlogId: in.GetBlogId()}, nil
}

func (*exampleserver) ListBlog(in *blogpb.ListBlogRequest, stream blogpb.BlogService_ListBlogServer) error {
	fmt.Println("ListBlog request")
	curr, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown error %v", err))
	}

	defer curr.Close(context.Background())

	for curr.Next(context.Background()) {
		data := &blogItem{}
		err := curr.Decode(data)
		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("unknown error while decoding %v", err))
		}
		stream.Send(&blogpb.ListBlogResponse{Blog: DatatoBlogpb(data)})

	}
	if err := curr.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("unknown error : %v", err),
		)
	}
	return nil

}

func main() {
	// if we crash we get the file name and the line number

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("hello World")

	// connect to mongodb
	fmt.Println("connecting to mongodb")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("mydb").Collection("blog")

	fmt.Println("connected to mongodb")

	lis, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	tls := false

	opts := []grpc.ServerOption{}

	if tls {
		certFile := "D:/Udemy/go/gRPC/ssl/server.crt"
		keyFile := "D:/Udemy/go/gRPC/ssl/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)

		if sslErr != nil {
			log.Fatalf("Failed loading certificates : %v", sslErr)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	blogpb.RegisterBlogServiceServer(s, &exampleserver{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		fmt.Println("Starting the server..")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to server : %v", err)
		}
	}()

	// Wait for control c to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Blog the program until a signal was recieved
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the port 5000")
	lis.Close()
	fmt.Println("Closing the db connection to mongodb")
	client.Disconnect(context.TODO())
	fmt.Println("Ending the program")

}
