// Copyright 2017 Google Inc. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to writing, software distributed
// under the License is distributed on a "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type scene struct {
	bg        *sdl.Texture
	bird      *bird
	pipes     *pipes
	framerate uint32
}

func newScene(r *sdl.Renderer) (*scene, error) {
	bg, err := img.LoadTexture(r, "res/imgs/background.png")
	if err != nil {
		return nil, fmt.Errorf("could not load background image: %v", err)
	}

	b, err := newBird(r)
	if err != nil {
		return nil, err
	}

	ps, err := newPipes(r)
	if err != nil {
		return nil, err
	}

	return &scene{bg: bg, bird: b, pipes: ps, framerate: 120}, nil
}

func (s *scene) run(r *sdl.Renderer) error {
	errc := make(chan error)
	defer close(errc)
	events := make(chan sdl.Event)
	donec := make(chan bool)
	defer close(donec)

	go func() {
		for {
			select {
			case <-donec:
				close(events)
			default:
				events <- sdl.WaitEvent()
			}
		}
	}()

	wg := sync.WaitGroup{}
	for {
		select {
		case err := <-errc:
			return err
		case <-donec:
			return nil
		default:
			wg.Add(1)
			go func() {
				select {
				case e := <-events:
					if done := s.handleEvent(e); done {
						go func() {
							donec <- done
						}()
					}
				default:
					s.update()

					if s.bird.isDead() {
						sdl.Do(func() {
							drawTitle(r, "Game Over")
						})
						time.Sleep(time.Second)
						s.restart()
					}

					if err := s.paint(r); err != nil {
						errc <- err
					}
				}
				wg.Done()
			}()
			wg.Wait()
		}
	}
}

func (s *scene) handleEvent(event sdl.Event) bool {
	switch event.(type) {
	case *sdl.QuitEvent:
		return true
	case *sdl.MouseButtonEvent:
		s.bird.jump()
	case *sdl.MouseMotionEvent, *sdl.WindowEvent, *sdl.TouchFingerEvent, *sdl.CommonEvent:
	default:
		log.Printf("unknown event %T", event)
	}
	return false
}

func (s *scene) update() {
	s.bird.update()
	s.pipes.update()
	s.pipes.touch(s.bird)
}

func (s *scene) restart() {
	s.bird.restart()
	s.pipes.restart()
}

func (s *scene) paint(r *sdl.Renderer) error {
	r.Clear()
	var err error
	sdl.Do(func() {
		err = r.Copy(s.bg, nil, nil)
	})
	if err != nil {
		return fmt.Errorf("could not copy background: %v", err)
	}
	sdl.Do(func() {
		err = s.bird.paint(r)
	})
	if err != nil {
		return err
	}
	sdl.Do(func() {
		err = s.pipes.paint(r)
	})
	if err != nil {
		return err
	}

	sdl.Do(func() {
		r.Present()
		sdl.Delay(1000 / s.framerate)
	})
	return nil
}

func (s *scene) destroy() {
	s.bg.Destroy()
	s.bird.destroy()
	s.pipes.destroy()
}
