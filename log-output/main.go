package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type PingPongResp struct {
	Pongs int64 `json:"pongs"`
}

type GreetingResp struct {
	Greeting string `json:"greetings"`
}

var (
	port        = os.Getenv("PORT")
	pingPongURL = os.Getenv("PING_PONG_URL")
	message     = os.Getenv("MESSAGE")
	greeterURL  = os.Getenv("GREETER_URL")
)

var randomString = uuid.New().String()

func main() {
	if port == "" {
		log.Fatal().Msg("$PORT must be set")
	}

	if pingPongURL == "" {
		log.Fatal().Msg("$PING_PONG_URL must be set")
	}

	if message == "" {
		log.Fatal().Msg("$MESSAGE must be set")
	}

	if greeterURL == "" {
		log.Fatal().Msg("$GREETER_URL must be set")
	}

	fileContent, err := os.ReadFile("/usr/src/app/files/information.txt")
	var fileContentStr string
	if err != nil {
		fileContentStr = "file not found"
	} else {
		fileContentStr = string(fileContent)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":     "Welcome to the Log Output Service!",
			"status_code": http.StatusOK,
		})
	})

	router.GET("/log", func(c *gin.Context) {
		count, err := getPingCount()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		greeting, err := getGreeting()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		log.Info().Msg(greeting)

		timestamp := time.Now().Format(time.RFC3339)

		c.JSON(http.StatusOK, gin.H{
			"file content": fileContentStr,
			"env variable": fmt.Sprintf("MESSAGE=%s", message),
			timestamp:      randomString,
			"Ping / Pongs": count,
			"Greetings":    greeting,
		})
	})

	err = router.Run(":" + port)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server:")
	}
}

func getPingCount() (int64, error) {
	u, err := url.Parse(pingPongURL)
	if err != nil {
		return 0, fmt.Errorf("unable to parse PING_PONG_URL: %w", err)
	}
	u.Path = "/pings"

	resp, err := http.Get(u.String())
	if err != nil {
		return 0, fmt.Errorf("unable to get ping count: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("unable to read response body: %w", err)
	}

	var pingPongResp PingPongResp

	if err := json.Unmarshal(body, &pingPongResp); err != nil {
		return 0, fmt.Errorf("unable to unmarshal response: %w", err)
	}

	return pingPongResp.Pongs, nil
}

func getGreeting() (string, error) {
	u, err := url.Parse(greeterURL)
	if err != nil {
		return "", fmt.Errorf("unable to parse GREETER_URL: %w", err)
	}
	u.Path = "/greeter"

	resp, err := http.Get(u.String())
	if err != nil {
		return "", fmt.Errorf("unable to get greeting: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("unable to read response body: %w", err)
	}

	var greetingResp GreetingResp

	if err := json.Unmarshal(body, &greetingResp); err != nil {
		return "", fmt.Errorf("unable to unmarshal response: %w", err)
	}

	return greetingResp.Greeting, nil
}
