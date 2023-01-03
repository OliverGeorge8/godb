package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Model struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
}

type TimestampsModel struct {
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func (TimestampsModel) BeforeUpdate(m interface{}) {

}

func (m Model) GetID() string {
	return m.ID.Hex()
}

func (m Model) setID(id interface{}) {
	if objectId, ok := id.(primitive.ObjectID); ok {
		m.ID = objectId
		return
	} else if strObjectId, ok := id.(string); ok && primitive.IsValidObjectID(strObjectId) {
		m.ID, _ = primitive.ObjectIDFromHex(strObjectId)
		return
	}

}

type IModel interface {
	GetID() string
	SetID(id interface{})
}
