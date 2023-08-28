package rulesengine

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"correspondence-composer/models"
)

type gateway struct {
	config Config
}

type Config struct {
	Username        string
	Password        string
	AuthEndpoint    string
	AuthClientCode  string
	ExecuteEndpoint string
	ExecuteClient   string
}

// nolint
func New(config Config) *gateway {
	return &gateway{
		config: config,
	}
}

func (g *gateway) ExecuteRules(rules []*models.Rule) (*models.RulesAdminResponse, error) {
	if len(rules) < 1 {
		return nil, errors.New("no rules provided")
	}

	token, err := g.getAuthToken()
	if err != nil {
		return nil, err
	}

	resp, err := g.executeRules(token, rules)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *gateway) executeRules(token *models.Token, rules []*models.Rule) (*models.RulesAdminResponse, error) {
	rulesRequest := g.buildRulesRequest(rules)
	requestBody, err := json.Marshal(rulesRequest)
	if err != nil {
		fmt.Printf("Error building rule execution request parameters: %v\n", err)
		return nil, err
	}

	client := &http.Client{}
	r, _ := http.NewRequest("POST", g.config.ExecuteEndpoint, bytes.NewBuffer(requestBody))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("authorization", "Bearer "+token.Token)
	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("Error executing rule: %v\n", err)
		return &models.RulesAdminResponse{}, err
	}

	body, err := g.responseBody(resp)
	if err != nil {
		return nil, err
	}

	var rulesResponse models.RulesAdminResponse
	err = json.Unmarshal([]byte(string(body)), &rulesResponse)
	if err != nil {
		fmt.Printf("Error decoding rule execution response: %v\n", err)
		return &rulesResponse, err
	}

	return &rulesResponse, nil
}

func (g *gateway) getAuthToken() (*models.Token, error) {
	data := url.Values{
		"username":   {g.config.Username},
		"password":   {g.config.Password},
		"clientCode": {g.config.AuthClientCode},
	}

	resp, err := http.PostForm(g.config.AuthEndpoint, data)
	if err != nil {
		fmt.Printf("Error fetching token: %v\n", err)
		return nil, err
	}

	body, err := g.responseBody(resp)
	if err != nil {
		return nil, err
	}

	var token models.Token
	err = json.Unmarshal([]byte(string(body)), &token)
	if err != nil {
		fmt.Printf("Error decoding token response: %v\n", err)
		return nil, err
	}

	return &token, nil
}

// This function builds the request parameters in the format that the rules engine expects.
// Example:
//
//	{
//		"client": "ZINNIA",
//		"source": "Camunda",
//		"rules": [
//			{
//				"ruleName": "CDS_TopicSelector",
//				"version": 1,
//				"input": {
//					"type": "corro.request.received"
//				}
//			}
//		]
//	}
func (g *gateway) buildRulesRequest(rules []*models.Rule) *models.RulesAdminRequest {
	rulesRequest := &models.RulesAdminRequest{
		Client: g.config.ExecuteClient,
		Source: "Camunda",
		Rules:  rules,
	}

	return rulesRequest
}

func (g *gateway) responseBody(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return nil, err
	}

	defer resp.Body.Close()

	return body, nil
}
