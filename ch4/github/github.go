// O pacote github disponibiliza uma API Go para o sistema de
// acompanhamento de problemas do GitHub
// Veja https://developer.github.com/v3/search/#search-issues
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // em formato markdown
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues faz uma consulta ao sistema de acompanhamento de problemas
// do GitHub
func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// Devemos fechar resp.Body em todos os paths de execução
	// (O capítulo 5 apresenta 'defer', que simplificará isto)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
