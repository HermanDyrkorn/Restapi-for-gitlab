package assignment2

import (
	"strconv"
	"testing"
)

func TestFirestoreDatabase(t *testing.T) {

	InitDatabase()
	//testing on a test collection database
	DB.CollectionName = "testDB"

	webhook := WebhookRegistration{ID: "1", URL: "URL", Event: "Event", Timestamp: "time"}
	err := Save(&webhook)
	if err != nil {
		t.Error(err)
	}
	res, err := FindID(webhook.ID)
	if err != nil {
		t.Error(err)
	}
	// test if res contains Alice
	if res[0].ID != webhook.ID {
		t.Errorf("Webhook.ID does not match!")
		return
	}
	if res[0].URL != webhook.URL {
		t.Errorf("Webhook.URL does not match!")
		return
	}
	if res[0].Event != webhook.Event {
		t.Errorf("Webhook.Event does not match!")
		return
	}
	if res[0].Timestamp != webhook.Timestamp {
		t.Errorf("Webhook.Timestamp does not match!")
		return
	}

	// we should clean-up
	err = DeleteWebhook(res[0].ID)
	if err != nil {
		t.Errorf("ERROR: removal of Webhook\n%v\n", err)
	}
	resAfterRemove, err := FindID(res[0].ID)
	// THERE MUST BE AN ERROR
	if len(resAfterRemove) != 0 {
		t.Error("FindID has not failed for deleted document!")
		return
	}
	//puts testdata in the database
	var webhooks [6]WebhookRegistration
	for i := 0; i <= 5; i++ {
		hook := WebhookRegistration{ID: strconv.Itoa(i), URL: "URL" + strconv.Itoa(i),
			Event: "Event" + strconv.Itoa(i), Timestamp: "time" + strconv.Itoa(i)}
		err := Save(&hook)
		if err != nil {
			t.Error("Save has failed to add document!")
		}
		webhooks[i] = hook
	}
	//gets all the data from the database
	everyWebhook, err := ReturnAllWebhooks()
	if err != nil {
		t.Error("Cant find any webhook")
	}
	//if the length of testdata is not the same as the database returned
	if len(everyWebhook) != len(webhooks) {
		t.Error("Did not return every registerd webhook")
	}

	//deleting everything from the database
	for j := range everyWebhook {
		err = DeleteWebhook(everyWebhook[j].ID)
		if err != nil {
			t.Errorf("ERROR: removal of Webhook\n%v\n", err)
		}
	}
}
