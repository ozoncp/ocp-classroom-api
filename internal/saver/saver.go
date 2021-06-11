package saver

import (
	"errors"
	"time"

	"github.com/ozoncp/ocp-classroom-api/internal/flusher"
	"github.com/ozoncp/ocp-classroom-api/internal/models"

	"github.com/rs/zerolog/log"
)

type Saver interface {
	Init()
	Save(classroom models.Classroom)
	Close()
}

type Policy int

const (
	Policy_DropAll = iota
	Policy_DropFirst
)

func New(capacity uint, policy Policy, interval time.Duration, flusher flusher.Flusher) (Saver, error) {

	if capacity == 0 {
		return nil, errors.New("capacity is 0")
	}

	if interval == 0 {
		return nil, errors.New("interval is 0")
	}

	if flusher == nil {
		return nil, errors.New("flusher is nil")
	}

	return &saver{
		capacity: capacity,
		policy:   policy,
		interval: interval,
		flusher:  flusher,
	}, nil
}

type saver struct {
	capacity uint
	policy   Policy
	interval time.Duration
	flusher  flusher.Flusher

	ticker *time.Ticker

	classrooms  []models.Classroom
	classroomCh chan models.Classroom

	shouldCloseCh chan struct{}
	isClosedCh    chan struct{}
}

func (s *saver) Init() {

	s.ticker = time.NewTicker(s.interval)

	s.classrooms = make([]models.Classroom, 0, s.capacity)
	s.classroomCh = make(chan models.Classroom)

	s.shouldCloseCh = make(chan struct{})
	s.isClosedCh = make(chan struct{})

	go loop(s)
}

func (s *saver) Save(classroom models.Classroom) {
	s.classroomCh <- classroom
}

func (s *saver) Close() {

	s.ticker.Stop()

	s.shouldCloseCh <- struct{}{}

	<-s.isClosedCh
}

func loop(s *saver) {

	for {
		select {

		case classroom := <-s.classroomCh:

			s.save(&classroom)

		case <-s.ticker.C:

			s.flush()

		case <-s.shouldCloseCh:

			s.flush()

			log.Debug().Str("package", "saver").Msg("closing")

			s.isClosedCh <- struct{}{}

			return
		}

	}
}

func (s *saver) save(classroom *models.Classroom) {

	if uint(len(s.classrooms)) == s.capacity {

		if s.policy == Policy_DropAll {

			s.classrooms = s.classrooms[:0]

			log.Debug().Str("package", "saver").Msg("droping all")

		} else if s.policy == Policy_DropFirst {

			s.classrooms = s.classrooms[1:]

			log.Debug().Str("package", "saver").Msg("droping first")
		}
	}

	s.classrooms = append(s.classrooms, *classroom)

	log.Debug().Str("package", "saver").Msgf("saving, classrooms: %v", s.classrooms)
}

func (s *saver) flush() {
	s.classrooms = s.flusher.Flush(s.classrooms)

	log.Debug().Str("package", "saver").Msg("flushing")
}
