package saver

import (
	"context"
	"errors"
	"time"

	"github.com/ozoncp/ocp-classroom-api/internal/flusher"
	"github.com/ozoncp/ocp-classroom-api/internal/models"

	"github.com/rs/zerolog/log"
)

// TODO: comment everything here

type Saver interface {
	Init(ctx context.Context)
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

		ticker: time.NewTicker(interval),

		classrooms:  make([]models.Classroom, 0, capacity),
		classroomCh: make(chan models.Classroom),

		shouldCloseCh: make(chan struct{}),
		isClosedCh:    make(chan struct{}),
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

func (s *saver) Init(ctx context.Context) {

	go func() {

		for {
			select {

			case classroom := <-s.classroomCh:

				s.save(&classroom)

			case <-s.ticker.C:

				s.flush(ctx)

			case <-s.shouldCloseCh:

				s.flush(ctx)

				log.Debug().Str("package", "saver").Msg("closing")

				s.isClosedCh <- struct{}{}

				return
			}

		}
	}()
}

func (s *saver) Save(classroom models.Classroom) {
	s.classroomCh <- classroom
}

func (s *saver) Close() {

	s.ticker.Stop()

	s.shouldCloseCh <- struct{}{}

	<-s.isClosedCh
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

func (s *saver) flush(ctx context.Context) {

	s.classrooms = s.flusher.Flush(ctx, nil, s.classrooms)

	log.Debug().Str("package", "saver").Msg("flushing")
}
