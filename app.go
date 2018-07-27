package main

import (
  "github.com/google/go-github/github"
  "golang.org/x/oauth2"
  "context"
  "fmt"
  "os"
)

type PullRequest struct {
  Title, Author  string
}

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

func fetchPullRequests(client *github.Client, ctx context.Context, repo string) []*github.PullRequest {
  opts :=  &github.PullRequestListOptions{ Direction: "asc"}
  prs, _, err := client.PullRequests.List(ctx, "deliveroo", repo, opts)

  if err != nil {
    fmt.Printf("Problem in getting repository information %v\n", err)
    os.Exit(1)
  }

  return prs
}

func getPullRequests(client *github.Client, ctx context.Context, repo string) []PullRequest {
  fetchedPrs := fetchPullRequests(client, ctx, repo)

  prs := make([]PullRequest, 0)

  for _, pr := range fetchedPrs {
    title := *pr.Title
    author := *pr.User.Login
    prs = append(prs, PullRequest{title, author})
  }

  return prs
}

func main() {
  token := getToken()
  ctx := context.Background()
  client := getClient(ctx, token)
  repo := getRepoOptions()

  //
  // result, _, err := client.Search.Issues(ctx, "is:open+is:pr+review-requested:hammerfunk", opts)
  // result, _, err := client.Search.Issues(ctx, "windows+label:bug+state:open", opts)
  // result, _, err := client.Search.Users(ctx, "tom+repos:%3E42+followers:%3E1000", nil)
  // result, _, err := client.Search.Users(ctx, "hammerfunk", nil)
  // https://api.github.com/search/issues/?q\=is:open+is:pr+review-requested:annebyrne
  // result, _, err := client.Organizations.List(ctx, "hammerfunk", nil)

  user := getCurrentUser(ctx, client)

  pullRequests := getPullRequests(client, ctx, repo)

  // fmt.Printf("RESULT result %v\n", response)
  fmt.Printf("PullRequests %v\n", pullRequests)
  fmt.Printf("User %v\n", user)
}
