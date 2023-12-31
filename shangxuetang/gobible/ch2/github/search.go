package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)
func SearchIssues(terms []string) (*IssuesSearchResult, error){
	fmt.Println("======"+strings.Join(terms, " ")+"======")
    q := url.QueryEscape(strings.Join(terms, " "))
	fmt.Println(IssuesURL+"?q="+q)
    resp, err := http.Get(IssuesURL + "?q=" + q)
    if err != nil {
        return nil, err
    }

    // We must close resp.Body on all execution paths.
    // (Chapter 5 presents 'defer', which makes this simpler.)
    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("search query failed: %s", resp.Status)
    }

    var result IssuesSearchResult
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        resp.Body.Close()
        return nil, err
    }
    resp.Body.Close()
    return &result, nil
}