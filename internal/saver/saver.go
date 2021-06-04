package saver

import (
	"errors"
	"fmt"
	"time"

	"github.com/ozoncp/ocp-classroom-api/internal/flusher"
	"github.com/ozoncp/ocp-classroom-api/internal/models"
)

type Saver interface {
	Init() error
	Save(classroom models.Classroom)
	Close()
}

type Policy int

const (
	Policy_DropAll = iota
	Policy_DropFirst
)

func NewSaver(capacity uint, policy Policy, interval time.Duration, flusher flusher.Flusher) Saver {

	return &saver{
		capacity: capacity,
		policy:   policy,
		interval: interval,
		flusher:  flusher,
	}
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

func (s *saver) Init() error {

	if s.capacity < 1 {
		return errors.New("capacity < 1")
	}

	if s.flusher == nil {
		return errors.New("flusher is nil")
	}

	if s.interval == 0 {
		return errors.New("interval == 0")
	}

	s.ticker = time.NewTicker(s.interval)

	s.classrooms = make([]models.Classroom, 0, s.capacity)
	s.classroomCh = make(chan models.Classroom)

	s.shouldCloseCh = make(chan struct{})
	s.isClosedCh = make(chan struct{})

	go loop(s)

	return nil
}

func (s *saver) Save(classroom models.Classroom) {

	s.classroomCh <- classroom
}

func (s *saver) Close() {

	s.ticker.Stop()

	s.shouldCloseCh <- struct{}{}

	<-s.isClosedCh

	close(s.classroomCh)
	close(s.shouldCloseCh)
	close(s.isClosedCh)
}

func loop(s *saver) {

	for {
		select {

		case classroom := <-s.classroomCh:

			s.loop_save(&classroom)

		case <-s.ticker.C:

			s.loop_flush()

		case <-s.shouldCloseCh:

			s.loop_flush()

			fmt.Println("closing...")

			s.isClosedCh <- struct{}{}

			return
		}

	}
}

func (s *saver) loop_save(classroom *models.Classroom) {

	if uint(len(s.classrooms)) == s.capacity {

		if s.policy == Policy_DropAll {

			s.classrooms = s.classrooms[:0]

			fmt.Println("droping all")

		} else if s.policy == Policy_DropFirst {

			s.classrooms = s.classrooms[1:]

			fmt.Println("droping first")
		}
	}

	s.classrooms = append(s.classrooms, *classroom)

	fmt.Println("saving... classrooms:", s.classrooms)
}

func (s *saver) loop_flush() {
	s.classrooms = s.flusher.Flush(s.classrooms)

	fmt.Println("flushing... classrooms:", s.classrooms)
}
