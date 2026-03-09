package stat

import (
	"Documents/Web_GO/pkg/event"
	"log"
)

type StatServiceDeps struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

type StatService struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

func NewStatService(deps *StatServiceDeps) *StatService {
	return &StatService{
		EventBus:       deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

func (s *StatService) AddClick() {
	for msg := range s.EventBus.Subscribe() {
		if msg.Type == event.LinkVisitedEvent {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Fatal("Bad EventLinkVisited Data: ", msg.Data)
			}
			s.StatRepository.AddClick(id)
		}
	}
}
