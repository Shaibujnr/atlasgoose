package main

import (
	"fmt"
	"github.com/Shaibujnr/atlasgoose/models"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createUser(db *gorm.DB, firstName, lastName, email string) (*models.User, error) {
	user := &models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
	result := db.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func addBlogPost(db *gorm.DB, userID uint, title, content string) (*models.BlogPost, error) {
	blogPost := &models.BlogPost{
		UserID:  userID,
		Title:   title,
		Content: content,
	}
	result := db.Save(blogPost)
	if result.Error != nil {
		return nil, result.Error
	}
	return blogPost, nil
}

func main() {
	dsn := "host=db user=tdb password=tdb dbname=tdb port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Transaction(func(tx *gorm.DB) error {
		for u := 0; u < 5; u++ {
			firstName := gofakeit.FirstName()
			lastName := gofakeit.LastName()
			email := gofakeit.Email()
			user, err := createUser(tx, firstName, lastName, email)
			if err != nil {
				panic(err)
			}
			for i := 0; i < 10; i++ {
				title := fmt.Sprintf("Title%d-%s", i, gofakeit.SentenceSimple())
				content := gofakeit.Sentence(500)
				_, err := addBlogPost(tx, user.ID, title, content)
				if err != nil {
					panic(err)
				}
			}
		}
		return nil
	})

}
