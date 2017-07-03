package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	que "github.com/bgentry/que-go"
	"github.com/jackc/pgx"

	"github.com/zhiruchen/PlayQuego/config"
	"github.com/zhiruchen/PlayQuego/sql"
	"github.com/zhiruchen/PlayQuego/task"
)

var (
	qc      *que.Client
	pgxpool *pgx.ConnPool
)

func main() {
	var err error
	pgxpool, qc, err = sql.Setup(config.DbURL)
	if err != nil {
		log.Fatalf("setup db error: %v", err)
	}

	defer pgxpool.Close()

	wm := que.WorkMap{
		config.QueueName1: task.Add,
		config.QueueName2: task.Mul,
	}

	workers := que.NewWorkerPool(qc, wm, 2)

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	go workers.Start()

	// Wait for a signal
	sig := <-sigCh
	log.Printf("Signal received. Shutting down. %v\n", sig)

	workers.Shutdown()
}
