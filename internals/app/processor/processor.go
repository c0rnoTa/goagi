package processor

import (
	"bufio"
	"fmt"
	"github.com/c0rnoTa/goagi/internals/app/models/repository/mysql"
	"github.com/zaf/agi"
	"log"
)

type Processor struct {
	storage *mysql.Storage
	agi     *agi.Session
}

func NewProcessor(storage *mysql.Storage, rw *bufio.ReadWriter) *Processor {
	myAgi := agi.New()
	err := myAgi.Init(rw)
	if err != nil {
		log.Fatalf("unable to create new processor: error Parsing AGI environment: %v\n", err)
	}
	return &Processor{storage: storage, agi: myAgi}
}

func (p *Processor) Verbose(msg interface{}) error {
	if msg == "" {
		log.Printf("Processor verbose method called without msg. Nothing to print.")
		return nil
	}

	if _, err := p.agi.Verbose(msg); err != nil {
		return fmt.Errorf("unable to process verbose msg: %w", err)
	}

	return nil
}
