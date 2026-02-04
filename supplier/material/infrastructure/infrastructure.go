package infrastructure

import (
	"context"
	"domain/supplier/material/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	DB *mongo.Database
}

func NewMongoClient(uri, dbName string) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// 確認真的連得上
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &Client{
		DB: client.Database(dbName),
	}, nil
}

type MaterialRepository struct {
	collection string
	db         *mongo.Database
}

func (r *MaterialRepository) On(s string, ctx context.Context, material *domain.Material) {
	panic("unimplemented")
}

func NewMaterialRepository(db *mongo.Database) *MaterialRepository {
	return &MaterialRepository{
		db:         db,
		collection: "materials",
	}
}

type MaterialWiehId struct {
	domain.Material
	ID string `bson:"_id,omitempty"`
}

type CreateMaterialInfo struct {
	Id           primitive.ObjectID
	MaterialNo   string
	Name         string
	Attributes   []*domain.ProductAttrs
	Unit         string
	IsCustomized *bool
	IsApproved   *bool
	IsInquiry    *bool
	Status       string
	Category     []string
	Tags         []string
	Companies    []string
}

func (r *MaterialRepository) Create(ctx context.Context, m *CreateMaterialInfo) (*CreateMaterialInfo, error) {
	result, err := r.db.Collection(r.collection).InsertOne(ctx, m)

	if err != nil {
		return nil, err
	}

	m.Id = result.InsertedID.(primitive.ObjectID)
	return m, nil
}

func (r *MaterialRepository) Update(ctx context.Context, m *CreateMaterialInfo) (*CreateMaterialInfo, error) {
	update := bson.M{
		"$set": bson.M{
			"name":         m.Name,
			"attributes":   m.Attributes,
			"isCustomized": m.IsCustomized,
			"isApproved":   m.IsApproved,
			"isInquiry":    m.IsInquiry,
			"status":       m.Status,
			"category":     m.Category,
			"tags":         m.Tags,
			"companies":    m.Companies,
			"updatedAt":    time.Now(),
		},
	}

	_, err := r.db.Collection(r.collection).UpdateByID(ctx, m.Id, update)

	if err != nil {
		return nil, err
	}

	return m, nil
}
