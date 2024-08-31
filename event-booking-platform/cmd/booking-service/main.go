package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"

    bookingpb "path/to/your/repo/api/proto/booking"
    eventpb "path/to/your/repo/api/proto/event"
    "path/to/your/repo/internal/app/booking"
)

func main() {
    listener, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("Failed to listen on port 50052: %v", err)
    }

    grpcServer := grpc.NewServer()
    eventClient := eventpb.NewEventServiceClient(grpcServer)
    bookingServer := booking.NewBookingServiceServer(eventClient)

    bookingpb.RegisterBookingServiceServer(grpcServer, bookingServer)
    reflection.Register(grpcServer)

    log.Println("Starting Booking gRPC server on port 50052...")
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatalf("Failed to serve gRPC server: %v", err)
    }
}
