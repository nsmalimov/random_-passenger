package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"random_passenger_driver/configs"
	"random_passenger_driver/internal/coordinategen"
	"random_passenger_driver/internal/driversprocessor"
	"random_passenger_driver/internal/ordergen"
	pb "random_passenger_driver/internal/proto"

	"google.golang.org/grpc"
)

type server struct {
	driverProcessor *driversprocessor.DriversProcessor
	orderGen        *ordergen.OrderGenService
	config          *configs.Config
}

func (s *server) randInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func (s server) Driver(srv pb.PassengerDriver_DriverServer) error {
	log.Println("Start Driver server")

	ctx := srv.Context()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		driver := s.driverProcessor.Driver()

		resp := pb.ResponseDriver{
			Id:        driver.ID,
			CarModel:  driver.CarModel,
			Latitude:  driver.Latitude,
			Longitude: driver.Longitude,
		}
		if err := srv.Send(&resp); err != nil {
			log.Printf("Error when try srv.Send, err: %s", err)
		}

		rSleepSec := s.randInt(s.config.MinSecSleepDriver, s.config.MaxSecSleepDriver)
		time.Sleep(time.Duration(rSleepSec) * time.Second)
	}
}

func (s server) Order(srv pb.PassengerDriver_OrderServer) error {
	log.Println("Start Order server")

	ctx := srv.Context()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		order := s.orderGen.GenOrder()

		resp := pb.ResponseOrder{
			Id:            order.ID,
			Username:      order.Username,
			LatitudeTo:    order.LatitudeTo,
			LongitudeTo:   order.LongitudeTo,
			LatitudeFrom:  order.LatitudeFrom,
			LongitudeFrom: order.LongitudeFrom,
		}
		if err := srv.Send(&resp); err != nil {
			log.Printf("Error when try srv.Send, err: %s", err)
		}

		rSleepSec := s.randInt(s.config.MinSecSleepOrder, s.config.MaxSecSleepOrder)
		time.Sleep(time.Duration(rSleepSec) * time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	configFilepathF := flag.String("config-file", "", "path to .yaml config file")

	flag.Parse()

	configFilepath := *configFilepathF

	if configFilepath == "" {
		log.Printf("No configFilepath set it by flag, -h for info")
		return
	}

	cfg, err := configs.New(configFilepath)

	if err != nil {
		log.Fatalf("Error when try configs.New, err: %s", err)
	}

	coordGen := coordinategen.New(
		cfg.CentralLatitude,
		cfg.CentralLongitude,
	)

	orderGenS := ordergen.New(cfg.PathToNamesData, coordGen)
	driverProcessor := driversprocessor.New(coordGen)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatalf("Error when try net.Listen, err: %s", err)
	}

	s := grpc.NewServer()
	pb.RegisterPassengerDriverServer(s, server{
		orderGen:        orderGenS,
		driverProcessor: driverProcessor,
		config:          cfg,
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error when try s.Serve, err: %s", err)
	}
}
