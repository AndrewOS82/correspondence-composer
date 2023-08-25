package rulesengine

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"correspondence-composer/models"
)

type gateway struct{}

// nolint
func New() *gateway {
	return &gateway{}
}

func (g *gateway) ExecuteRules(rules []*models.Rule) (*models.RulesAdminResponse, error) {
	// TODO: set up a proper config file at the root level of the application that loads .env
	g.loadEnvFile()

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
	apiURL := os.Getenv("RULES_ENGINE_EXECUTE_ENDPOINT")

	rulesRequest := g.buildRulesRequest(rules)
	requestBody, err := json.Marshal(rulesRequest)
	if err != nil {
		fmt.Printf("Error building rule execution request parameters: %v\n", err)
		return nil, err
	}

	client := &http.Client{}
	r, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
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
	apiURL := os.Getenv("RULES_ENGINE_AUTH_ENDPOINT")
	username := os.Getenv("RULES_ENGINE_USERNAME")
	password := os.Getenv("RULES_ENGINE_PASSWORD")
	clientCode := os.Getenv("RULES_ENGINE_AUTH_CLIENT_CODE")
	data := url.Values{
		"username":   {username},
		"password":   {password},
		"clientCode": {clientCode},
	}

	resp, err := http.PostForm(apiURL, data)
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
	client := os.Getenv("RULES_ENGINE_EXECUTE_CLIENT")
	rulesRequest := &models.RulesAdminRequest{
		Client: client,
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

// TODO: move this into a shared util and set up a proper config file
func (g *gateway) loadEnvFile() {
	path, _ := os.Getwd()
	for {
		envFile := path + "/.env"
		if _, err := os.Stat(envFile); err == nil {
			err := godotenv.Load(envFile)
			if err != nil {
				fmt.Println(err)
			}
		}
		if len(path) <= 1 {
			break
		}

		path = filepath.Dir(path)
	}
}
