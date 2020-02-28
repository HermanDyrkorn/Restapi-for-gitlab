package assignment2

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

//Database collectionname
const databasecollection string = "webhooks"

//DB database
var DB = FirestoreDatabase{CollectionName: databasecollection}

//InitDatabase function initializes the database
func InitDatabase() error {
	DB.Ctx = context.Background()
	sa := option.WithCredentialsFile("./assignment2-3bd71-e15779190f0c.json")
	app, err := firebase.NewApp(DB.Ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	DB.Client, err = app.Firestore(DB.Ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

//Save function saves a webhook in the database
func Save(w *WebhookRegistration) error {
	ref := DB.Client.Collection(DB.CollectionName).NewDoc()
	w.ID = ref.ID
	_, err := ref.Set(DB.Ctx, w)
	if err != nil {
		fmt.Println("ERROR saving student to Firestore DB: ", err)
		return errors.Wrap(err, "Error in FirebaseDatabase.Save()")
	}

	return nil
}

//GetData from database function
func GetData(w http.ResponseWriter) error {
	iter := DB.Client.Collection(DB.CollectionName).Documents(DB.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		err2 := json.NewEncoder(w).Encode(doc.Data())
		if err2 != nil {
			return err2
		}
	}
	return nil
}

//ReturnAllWebhooks returns all the webhooks from the database
func ReturnAllWebhooks() ([]WebhookRegistration, error) {
	var webhooks []WebhookRegistration

	iter := DB.Client.Collection(DB.CollectionName).Documents(DB.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err == iterator.Done {
			return nil, err
		}
		webhook := WebhookRegistration{}
		e := doc.DataTo(&webhook)
		if e != nil {
			return nil, e
		}
		webhooks = append(webhooks, webhook)
	}
	return webhooks, nil
}

//FindID function find a webhook with id
func FindID(id string) ([]WebhookRegistration, error) {
	iter := DB.Client.Collection(DB.CollectionName).Where("ID", "==", id).Documents(DB.Ctx)
	var webhook = []WebhookRegistration{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Printf("Error when querying for name: %v\n%v\n", id, err)
		}

		hook := WebhookRegistration{}
		err = doc.DataTo(&hook)
		if err != nil {
			fmt.Println("Error when converting retrieved document to Student struct: ", err)
		}
		webhook = append(webhook, hook)
	}
	return webhook, nil

}

//DeleteWebhook deletes a webhook from the database
func DeleteWebhook(id string) error {
	_, err := DB.Client.Collection(DB.CollectionName).Doc(id).Delete(DB.Ctx)
	if err != nil {
		fmt.Printf("ERROR deleting webhook (%v) from Firestore DB: %v\n", id, err)
		return errors.Wrap(err, "Error in FirebaseDatabase.Delete()")
	}
	return nil

}
