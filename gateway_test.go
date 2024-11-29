package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/cucumber/godog"
)

var (
	apiEndpoint  string
	servicePath  string
	response     *http.Response
	err          error
	cookies      []*http.Cookie
	responseBody map[string]interface{}
)

func myAPIEndpointIs(endpoint string) error {
	apiEndpoint = endpoint
	return nil
}

func iAccessedPath(path string) error {
	servicePath = path
	return nil
}

func iSendAGETRequest() error {
	url := apiEndpoint + servicePath
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	if len(cookies) > 0 {
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}
	}

	client := &http.Client{}
	response, err = client.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(&responseBody)
}

func theResponseStatusCodeShouldBe(expectedStatus int) error {
	if response.StatusCode != expectedStatus {
		return fmt.Errorf("expected status code %d, got %d", expectedStatus, response.StatusCode)
	}
	return nil
}

func theResponseShouldContain(expectedKey string, expectedResult string) error {
	result, exists := responseBody[expectedKey]
	if !exists && result != expectedResult {
		return fmt.Errorf("expected key %s not found in response body", expectedKey)
	}
	return nil
}

func iSetCookiesSessionId() error {
	cookies = response.Cookies()
	return nil
}

func TestMain(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Given(`^my API endpoint is "([^"]*)"`, myAPIEndpointIs)
	sc.When(`^I accessed path "([^"]*)"`, iAccessedPath)
	sc.When(`^I send a GET request$`, iSendAGETRequest)
	sc.Then(`^the response status code should be (\d+)$`, theResponseStatusCodeShouldBe)
	sc.Then(`^the "([^"]*)" should contain "([^"]*)"$`, theResponseShouldContain)
	sc.Step(`^I set Cookies SessionId from response header$`, iSetCookiesSessionId)
}
