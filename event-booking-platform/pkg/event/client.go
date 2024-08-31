package event

import (
    "context"

    eventpb "path/to/your/repo/api/proto/event"
    "google.golang.org/grpc"
)

type EventClient struct {
    client eventpb.EventServiceClient
}

func NewEventClient(conn *grpc.ClientConn) *EventClient {
    return &EventClient{
        client: eventpb.NewEventServiceClient(conn),
    }
}

func (c *EventClient) CreateEvent(ctx context.Context, name, location string, date *eventpb.Timestamp) (*eventpb.CreateEventResponse, error) {
    return c.client.CreateEvent(ctx, &eventpb.CreateEventRequest{
        Name:     name,
        Location: location,
        Date:     date,
    })
}

func (c *EventClient) GetEvent(ctx context.Context, id string) (*eventpb.GetEventResponse, error) {
    return c.client.GetEvent(ctx, &eventpb.GetEventRequest{
        Id: id,
    })
}
