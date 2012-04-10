package fmod

import (
  "runtime"
)

var critical_in chan func()
var critical_out chan struct{}
// Runs the func, f, on the fmod thread, which is locked to an Os thread.
// This function does not return until after f has returned.
func thread(f func()) {
  critical_in <- f
  <-critical_out
}

func init() {
  critical_in = make(chan func())
  critical_out = make(chan struct{})
  go func() {
    runtime.LockOSThread()
    for {
      f := <-critical_in
      f()
      critical_out <- struct{}{}
    }
  } ()
}
