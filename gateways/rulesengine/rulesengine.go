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
	config    Config
	authToken *models.Token
}

type Config struct {
	Username        string
	Password        string
	AuthEndpoint    string
	ClientCode      string
	ExecuteEndpoint string
}

// nolint
func New(config Config) *gateway {
	token, _ := getAuthToken(config)

	return &gateway{
		authToken: token,
		config:    config,
	}
}

func (g *gateway) ExecuteRules(rules []*models.Rule) (*models.RulesEngineResponse, error) {
	if len(rules) < 1 {
		return nil, errors.New("no rules provided")
	}

	rulesRequest := g.buildRulesRequest(rules)
	resp, err := g.executeRulesRequest(rulesRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		g.refetchToken()
		resp, err = g.executeRulesRequest(rulesRequest)
		if err != nil {
			return nil, err
		}
	}

	body, err := g.responseBody(resp)
	if err != nil {
		return nil, err
	}

	var rulesResponse models.RulesEngineResponse
	err = json.Unmarshal([]byte(string(body)), &rulesResponse)
	if err != nil {
		fmt.Printf("Error decoding rule execution response: %v\n", err)
		return nil, err
	}

	return &rulesResponse, nil
}

func getAuthToken(config Config) (*models.Token, error) {
	data := url.Values{
		"username":   {config.Username},
		"password":   {config.Password},
		"clientCode": {config.ClientCode},
	}

	resp, err := http.PostForm(config.AuthEndpoint, data)
	if err != nil {
		fmt.Printf("Error fetching token: %v\n", err)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading token response: %v\n", err)
		return nil, err
	}

	defer resp.Body.Close()

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
func (g *gateway) buildRulesRequest(rules []*models.Rule) *models.RulesEngineRequest {
	rulesRequest := &models.RulesEngineRequest{
		Client: g.config.ClientCode,
		Source: "Camunda",
		Rules:  rules,
	}

	return rulesRequest
}

func (g *gateway) executeRulesRequest(rulesRequest *models.RulesEngineRequest) (*http.Response, error) {
	token := g.authToken.Token
	requestBody, err := json.Marshal(rulesRequest)
	if err != nil {
		fmt.Printf("Error building rule execution request parameters: %v\n", err)
		return nil, err
	}

	client := &http.Client{}
	r, _ := http.NewRequest("POST", g.config.ExecuteEndpoint, bytes.NewBuffer(requestBody))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("authorization", "Bearer "+token)
	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("Error executing rule: %v\n", err)
		return nil, err
	}

	return resp, nil
}

func (g *gateway) refetchToken() {
	token, _ := getAuthToken(g.config)
	g.authToken = token
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
