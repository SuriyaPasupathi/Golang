package gateway

import (
    "context"

    "google.golang.org/grpc"
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type GatewayClient struct {
    mux *runtime.ServeMux
}

func NewGatewayClient() *GatewayClient {
    return &GatewayClient{
        mux: runtime.NewServeMux(),
    }
}

func (c *GatewayClient) RegisterEventServiceGateway(ctx context.Context, endpoint string, opts []grpc.DialOption) error {
    return eventpb.RegisterEventServiceHandlerFromEndpoint(ctx, c.mux, endpoint, opts)
}

func (c *GatewayClient) RegisterBookingServiceGateway(ctx context.Context, endpoint string, opts []grpc.DialOption) error {
    return bookingpb.RegisterBookingServiceHandlerFromEndpoint(ctx, c.mux, endpoint, opts)
}

func (c *GatewayClient) StartHTTPServer(address string) error {
    return http.ListenAndServe(address, c.mux)
}
