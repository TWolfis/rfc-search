package search

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"regexp"

	"github.com/TWolfis/ietfrfc"
)

// linked list of RFC files
type RfcList struct {
	head *RfcNode
	tail *RfcNode
}

type RfcNode struct {
	value ietfrfc.RFC
	next  *RfcNode
	prev  *RfcNode
}

func readRfc(file string) (ietfrfc.RFC, error) {

	var rfc ietfrfc.RFC

	data, err := os.ReadFile(file)
	if err != nil {
		return ietfrfc.RFC{}, err
	}

	// Parse RFC
	if err := json.Unmarshal(data, &rfc); err != nil {
		return ietfrfc.RFC{}, err
	}

	return rfc, nil
}

// Create list with RFC's
func NewRfcList(srcDir string) *RfcList {
	var rl RfcList
	files, err := os.ReadDir(srcDir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		rfc, err := readRfc(path.Join(srcDir, file.Name()))
		if err != nil {
			panic(err)
		}

		rl.AddRfc(rfc)
	}

	return &rl
}

// Add RFC to list
func (rl *RfcList) AddRfc(rfc ietfrfc.RFC) {
	node := &RfcNode{value: rfc}

	if rl.head == nil {
		rl.head = node
		rl.tail = node
	} else {
		rl.tail.next = node
		node.prev = rl.tail
		rl.tail = node
	}
}

func (rl *RfcList) SearchRfc(query string) (ietfrfc.RFC, error) {
	regex := regexp.MustCompile(`(?i)` + query) // case insensitive search

	// Iterate through list
	current := rl.head
	for current != nil {

		if regex.MatchString(current.value.Title) {
			return current.value, nil
		}

		// Move to next node
		current = current.next
	}
	return ietfrfc.RFC{}, errors.New("query yielded no result")
}
