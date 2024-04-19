package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "time"

    "google.golang.org/grpc"
    pb "Tierra/gems_grpc" 
)

type TierraServer struct {
	pb.UnimplementedTierraServer
    ATCounter     int32
    MPCounter     int32
    ATMaxCapacity int32 // Maximum capacity for AT ammo
    MPMaxCapacity int32 // Maximum capacity for MP ammo
}

func (s *TierraServer) SolicitarM(ctx context.Context, req *pb.Solicitud) (*pb.Respuesta, error) {

    if s.hasEnoughAmmo(req.AT, req.MP) {
        s.subtractAmmo(req.AT, req.MP)
        fmt.Printf("\nRecepcion de solicitud desde equipo %d, %d AT y %d MP -- APROBADA -- AT EN SISTEMA: %d ; MP EN SISTEMA: %d\n", req.ID,req.AT,req.MP,s.ATCounter, s.MPCounter)
        return &pb.Respuesta{Success: true}, nil
    } else {
        fmt.Printf("\nRecepcion de solicitud desde equipo %d, %d AT y %d MP -- DENEGADA -- AT EN SISTEMA: %d ; MP EN SISTEMA: %d\n", req.ID, req.AT, req.MP, s.ATCounter, s.MPCounter)
        return &pb.Respuesta{Success: false}, nil
    }
}

func (s *TierraServer) hasEnoughAmmo(ATRequested, MPRequested int32) bool {
    return s.ATCounter >= ATRequested && s.MPCounter >= MPRequested
}

func (s *TierraServer) subtractAmmo(ATRequested, MPRequested int32) {
    s.ATCounter -= ATRequested
    s.MPCounter -= MPRequested
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    server := grpc.NewServer()
    tierraServer := &TierraServer{
        ATMaxCapacity: 50, // Maximum capacity for AT ammo
        MPMaxCapacity: 20, // Maximum capacity for MP ammo
    }
    pb.RegisterTierraServer(server, tierraServer)

    // Start a goroutine to periodically add ammo
    go func() {
        for {
            time.Sleep(5 * time.Second)
            tierraServer.addAmmo(10, 5) // Add 10 AT ammo and 5 MP ammo
        }
    }()

    if err := server.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

func (s *TierraServer) addAmmo(ATToAdd, MPToAdd int32) {
    // Add AT ammo
    if s.ATCounter+ATToAdd <= s.ATMaxCapacity {
        s.ATCounter += ATToAdd
    } else {
        s.ATCounter = s.ATMaxCapacity
    }

    // Add MP ammo
    if s.MPCounter+MPToAdd <= s.MPMaxCapacity {
        s.MPCounter += MPToAdd
    } else {
        s.MPCounter = s.MPMaxCapacity
    }
}
