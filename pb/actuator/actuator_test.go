package actuator_test

import (
	context "context"
	"testing"

	empty "github.com/golang/protobuf/ptypes/empty"
	core "github.com/imhshekhar47/go-grpc-api/core"
	actuator "github.com/imhshekhar47/go-grpc-api/pb/actuator"
	. "github.com/onsi/gomega"
)

func TestGetHealth(t *testing.T) {
	g := NewGomegaWithT(t)
	server := actuator.NewServer(core.DefaultAppConfig)
	// call
	response, err := server.GetHealth(context.Background(), new(empty.Empty))
	if err != nil {
		t.Errorf("Failed with error: %v\n", err)
	}
	g.Expect(response).ToNot(BeNil(), "Response should not be nil")
	t.Log("Response: ", response)

}

func TestGetInfo(t *testing.T) {
	g := NewGomegaWithT(t)
	server := actuator.NewServer(core.DefaultAppConfig)
	// call
	response, err := server.GetInfo(context.Background(), new(empty.Empty))
	if err != nil {
		t.Errorf("Failed with error: %v\n", err)
	}
	g.Expect(response).ToNot(BeNil(), "Response should not be nil")
	t.Log("Response: ", response)
}
