Feature: apiTest
  In hitting consumerloan, cashloan, risk service
  As a hungry gopher
  I need to be able to receive 200 statuscode

    Scenario: Hit Consumerloan Service
        Given my API endpoint is "http://localhost:8080/gateway"
        When I accessed path "/consumerloan"
        When I send a GET request
        Then the response status code should be 200
        Then the "Result" should contain "OK"

    Scenario: Hit Cashloan Service
        Given my API endpoint is "http://localhost:8080/gateway"
        When I accessed path "/cashloan"
        When I send a GET request
        Then the response status code should be 200
        Then the "Result" should contain "OK"

    Scenario: Hit Risk Service
        Given my API endpoint is "http://localhost:8080/gateway"
        When I accessed path "/risk"
        When I send a GET request
        Then the response status code should be 200
        Then the "Result" should contain "OK"

    Scenario: Hit Undefined Path
        Given my API endpoint is "http://localhost:8080/gateway"
        When I accessed path "/risk/abc"
        When I send a GET request
        Then the response status code should be 500

    Scenario: Hit /search on Consumerloan Service
        Given my API endpoint is "http://localhost:8080/gateway"
        When I accessed path "/consumerloan"
        When I send a GET request
        And  I set Cookies SessionId from response header
        When I accessed path "/consumerloan/search"
        When I send a GET request
        Then the response status code should be 200


