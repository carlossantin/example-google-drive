package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	authSrv := authService{}
	client, err := authSrv.GetGoogleClient(ctx, "./credentials/credentials.json")
	if err != nil {
		panic(err)
	}

	// Create Google Drive Service
	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %+v", err)
	}

	// List files in Google Drive
	r, err := srv.Files.List().
		PageSize(10).
		Fields("nextPageToken, files(id, name)").
		Do()

	if err != nil {
		log.Fatalf("Unable to retrieve files: %+v", err)
	}
	fmt.Println("Files:")
	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Files {
			fmt.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}
}
