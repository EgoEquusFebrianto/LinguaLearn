package seed

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DictionarySeed struct {
	ID          int      `json:"id"`
	Word        string   `json:"word"`
	Translate   []string `json:"translate"`
	Synonym     []string `json:"synonym"`
	Explanation string   `json:"explanation"`
	Type        string   `json:"type"`
}

type BankWord struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	WordUUID string        `bson:"word_uuid"`
	Word     string        `bson:"word"`
	Senses   []Sense       `bson:"senses"`
}

type Sense struct {
	Translations []string `bson:"translations"`
	Synonyms     []string `bson:"synonyms"`
	Type         string   `bson:"type"`
	Description  string   `bson:"description"`
}

var linguaLearnNamespace = uuid.MustParse(
	"6ba7b810-9dad-11d1-80b4-00c04fd430c8",
)

func generateWordUUID(word string) string {
	return uuid.NewSHA1(
		linguaLearnNamespace,
		[]byte(word),
	).String()
}

func loadDictionary(path string ) ([]DictionarySeed, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read dictionary file: %w", err)
	}

	var records []DictionarySeed

	if err := json.Unmarshal(data, &records); err != nil {
		return nil, fmt.Errorf("Failed to decode dictionary JSON: %w", err)
	}

	return records, nil
}

func transformDictionaryGrouped(records []DictionarySeed) map[string]BankWord {
	wordMap := make(map[string]BankWord)

	for _, record := range records {
		sense := Sense{
			Translations: record.Translate,
			Synonyms:     record.Synonym,
			Type:         record.Type,
			Description:  record.Explanation,
		}

		if existing, exists := wordMap[record.Word]; exists {
			// Word sudah ada, tambahkan sense baru ke array
			existing.Senses = append(existing.Senses, sense)
			wordMap[record.Word] = existing
		} else {
			// Word baru, buat dokumen baru
			wordMap[record.Word] = BankWord{
				WordUUID: generateWordUUID(record.Word),
				Word:     record.Word,
				Senses:   []Sense{sense},
			}
		}
	}

	return wordMap
}

func SeedDictionary(
	ctx context.Context,
	collection *mongo.Collection,
	records []DictionarySeed,
) error {
	wordMap := transformDictionaryGrouped(records)

	// Log statistik
	fmt.Printf("Total records: %d\n", len(records))
	fmt.Printf("Unique words: %d\n", len(wordMap))
	
	// Hitung total senses
	totalSenses := 0
	for _, word := range wordMap {
		totalSenses += len(word.Senses)
	}
	fmt.Printf("Total senses: %d\n", totalSenses)

	models := make([]mongo.WriteModel, 0, len(wordMap))

	for _, document := range wordMap {
		filter := bson.D{
			{Key: "word_uuid", Value: document.WordUUID},
		}

		// Gunakan $set untuk overwrite seluruh dokumen
		// Atau gunakan $push untuk menambah sense
		update := bson.D{
			{Key: "$set", Value: document},
		}

		models = append(
			models,
			mongo.NewUpdateOneModel().
				SetFilter(filter).
				SetUpdate(update).
				SetUpsert(true),
		)
	}

	if len(models) == 0 {
		return nil
	}

	_, err := collection.BulkWrite(ctx, models)
	if err != nil {
		return fmt.Errorf("Failed to seed bank_words: %w", err)
	}

	return nil
}

func RunDictionaryseed(
	client *mongo.Client,
	databaseName string,
	filePath string,
) error {
	records, err := loadDictionary(filePath)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		60 * time.Second,
	)
	defer cancel()

	collection := client.Database(databaseName).Collection("bank_words")

	if err := SeedDictionary(ctx, collection, records); err != nil {
		return err
	}

	return nil
}