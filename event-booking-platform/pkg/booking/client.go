package booking

import (
    "context"

    bookingpb "path/to/your/repo/api/proto/booking"
    "google.golang.org/grpc"
)

type BookingClient struct {
    client bookingpb.BookingServiceClient
}

func NewBookingClient(conn *grpc.ClientConn) *BookingClient {
    return &BookingClient{
        client: bookingpb.NewBookingServiceClient(conn),
    }
}

func (c *BookingClient) CreateBooking(ctx context.Context, userID, eventID string) (*bookingpb.CreateBookingResponse, error) {
    return c.client.CreateBooking(ctx, &bookingpb.CreateBookingRequest{
        UserId:  userID,
        EventId: eventID,
    })
}

func (c *BookingClient) GetBooking(ctx context.Context, id string) (*bookingpb.GetBookingResponse, error) {
    return c.client.GetBooking(ctx, &bookingpb.GetBookingRequest{
        Id: id,
    })
}
