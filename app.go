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
  // Reviewers      []string
}

func getClient(ctx context.Context, token string) *github.Client {
  ts := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: token},
  )
  tc := oauth2.NewClient(ctx, ts)

  return github.NewClient(tc)
}

func getRepoOptions() string {
  userInput := os.Args[2]
  return userInput
}

func getToken() string {
  token := os.Args[1]
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

func requestedReview(pr github.PullRequest, user string) bool {
  for _, reviewer := range pr.RequestedReviewers {
    if *reviewer.Login == user {
      return true
    }
  }
  return false
}

func getPullRequests(client *github.Client, ctx context.Context, user string, repo string) []PullRequest {
  fetchedPrs := fetchPullRequests(client, ctx, repo)

  prs := make([]PullRequest, 0)

  for _, pr := range fetchedPrs {
    if requestedReview(*pr, user) {
      title := *pr.Title
      author := *pr.User.Login
      prs = append(prs, PullRequest{title, author})
    }
  }

  return prs
}

func main() {
  token := getToken()
  ctx := context.Background()
  client := getClient(ctx, token)
  repo := getRepoOptions()

  user := getCurrentUser(ctx, client)
  pullRequests := getPullRequests(client, ctx, user, repo)

  fmt.Printf("PullRequests %v\n", pullRequests)

}
