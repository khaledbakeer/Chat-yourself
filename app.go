package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

/*
a ver simple project, lets two instances of chat openapi to talk to eachother.

./app -key1=key1 -key2=key2 -start="was das Thema sein soll"
*/

func main() {
	key1 := flag.String("key1", "", "OpenAI API Key for the first instance")
	key2 := flag.String("key2", "", "OpenAI API Key for the second instance")
	start := flag.String("start", "Who are you?", "Start the conversation with")

	flag.Parse()

	var msg *string
	msg = start
	fmt.Println("\nThema ist: ")
	fmt.Println(*msg)

	_ = checkResponse(msg, *key1)
	res := fmt.Sprintf("%s %s", *msg, " and you?")
	msg = checkResponse(&res, *key2)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("A: ")
			msg = checkResponse(msg, *key1)
		} else {
			fmt.Println("B: ")
			msg = checkResponse(msg, *key2)
		}
		fmt.Println(*msg)

		time.Sleep(10 * time.Second)
		fmt.Println("-----------------------------------------------------------------------------------")
	}

}

func sendPrompt(msg *string, key string) *string {

	data := createPayload(msg)
	payloadBytes, _ := json.Marshal(data)
	reader := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", reader)
	if err != nil {
		panic(fmt.Sprintf("connection error: %s", err.Error()))
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", key))

	netTransport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 50 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 50 * time.Second,
	}
	netClient := &http.Client{
		Timeout:   time.Second * 100,
		Transport: netTransport,
	}

	resp, err := netClient.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var source = ApiResponse{}

	err = json.Unmarshal(b, &source)
	if err != nil {
		panic(err.Error())
	}

	if len(source.Choices) <= 0 {
		return sendPrompt(msg, key)
	}

	return &source.Choices[0].Text
}

func checkResponse(msg *string, key string) *string {
	if resp := sendPrompt(msg, key); resp == nil || *resp == "" || *resp == " " {
		fmt.Printf("%s", "still thinking!! ")
		time.Sleep(1 * time.Second)
		return checkResponse(msg, key)
	} else {
		return resp
	}
}

func createPayload(msg *string) *Payload {
	return &Payload{
		Model:            "text-davinci-003",
		Prompt:           msg,
		Temperature:      0.9,
		MaxTokens:        1500,
		TopP:             1,
		FrequencyPenalty: 0.9,
		PresencePenalty:  0.9,
		Stop:             []string{" Human: ", " AI: "},
	}
}
