package repository

type EventRepositoryImpl struct {
}

func NewEventRepository() EventRepository {
	return &EventRepositoryImpl{}
}

