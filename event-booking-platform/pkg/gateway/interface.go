package gateway

import (
    "context"

    "google.golang.org/grpc"
)

type GatewayService interface {
    RegisterEventServiceGateway(ctx context.Context, endpoint string, opts []grpc.DialOption) error
    RegisterBookingServiceGateway(ctx context.Context, endpoint string, opts []grpc.DialOption) error
    StartHTTPServer(address string) error
}
