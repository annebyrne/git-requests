package main

import (
  "github.com/google/go-github/github"
  "golang.org/x/oauth2"
  "context"
  "fmt"
  "os"
)

func main() {
  ctx := context.Background()
  ts := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: "TOKEN"},
  )
  tc := oauth2.NewClient(ctx, ts)

  // get go-github client
  client := github.NewClient(tc)

  // opts :=  &github.SearchOptions{Sort: "created", Order: "asc"}
  // result, _, err := client.Search.Issues(ctx, "is:open+is:pr+review-requested:hammerfunk", opts)
  // result, _, err := client.Search.Issues(ctx, "windows+label:bug+state:open", opts)
  // result, _, err := client.Search.Users(ctx, "tom+repos:%3E42+followers:%3E1000", nil)
  // result, _, err := client.Search.Users(ctx, "hammerfunk", nil)
  result, _, err := client.Activity.ListNotifications(ctx, nil)
  // https://api.github.com/search/issues/?q\=is:open+is:pr+review-requested:annebyrne
  // result, _, err := client.Organizations.List(ctx, "hammerfunk", nil)

  if err != nil {
    fmt.Printf("Problem in getting repository information %v\n", err)
    os.Exit(1)
  }

  fmt.Printf("RESULT result %v\n", result)
}
