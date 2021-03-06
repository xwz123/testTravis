package main

import (
	"encoding/json"
	"fmt"
)

const (
	jsonNote = `{"id":"123","action":"add note","comment":"comment dd"}`
	jsonPr   = `{"id":"123","action":"open","pull_request":"add pr"}`
)

type noteHook struct {
	ID      string `json:"id"`
	Action  string `json:"action"`
	Comment string `json:"comment"`
}

type prHook struct {
	ID          string `json:"id"`
	Action      string `json:"action"`
	PullRequest string `json:"pull_request"`
}

func sayTravisHello() {
	fmt.Println("say hello by travis ci")
}

func parsePayload(payLoad []byte, v interface{}) error {
	err := json.Unmarshal(payLoad, v)
	if err != nil {
		return err
	}
	switch t := v.(type) {
	case *noteHook:
		fmt.Println(t.Comment)
	case *prHook:
		fmt.Println(t.PullRequest)
	default:
		fmt.Printf("unexpected type %T", v)
	}
	return nil
}
