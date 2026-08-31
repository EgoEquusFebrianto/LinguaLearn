package models

import "go.mongodb.org/mongo-driver/v2/bson"

type BankWord struct {
	ID       bson.ObjectID `bson:"_id,omitempty" json:"id"`
	WordUUID string        `bson:"word_uuid" json:"word_uuid"`
	Word     string        `bson:"word" json:"word"`
	Senses   []Sense       `bson:"senses" json:"senses"`
}

type Sense struct {
	Translations []string `bson:"translations" json:"translations"`
	Synonyms     []string `bson:"synonyms" json:"synonyms"`
	Type         string   `bson:"type" json:"type"`
	Description  string   `bson:"description" json:"description"`
}