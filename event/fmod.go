package fmod_event

// #cgo darwin  LDFLAGS: -Llib -lfmodevent
// #cgo CFLAGS: -Iinc -I../ex/inc
// #include "fmod_event.h"
// #include "stdlib.h"
import "C"

import (
  "fmt"
  "errors"
  "unsafe"
  "github.com/runningwild/fmod/base"
  "github.com/runningwild/fmod/ex"
)

type EventSystem struct {
  system *C.FMOD_EVENTSYSTEM
}

// FMOD_RESULT F_API FMOD_EventSystem_Create(FMOD_EVENTSYSTEM **eventsystem);
func EventSystemCreate() (*EventSystem, error) {
  var event EventSystem
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_EventSystem_Create(&event.system)
  })
  err := base.ResultToError(ferr)
  if err != nil {
    return nil, err
  }
  return &event, err
}

// FMOD_RESULT F_API FMOD_EventSystem_Init              (FMOD_EVENTSYSTEM *eventsystem, int maxchannels, FMOD_INITFLAGS flags, void *extradriverdata, FMOD_EVENT_INITFLAGS eventflags);
func (sys *EventSystem) Init(max_channels int, init_flags fmod_ex.InitFlags, extradriverdata interface{}, event_init_flags InitFlags) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_EventSystem_Init(sys.system, C.int(max_channels), C.FMOD_INITFLAGS(init_flags), unsafe.Pointer(uintptr(0)), C.FMOD_EVENT_INITFLAGS(event_init_flags))
  })
  return base.ResultToError(ferr)
}

// TODO: Bind this to a finalizer
// FMOD_RESULT F_API FMOD_EventSystem_Release           (FMOD_EVENTSYSTEM *eventsystem);
func (sys *EventSystem) Release() error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_EventSystem_Release(sys.system)
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_EventSystem_Update            (FMOD_EVENTSYSTEM *eventsystem);
func (sys *EventSystem) Update() error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_EventSystem_Update(sys.system)
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_EventSystem_SetMediaPath      (FMOD_EVENTSYSTEM *eventsystem, const char *path);
func (sys *EventSystem) SetMediaPath(path string) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    cpath := C.CString(path)
    ferr = C.FMOD_EventSystem_SetMediaPath(sys.system, cpath)
    C.free(unsafe.Pointer(cpath))
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_EventSystem_SetPluginPath     (FMOD_EVENTSYSTEM *eventsystem, const char *path);
func (sys *EventSystem) SetPluginPath(path string) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    cpath := C.CString(path)
    ferr = C.FMOD_EventSystem_SetPluginPath(sys.system, cpath)
    C.free(unsafe.Pointer(cpath))
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_EventSystem_GetVersion        (FMOD_EVENTSYSTEM *eventsystem, unsigned int *version);
func (sys *EventSystem) GetVersion() (uint32, error) {
  var ferr C.FMOD_RESULT
  var version C.uint
  base.Thread(func() {
    ferr = C.FMOD_EventSystem_GetVersion(sys.system, &version)
  })
  return uint32(version), base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_EventSystem_GetInfo           (FMOD_EVENTSYSTEM *eventsystem, FMOD_EVENT_SYSTEMINFO *info);
func (sys *EventSystem) GetInfo() error {
  return errors.New("fmod_event.GetInfo() is not implemented.")
}

// FMOD_RESULT F_API FMOD_EventSystem_GetSystemObject   (FMOD_EVENTSYSTEM *eventsystem, FMOD_SYSTEM **system);
func (sys *EventSystem) GetSystemObject() error {
  return errors.New("fmod_event.GetSystemObject() is not implemented.")
}

// FMOD_RESULT F_API FMOD_EventSystem_GetMusicSystem    (FMOD_EVENTSYSTEM *eventsystem, FMOD_MUSICSYSTEM **musicsystem);
func (sys *EventSystem) GetMusicSystem() error {
  return errors.New("fmod_event.GetMusicSystem() is not implemented.")
}

// FMOD_RESULT F_API FMOD_EventSystem_SetLanguage       (FMOD_EVENTSYSTEM *eventsystem, const char *language);
func (sys *EventSystem) SetLanguage() error {
  return errors.New("fmod_event.SetLanguage() is not implemented.")
}

// FMOD_RESULT F_API FMOD_EventSystem_GetLanguage       (FMOD_EVENTSYSTEM *eventsystem, char *language);
func (sys *EventSystem) GetLanguage() error {
  return errors.New("fmod_event.GetLanguage() is not implemented.")
}

// FMOD_RESULT F_API FMOD_EventSystem_RegisterDSP       (FMOD_EVENTSYSTEM *eventsystem, FMOD_DSP_DESCRIPTION *description, unsigned int *handle);
func (sys *EventSystem) RegisterDSP() error {
  return errors.New("fmod_event.RegisterDSP() is not implemented.")
}

func main() {
  var p unsafe.Pointer
  fmt.Printf("%v\n", p)
  base.Thread(func() {})
}
