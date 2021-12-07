package stream_test

import (
	context "context"
	"errors"
	"fmt"
	"testing"

	"github.com/imhshekhar47/go-grpc-api/pb/stream"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
)

type StreamMock struct {
	grpc.ServerStream
	ctx                   context.Context
	recvToServerChannel   chan *stream.PlayerScoreEvent
	sentFromServerChannel chan *stream.PlayerScore
}

func NewStreamMock() *StreamMock {
	return &StreamMock{
		ctx:                   context.Background(),
		recvToServerChannel:   make(chan *stream.PlayerScoreEvent, 10),
		sentFromServerChannel: make(chan *stream.PlayerScore, 10),
	}
}

func (ms *StreamMock) SendAndClose(m *stream.PlayerScore) error {
	ms.sentFromServerChannel <- m
	return nil
}

func (ms *StreamMock) Recv() (*stream.PlayerScoreEvent, error) {
	req, more := <-ms.recvToServerChannel
	if !more {
		return nil, errors.New("empty")
	}

	return req, nil
}

func (ms *StreamMock) Context() context.Context {
	return ms.ctx
}

func (ms *StreamMock) SentFromClient(evt *stream.PlayerScoreEvent) error {
	ms.recvToServerChannel <- evt
	return nil
}

func (ms *StreamMock) RecvToClient() (*stream.PlayerScore, error) {
	resp, more := <-ms.sentFromServerChannel
	if !more {
		return nil, errors.New("Empty")
	}

	return resp, nil
}

func TestTrackScore(t *testing.T) {
	g := NewGomegaWithT(t)
	mockStream := NewStreamMock()
	go func() {
		server := stream.NewServer()

		err := server.TrackScore(mockStream)
		if err != nil {
			t.Errorf("Filed with error %v:\n", err)
		}

		close(mockStream.recvToServerChannel)
		close(mockStream.sentFromServerChannel)
		t.Log("Server up and running")
	}()

	limit := 10
	var event *stream.PlayerScoreEvent
	for i := 0; i <= limit; i++ {

		if i == 0 {
			event = &stream.PlayerScoreEvent{
				Data: &stream.PlayerScoreEvent_Info{
					Info: &stream.PlayerInfo{
						PlayerId: fmt.Sprintf("player_%v", i),
					},
				},
			}
		} else if i == limit {
			// TODO: How to notify server that stream is complete
			event = &stream.PlayerScoreEvent{
				Data: &stream.PlayerScoreEvent_Score{
					Score: int32(0),
				},
			}
		} else {
			event = &stream.PlayerScoreEvent{
				Data: &stream.PlayerScoreEvent_Score{
					Score: int32(10 + i),
				},
			}
		}

		t.Logf("Sending event %v to server\n", i)
		err := mockStream.SentFromClient(event)
		if err != nil {
			t.Errorf("Filed to sent event %v to server with error %v\n", i, err.Error())
			return
		}
	}

	resp, err := mockStream.RecvToClient()
	if err != nil {
		t.Errorf("Failed to receive response from server with error %v\n", err.Error())
	}
	g.Expect(resp).ToNot(BeNil(), "Response should not be nil")
	g.Expect(resp.TotalScore).To(BeEquivalentTo(135), "Total score should match")
	t.Logf("Response[Player=%v, total_score=%v]\n", resp.Info.PlayerId, resp.TotalScore)
}
