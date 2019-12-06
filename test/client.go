package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "random_passenger_driver/internal/proto"

	"google.golang.org/grpc"
)

type Config struct {
	Host            string
	Port            int
	SecListenStream int
}

func main() {
	cfg := Config{
		// Host:            "localhost",
		Host:            "80.93.182.105",
		Port:            50005,
		SecListenStream: 10,
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error when try grpc.Dial, err: %s", err)
	}

	client := pb.NewPassengerDriverClient(conn)

	streamDriver, err := client.Driver(context.Background())
	if err != nil {
		log.Fatalf("Error when try client.Driver, err: %s", err)
	}

	streamOrder, err := client.Order(context.Background())
	if err != nil {
		log.Fatalf("Error when try client.Order, err: %s", err)
	}

	ctxDriver := streamDriver.Context()
	ctxOrder := streamOrder.Context()

	done := make(chan bool)

	go func() {
		for {
			resp, err := streamDriver.Recv()
			if err == io.EOF {
				close(done)
				return
			}

			if err != nil {
				log.Fatalf("Error when try stream.Recv, err: %s", err)
			}

			log.Printf("Receive from server (driver): %v", resp)
		}
	}()

	go func() {
		for {
			resp, err := streamOrder.Recv()
			if err == io.EOF {
				close(done)
				return
			}

			if err != nil {
				log.Fatalf("Error when try stream.Recv, err: %s", err)
			}

			log.Printf("Receive from server (order): %v", resp)
		}
	}()

	timer := time.NewTimer(time.Duration(cfg.SecListenStream) * time.Second)

	go func() {
		<-timer.C

		<-ctxDriver.Done()

		if err := ctxDriver.Err(); err != nil {
			log.Printf("Error when try ctxDriver.Err, err: %s", err)
		}

		<-ctxOrder.Done()

		if err := ctxOrder.Err(); err != nil {
			log.Printf("Error when try ctxDriver.Err, err: %s", err)
		}

		close(done)
	}()

	<-done
}
