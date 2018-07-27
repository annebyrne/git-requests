package main

import (
  "github.com/google/go-github/github"
  "golang.org/x/oauth2"
  "context"
  "fmt"
  "os"
)

func getClient(ctx context.Context, token string) *github.Client {
  ts := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: token},
  )
  tc := oauth2.NewClient(ctx, ts)

  return github.NewClient(tc)
}

func getRepoOptions() string {
  userInput := os.Args[1]
  return userInput
}

func getToken() string {
  token := os.Args[2]
  return token
}

func getCurrentUser(ctx context.Context, client *github.Client) string {
  result, _, authError := client.Users.Get(ctx, "")
  user := *result.Login

  if authError!= nil {
    fmt.Printf("Problem in getting authenticated user information %v\n", authError)
    os.Exit(1)
  }
  return user
}

func main() {

  token := getToken()
  // get go-github client
  ctx := context.Background()
  client := getClient(ctx, token)

  repo := getRepoOptions()
  opts :=  &github.PullRequestListOptions{ Direction: "asc"}

  result, _, err := client.PullRequests.List(ctx, "deliveroo", repo, opts)
 
  //
  // result, _, err := client.Search.Issues(ctx, "is:open+is:pr+review-requested:hammerfunk", opts)
  // result, _, err := client.Search.Issues(ctx, "windows+label:bug+state:open", opts)
  // result, _, err := client.Search.Users(ctx, "tom+repos:%3E42+followers:%3E1000", nil)
  // result, _, err := client.Search.Users(ctx, "hammerfunk", nil)
  // https://api.github.com/search/issues/?q\=is:open+is:pr+review-requested:annebyrne
  // result, _, err := client.Organizations.List(ctx, "hammerfunk", nil)

  type PullRequest struct {
    Title, Author  string
  }

  user := getCurrentUser(ctx, client)

  title := *result[0].Title
  author := *result[0].User.Login
  pr := PullRequest{title, author}

  if err != nil {
    fmt.Printf("Problem in getting repository information %v\n", err)
    os.Exit(1)
  }

  

  fmt.Printf("RESULT result %v\n", response)
  fmt.Printf("RESULT result %v\n", pr)
  fmt.Printf("RESULT result %v\n", user)
}
