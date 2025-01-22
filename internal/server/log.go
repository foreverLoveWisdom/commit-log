package server

// "What problem does this file solve in the system?"

//  "What role does this file play in the system’s architecture?"
//
//     Is it handling requests?
//     Is it managing concurrency?
//     Is it defining data models?
//
// 2️⃣ "What are the key inputs and outputs of this file?"
//
//     What data does it receive?
//     What does it return or modify?
//
// 3️⃣ "How does this file interact with other parts of the system?"
//
//     Does it talk to a database?
//     Does it call external APIs?
//     Does it use goroutines or channels?

import (
	"fmt"
	"sync"
)

type Log struct {
	mu
	sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}

func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}

type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")
