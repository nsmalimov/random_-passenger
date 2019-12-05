package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"random_passenger_driver/configs"
	"random_passenger_driver/internal/coordinate_gen"
	"random_passenger_driver/internal/drivers_processor"
	"time"

	"random_passenger_driver/internal/order_gen"

	pb "random_passenger_driver/internal/proto"

	"google.golang.org/grpc"
)

type server struct{}

func (s server) Max(srv pb.Math_MaxServer) error {

	log.Println("start new server")
	var max int32
	ctx := srv.Context()

	for {

		// exit if context is done
		// or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from stream
		req, err := srv.Recv()
		if err == io.EOF {
			// return will close stream from server side
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}

		// continue if number reveived from stream
		// less than max
		if req.Num <= max {
			continue
		}

		// update max and send it to stream
		max = req.Num
		resp := pb.Response{Result: max}
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
		log.Printf("send new max=%d", max)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	cfg := configs.Config{
		CentralLatitude:  55.752818,
		CentralLongitude: 37.621753,
		Radius:           20000,
		PathToNamesData:  "../random_passenger/internal/order_gen/usernames",
	}

	coordGen := coordinate_gen.New(
		cfg.CentralLatitude,
		cfg.CentralLongitude,
	)

	orderGen := order_gen.New(cfg.PathToNamesData, coordGen)

	driverProcessor := drivers_processor.New(coordGen)

	fmt.Println(orderGen, driverProcessor)

	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMathServer(s, server{})

	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
