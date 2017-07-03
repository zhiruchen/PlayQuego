package main

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/bgentry/que-go"
	"github.com/jackc/pgx"

	"github.com/zhiruchen/PlayQuego/config"
	"github.com/zhiruchen/PlayQuego/sql"
)

var (
	qc      *que.Client
	pgxpool *pgx.ConnPool
)

func queueAdd() error {
	data := struct {
		A string `json:"a"`
		B string `json:"b"`
	}{
		A: "Hello",
		B: "Queue",
	}
	enc, err := json.Marshal(&data)
	if err != nil {
		return errors.New("Marshalling the IndexRequest")
	}

	j := que.Job{
		Type:  config.QueueName1,
		Args:  enc,
		RunAt: time.Now().Add(time.Duration(30 * time.Second)),
	}

	return qc.Enqueue(&j)
}

func queueMul() error {
	data := struct {
		Str   string `json:"str"`
		Times int32  `json:"times"`
	}{
		Str:   "Task",
		Times: 10,
	}

	enc, err := json.Marshal(&data)
	if err != nil {
		return errors.New("Marshalling the IndexRequest")
	}

	j := que.Job{
		Type:  config.QueueName2,
		Args:  enc,
		RunAt: time.Now().Add(time.Duration(60 * time.Second)),
	}

	return qc.Enqueue(&j)
}

func main() {
	var err error
	pgxpool, qc, err = sql.Setup(config.DbURL)
	if err != nil {
		log.Fatalf("connect to postgresql error: %v", err)
	}

	defer pgxpool.Close()

	queueAdd()
	queueMul()
	log.Println("job enqueue ...")
}
