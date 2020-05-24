package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	firebase "firebase.google.com/go"
)

type container struct {
	Status string    `json:"status"`
	Seen   time.Time `json:"seen"`
}

func main() {
	ctx := context.Background()

	var dryRun bool
	var projectID string
	var collectionName string
	var cluster string

	flag.BoolVar(&dryRun, "dry-run", false, "don't write to Firebase")
	flag.StringVar(&projectID, "project-id", "", "Firebase project ID")
	flag.StringVar(&collectionName, "name", "monitoring", "name of Firebase collection to write to")
	flag.StringVar(&cluster, "cluster", "", "name of the cluster/server you're monitoring")
	flag.Parse()

	if projectID == "" || cluster == "" {
		log.Println("one of project-id and cluster missing - falling back to dry-run")
		dryRun = true
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		log.Printf("failed to create docker client: %v", err)
		os.Exit(2)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		log.Printf("failed to list containers: %v", err)
		os.Exit(2)
	}

	now := time.Now()
	doc := make(map[string]container)
	for _, c := range containers {
		doc[c.Image] = container{Status: c.State, Seen: now}
		log.Printf("%s=%s %s\n", c.Image, c.Status, c.State)
	}

	if dryRun {
		if err := dumpResult(doc); err != nil {
			log.Print(err.Error())
			os.Exit(2)
		}
		os.Exit(0)
	}

	if err = save(ctx, projectID, collectionName, cluster, doc); err != nil {
		os.Exit(2)
	}

	log.Println("status updated successfully")
}

// save writes the list of container statuses to firebase
func save(ctx context.Context, projectID, collection, cluster string, data map[string]container) error {
	conf := firebase.Config{
		ProjectID: projectID,
	}
	app, err := firebase.NewApp(ctx, &conf)
	if err != nil {
		log.Printf("error initializing app: %v", err)
		return err
	}

	db, err := app.Firestore(ctx)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return err
	}

	_, err = db.Collection(collection).Doc(cluster).Set(ctx, data)
	if err != nil {
		log.Printf("failed to write data: %v", err)
		return err
	}

	return nil
}

func dumpResult(doc map[string]container) error {
	b, err := json.MarshalIndent(doc, "", " ")
	if err != nil {
		return err
	}
	log.Printf("result:\n%s", string(b))
	return nil
}
