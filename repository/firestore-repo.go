package repository

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/kristianespina/golang-rest-api/entity"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type repo struct{}

// NewFirestoreRepository creates as new repo
func NewFirestoreRepository() PostRepository {
	return &repo{}
}

const (
	projectID      string = "go-rest-api"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("C:\\Users\\Skye\\Desktop\\Experiment\\golang-rest-api\\keys\\go-rest-api-989a0-firebase-adminsdk-a70bt-64520e94e9.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("C:\\Users\\Skye\\Desktop\\Experiment\\golang-rest-api\\keys\\go-rest-api-989a0-firebase-adminsdk-a70bt-64520e94e9.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	var posts []entity.Post

	iter := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iter.Next()

		// Stop when done
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}

		// Get Post
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}

		// Append Post into posts slice
		posts = append(posts, post)
	}

	return posts, nil
}
