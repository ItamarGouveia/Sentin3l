package main

import (
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

func getAgentID() string {
	const file = "agent_id.txt"
	if _, err := os.Stat(file); err == nil {
		id, _ := ioutil.ReadFile(file)
		return string(id)
	}

	id := uuid.New().String()
	_ = ioutil.WriteFile(file, []byte(id), 0644)
	return id
}
