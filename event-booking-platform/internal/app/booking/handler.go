package booking

import (
    "context"
    "errors"

    bookingpb "path/to/your/repo/api/proto/booking"
    eventpb "path/to/your/repo/api/proto/event"
)

type BookingServiceServer struct {
    bookingpb.UnimplementedBookingServiceServer
    bookings    map[string]*bookingpb.GetBookingResponse
    eventClient eventpb.EventServiceClient
}

func NewBookingServiceServer(eventClient eventpb.EventServiceClient) *BookingServiceServer {
    return &BookingServiceServer{
        bookings:    make(map[string]*bookingpb.GetBookingResponse),
        eventClient: eventClient,
    }
}

func (s *BookingServiceServer) CreateBooking(ctx context.Context, req *bookingpb.CreateBookingRequest) (*bookingpb.CreateBookingResponse, error) {
    event, err := s.eventClient.GetEvent(ctx, &eventpb.GetEventRequest{Id: req.EventId})
    if err != nil {
        return nil, errors.New("failed to fetch event details")
    }

    id := "booking_" + generateID()
    booking := &bookingpb.GetBookingResponse{
        Id:      id,
        UserId:  req.UserId,
        Event:   event,
    }
    s.bookings[id] = booking
    return &bookingpb.CreateBookingResponse{
        Id:      id,
        Message: "Booking created successfully",
    }, nil
}

func (s *BookingServiceServer) GetBooking(ctx context.Context, req *bookingpb.GetBookingRequest) (*bookingpb.GetBookingResponse, error) {
    booking, exists := s.bookings[req.Id]
    if !exists {
        return nil, errors.New("booking not found")
    }
    return booking, nil
}
