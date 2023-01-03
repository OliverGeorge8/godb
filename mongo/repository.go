package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository[T any] struct {
	collection *mongo.Collection
}

func (r *Repository[T]) InsertOneWithCtx(ctx context.Context, documentModel interface{}) (*T, error) {
	beforeInsertHook, ok := (documentModel).(BeforeInsert)

	if ok {
		beforeInsertHook.BeforeInsert()
	}

	result, err := r.collection.InsertOne(ctx, documentModel)
	if err != nil {
		return nil, err
	}

	if doc, ok := documentModel.(IModel); ok {
		doc.SetID(result.InsertedID)
	}

	return documentModel.(*T), nil
}
func (r *Repository[T]) InsertOne(documentModel *T) (*T, error) {
	return r.InsertOneWithCtx(context.Background(), documentModel)
}

func (r *Repository[T]) UpdateOneWithCtx(ctx context.Context, documentModel interface{}) (*T, error) {

}

func (r *Repository[T]) UpdateOne(documentModel *T) (*T, error) {
	return r.UpdateOneWithCtx(context.Background(), documentModel)
}
