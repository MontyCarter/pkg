/*
Copyright 2019 The Knative Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testing

import (
	"time"
)

// FakeClock - a fake Clock
type FakeClock struct {
	Time    time.Time
	sleeper chan bool
}

// Now - returns the current time of the clock
func (c FakeClock) Now() time.Time {
	return c.Time
}

// Sleep - checks to see if duration has passed since being called
func (c FakeClock) Sleep(duration time.Duration) {
	startingTime := c.Now()
	for {
		if c.Now().Sub(startingTime) > duration {
			return
		}
		c.sleeper <- true
	}
}

// SetTime - Sets the current time of the FakeClock
func (c *FakeClock) SetTime(time time.Time) {
	c.Time = time
	<-c.sleeper
}

// NewFakeClock - creates a new fake clock, with starting time of time.Now()
func NewFakeClock() *FakeClock {
	now := time.Now()
	return &FakeClock{Time: now, sleeper: make(chan bool)}
}
