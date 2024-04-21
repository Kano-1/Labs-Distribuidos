package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "Lab3/Proto"

	"google.golang.org/grpc"
)

type ServicioMunicionServer struct {
	pb.UnimplementedServicioMunicionServer
	ATCounter     int32
	MPCounter     int32
	ATMaxCapacity int32 // Maximum capacity for AT ammo
	MPMaxCapacity int32 // Maximum capacity for MP ammo
}

func (s *ServicioMunicionServer) SolicitarM(ctx context.Context, req *pb.SolicitarMunicion) (*pb.RespuestaMunicion, error) {

	if s.revisarMunicion(req.AT, req.MP) {
		s.restarMunicion(req.AT, req.MP)
		fmt.Printf("\nRecepcion de solicitud desde equipo %d, %d AT y %d MP -- APROBADA -- \nAT EN SISTEMA: %d ; MP EN SISTEMA: %d\n", req.ID, req.AT, req.MP, s.ATCounter, s.MPCounter)
		return &pb.RespuestaMunicion{Success: true}, nil
	} else {
		fmt.Printf("\nRecepcion de solicitud desde equipo %d, %d AT y %d MP -- DENEGADA -- \nAT EN SISTEMA: %d ; MP EN SISTEMA: %d\n", req.ID, req.AT, req.MP, s.ATCounter, s.MPCounter)
		return &pb.RespuestaMunicion{Success: false}, nil
	}
}

func (s *ServicioMunicionServer) revisarMunicion(ATRequested, MPRequested int32) bool {
	return s.ATCounter >= ATRequested && s.MPCounter >= MPRequested
}

func (s *ServicioMunicionServer) restarMunicion(ATRequested, MPRequested int32) {
	s.ATCounter -= ATRequested
	s.MPCounter -= MPRequested
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	servicioMunicionServer := &ServicioMunicionServer{
		ATMaxCapacity: 50, // Maximum capacity for AT ammo
		MPMaxCapacity: 20, // Maximum capacity for MP ammo
	}
	pb.RegisterServicioMunicionServer(server, servicioMunicionServer)

	// Start a goroutine to periodically add ammo
	go func() {
		for {
			time.Sleep(5 * time.Second)
			servicioMunicionServer.crearMunicion(10, 5) // Add 10 AT ammo and 5 MP ammo
		}
	}()

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *ServicioMunicionServer) crearMunicion(ATToAdd, MPToAdd int32) {
	// Suma AT
	if s.ATCounter+ATToAdd <= s.ATMaxCapacity {
		s.ATCounter += ATToAdd
	} else {
		s.ATCounter = s.ATMaxCapacity
	}

	// Suma MP
	if s.MPCounter+MPToAdd <= s.MPMaxCapacity {
		s.MPCounter += MPToAdd
	} else {
		s.MPCounter = s.MPMaxCapacity
	}
}
