package stream

import (
	"io"
	"sync"

	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//PlayerState holds the stats of the player
type PlayerState struct {
	uuid  string
	score int32
}

type PlayerScoreCache interface {
	SaveScore(uuid string, score int32) (*PlayerState, error)
}

type InMemoryPlyerScoreCache struct {
	lock  sync.RWMutex
	cache map[string]*PlayerState
}

func NewInMemoryPlyerScoreCache() InMemoryPlyerScoreCache {
	return InMemoryPlyerScoreCache{
		cache: make(map[string]*PlayerState),
	}
}

func (c *InMemoryPlyerScoreCache) SaveScore(uuid string, score int32) (*PlayerState, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	state, found := c.cache[uuid]
	if found {
		log.Printf("[Info] Player %v have added score of %v to its last tally of %v\n", state.uuid, state.score, score)
		last_score := state.score
		state.score = last_score + score

		return state, nil
	}
	state = &PlayerState{
		uuid:  uuid,
		score: score,
	}
	log.Printf("[Info] Plyer %v have been added to the score board with score %v", uuid, score)

	c.cache[uuid] = state

	return state, nil
}

type Server struct {
	cache InMemoryPlyerScoreCache
	UnimplementedStreamServiceServer
}

func NewServer() *Server {
	return &Server{
		cache: NewInMemoryPlyerScoreCache(),
	}
}

func (s *Server) TrackScore(stream StreamService_TrackScoreServer) error {
	log.Print("[DEBUG] entry: TrackScore()")
	req, err := stream.Recv()
	if err != nil {
		log.Printf("[ERROR] Could not read data from request: %v\n", err.Error())
		return status.Error(codes.Unknown, "Could not read data request\n")
	}

	playerName := req.GetInfo().GetPlayerId()
	if len(playerName) < 5 {
		log.Printf("[ERROR] Invalid plyer name %v, should contain at least 5 character\n", playerName)
		return status.Errorf(codes.InvalidArgument, "Player name should contain at least 5 character\n")
	}

	total_score := int32(0)
	log.Printf("[INFO] Waiting for %v's scores \n", playerName)
	for {

		log.Println("[DEBUG] Waiting for data")

		err = stream.Context().Err()
		if err != nil {
			break
		}

		req, err := stream.Recv()
		if io.EOF == err {
			log.Printf("[INFO] Client explicitely closed the stream.\n")
			break
		}

		if err != nil {
			log.Printf("[ERROR] could not recieve score events: %v\n", err.Error())
			return status.Errorf(codes.Internal, "Could not recieve events\n")
		}

		score := int32(req.GetScore())
		log.Printf("[INFO] Player %v got %v score.\n", playerName, score)
		if score == 0 {
			log.Printf("[WARNING] Custom marker to inform server that this is the end of stream")
			break
		}
		total_score = total_score + int32(score)
	}

	state, err := s.cache.SaveScore(playerName, total_score)

	if err != nil {
		log.Printf("[ERROR] Failed to save player: %v", err.Error())
		return status.Errorf(codes.Internal, "Failed to save player %v", playerName)
	}

	res := &PlayerScore{
		Info: &PlayerInfo{
			PlayerId: state.uuid,
		},
		TotalScore: state.score,
	}

	err = stream.SendAndClose(res)
	if err != nil {
		log.Printf("[Error] Failed to send response and close stream, %v\n", err.Error())
		return status.Error(codes.Internal, "Failed to send response and close stream")
	}

	log.Printf("[INFO] Player %v saved successfully\n", playerName)
	log.Print("[DEBUG] exit: TrackScore()")
	return nil
}
