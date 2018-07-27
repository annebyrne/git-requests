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

  orgs, _, err := client.Repositories.List(ctx, "hammerfunk", nil)

  if err != nil {
    fmt.Printf("Problem in getting repository information %v\n", err)
    os.Exit(1)
  }

  fmt.Printf("RESULT orgs %v\n", orgs)
}
