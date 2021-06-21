package saver

import (
	"context"
	"errors"
	"time"

	"github.com/ozoncp/ocp-classroom-api/internal/flusher"
	"github.com/ozoncp/ocp-classroom-api/internal/models"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Saver is utility that saves classrooms
type Saver interface {
	Init(ctx context.Context)
	Save(classroom models.Classroom)
	Close()
}

// Policy defines behavior of Saver when classrooms in RAM are overflowed
type Policy int

const (
	// Policy_DropAll defines behavior to drop all classrooms from RAM
	Policy_DropAll = iota

	// Policy_DropFirst defines behavior to drop only first classroom in RAM
	Policy_DropFirst
)

// saver is thread-safe implementation of Saver interface
// that keeps classrooms in RAM and flushes them to storage by timer
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

// New returns thread-safe Saver instance that keeps classrooms in RAM and flushes them to storage by timer
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

// Init runs work in another goroutine
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

				logSaver().Msg("closing")

				s.isClosedCh <- struct{}{}

				return
			}

		}
	}()
}

// Save puts classroom through channel to RAM
func (s *saver) Save(classroom models.Classroom) {
	s.classroomCh <- classroom
}

// Close finishes work and flushes all classrooms in RAM to storage before closing
func (s *saver) Close() {

	s.ticker.Stop()

	s.shouldCloseCh <- struct{}{}

	<-s.isClosedCh
}

// save appends new classroom to RAM in order of policy
func (s *saver) save(classroom *models.Classroom) {

	if uint(len(s.classrooms)) == s.capacity {

		if s.policy == Policy_DropAll {

			s.classrooms = s.classrooms[:0]

			logSaver().Msg("droping all")

		} else if s.policy == Policy_DropFirst {

			s.classrooms = s.classrooms[1:]

			logSaver().Msg("droping first")
		}
	}

	s.classrooms = append(s.classrooms, *classroom)

	logSaver().Msgf("saving, classrooms: %v", s.classrooms)
}

// flush flushes classrooms in RAM to storage
func (s *saver) flush(ctx context.Context) {

	s.classrooms = s.flusher.Flush(ctx, nil, s.classrooms)

	logSaver().Msgf("flushing, classrooms: %v", s.classrooms)
}

// logSaver is convenient internal function for logging
func logSaver() *zerolog.Event {
	return log.Debug().Str("package", "saver")
}
