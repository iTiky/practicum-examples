package eventservice

type Event struct {
	Id int
}

//go:generate mockery --name=Creator
type Creator interface {
	Create() Event
}

type Service struct {
	creator Creator
}

func New(creator Creator) *Service {
	return &Service{
		creator: creator,
	}
}

func (s *Service) NewEvent() Event {
	return s.creator.Create()
}
