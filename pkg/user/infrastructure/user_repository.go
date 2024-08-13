package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/versionary-server/internal/database"
	apperrors "github.com/namhq1989/versionary-server/internal/utils/error"
	"github.com/namhq1989/versionary-server/pkg/user/domain"
	"github.com/namhq1989/versionary-server/pkg/user/infrastructure/dbmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	db             *database.Database
	collectionName string
}

func NewUserRepository(db *database.Database) UserRepository {
	r := UserRepository{
		db:             db,
		collectionName: database.Collections.User,
	}
	r.ensureIndexes()
	return r
}

func (r UserRepository) ensureIndexes() {
	var (
		ctx           = context.Background()
		opts          = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		isEmailUnique = true
		indexes       = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "email", Value: 1}},
				Options: &options.IndexOptions{
					Unique: &isEmailUnique,
				},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r UserRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r UserRepository) FindUserByID(ctx *appcontext.AppContext, userID string) (*domain.User, error) {
	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	var doc dbmodel.User
	if err = r.collection().FindOne(ctx.Context(), bson.M{
		"_id": uid,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	result := doc.ToDomain()
	return &result, nil
}

func (r UserRepository) UpdateUser(ctx *appcontext.AppContext, user domain.User) error {
	doc, err := dbmodel.User{}.FromDomain(user)
	if err != nil {
		return err
	}

	_, err = r.collection().UpdateByID(ctx.Context(), doc.ID, bson.M{"$set": doc})
	return err
}
