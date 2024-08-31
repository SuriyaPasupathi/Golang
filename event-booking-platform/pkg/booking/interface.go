package booking

import (
    "context"

    bookingpb "path/to/your/repo/api/proto/booking"
)

type BookingService interface {
    CreateBooking(ctx context.Context, userID, eventID string) (*bookingpb.CreateBookingResponse, error)
    GetBooking(ctx context.Context, id string) (*bookingpb.GetBookingResponse, error)
}
