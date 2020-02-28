package assignment2

import (
	"context"

	"cloud.google.com/go/firestore"
)

//Project struct
type Project struct {
	ID      int    `json:"id"`
	Name    string `json:"path_with_namespace"`
	Commits int
}

//Projectoutput struct
type Projectoutput struct {
	Name        string `json:"path_with_namespace"`
	CommitCount int    `json:"count"`
}

//Commit struct
type Commit struct {
	ID string `json:"id"`
}

//RepoOutput struct
type RepoOutput struct {
	Repos []Projectoutput `json:"Repos"`
	Auth  bool            `json:"Auth"`
}

//Status struct
type Status struct {
	Database int
	Gitlab   int
	Version  string
	Uptime   float64
}

//ProjectID struct
type ProjectID struct {
	ID int `json:"id"`
}

//Language struct
type Language struct {
	Name  string
	Count int
}

//LanguageName struct
type LanguageName struct {
	Name []string
}

//ReturnLanguage struct
type ReturnLanguage struct {
	Languages []string
	Auth      bool
}

//FirestoreDatabase struct
type FirestoreDatabase struct {
	CollectionName string
	Ctx            context.Context
	Client         *firestore.Client
}

//WebhookRegistration struct
type WebhookRegistration struct {
	ID        string `json:"ID"`
	URL       string `json:"url"`
	Event     string `json:"event"`
	Timestamp string `json:"Timestamp"`
}
