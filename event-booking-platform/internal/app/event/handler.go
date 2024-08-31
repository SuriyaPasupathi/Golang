package event

import (
    "context"
    "errors"

    eventpb "path/to/your/repo/api/proto/event"
)

type EventServiceServer struct {
    eventpb.UnimplementedEventServiceServer
    events map[string]*eventpb.GetEventResponse
}

func NewEventServiceServer() *EventServiceServer {
    return &EventServiceServer{events: make(map[string]*eventpb.GetEventResponse)}
}

func (s *EventServiceServer) CreateEvent(ctx context.Context, req *eventpb.CreateEventRequest) (*eventpb.CreateEventResponse, error) {
    id := "event_" + generateID()
    event := &eventpb.GetEventResponse{
        Id:       id,
        Name:     req.Name,
        Location: req.Location,
        Date:     req.Date,
    }
    s.events[id] = event
    return &eventpb.CreateEventResponse{
        Id:      id,
        Message: "Event created successfully",
    }, nil
}

func (s *EventServiceServer) GetEvent(ctx context.Context, req *eventpb.GetEventRequest) (*eventpb.GetEventResponse, error) {
    event, exists := s.events[req.Id]
    if !exists {
        return nil, errors.New("event not found")
    }
    return event, nil
}
