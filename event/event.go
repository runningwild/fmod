package fmod_event

// #include "fmod_event.h"
// #include "stdlib.h"
import "C"

import (
  "errors"
  "unsafe"
  "runtime"
  "github.com/MobRulesGames/fmod/base"
)

type Event struct {
  event *C.FMOD_EVENT

  // The group this was loaded from, nil if it wasn't loaded from a group.
  group *Group
}

func finalizeEvent(event *Event) {
  event.Release(true, true)
}

// FMOD_RESULT F_API FMOD_Event_Release                 (FMOD_EVENT *event, FMOD_BOOL freeeventdata, FMOD_BOOL waituntilready);
func (e *Event) Release(free_event_data, wait_until_ready bool) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_Event_Release(e.event, makeFmodBool(free_event_data), makeFmodBool(wait_until_ready))
  })
  runtime.SetFinalizer(e, nil)
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_Event_Start                   (FMOD_EVENT *event);
func (e *Event) Start() error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_Event_Start(e.event)
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_Event_Stop                    (FMOD_EVENT *event, FMOD_BOOL immediate);
func (e *Event) Stop(immediate bool) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_Event_Stop(e.event, makeFmodBool(immediate))
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_Event_GetInfo                 (FMOD_EVENT *event, int *index, char **name, FMOD_EVENT_INFO *info);
func (e *Event) GetInfo() (index int, name string, err error) {
  var ferr C.FMOD_RESULT
  var cindex C.int
  var pcname *C.char
  base.Thread(func() {
    ferr = C.FMOD_Event_GetInfo(e.event, &cindex, &pcname, nil)
  })
  if pcname != nil {
    name = C.GoString(pcname)
  }
  return int(cindex), name, base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_Event_GetState                (FMOD_EVENT *event, FMOD_EVENT_STATE *state);
func (e *Event) GetState() (State, error) {
  var ferr C.FMOD_RESULT
  var state C.FMOD_EVENT_STATE
  base.Thread(func() {
    ferr = C.FMOD_Event_GetState(e.event, &state)
  })
  return State(state), base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_Event_GetParentGroup          (FMOD_EVENT *event, FMOD_EVENTGROUP **group);
func (e *Event) GetParentGroup() error {
  return errors.New("Event.GetParentGroup() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_GetChannelGroup         (FMOD_EVENT *event, FMOD_CHANNELGROUP **channelgroup);
func (e *Event) GetChannelGroup() error {
  return errors.New("Event.GetChannelGroup() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_SetCallback             (FMOD_EVENT *event, FMOD_EVENT_CALLBACK callback, void *userdata);
func (e *Event) SetCallback() error {
  return errors.New("Event.SetCallback() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_GetParameter            (FMOD_EVENT *event, const char *name, FMOD_EVENTPARAMETER **parameter);
func (e *Event) GetParameter(name string) (*Param, error) {
  var ferr C.FMOD_RESULT
  var p Param
  base.Thread(func() {
    cname := C.CString(name)
    ferr = C.FMOD_Event_GetParameter(e.event, cname, &p.param)
    C.free(unsafe.Pointer(cname))
  })
  err := base.ResultToError(ferr)
  if err != nil {
    return nil, err
  }
  return &p, nil
}

// FMOD_RESULT F_API FMOD_Event_GetParameterByIndex     (FMOD_EVENT *event, int index, FMOD_EVENTPARAMETER **parameter);
func (e *Event) GetParameterByIndex() error {
  return errors.New("Event.GetParameterByIndex() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_GetNumParameters        (FMOD_EVENT *event, int *numparameters);
func (e *Event) GetNumParameters() error {
  return errors.New("Event.GetNumParameters() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_GetProperty             (FMOD_EVENT *event, const char *propertyname, void *value, FMOD_BOOL this_instance);
func (e *Event) GetProperty() error {
  return errors.New("Event.GetProperty() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_GetPropertyByIndex      (FMOD_EVENT *event, int propertyindex, void *value, FMOD_BOOL this_instance);
func (e *Event) GetPropertyByIndex() error {
  return errors.New("Event.GetPropertyByIndex() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_SetProperty             (FMOD_EVENT *event, const char *propertyname, void *value, FMOD_BOOL this_instance);
func (e *Event) SetProperty() error {
  return errors.New("Event.SetProperty() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_SetPropertyByIndex      (FMOD_EVENT *event, int propertyindex, void *value, FMOD_BOOL this_instance);
func (e *Event) SetPropertyByIndex() error {
  return errors.New("Event.SetPropertyByIndex() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_GetNumProperties        (FMOD_EVENT *event, int *numproperties);
func (e *Event) GetNumProperties() error {
  return errors.New("Event.GetNumProperties() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_GetPropertyInfo         (FMOD_EVENT *event, int *propertyindex, char **propertyname, FMOD_EVENTPROPERTY_TYPE *type);
func (e *Event) GetPropertyInfo() error {
  return errors.New("Event.GetPropertyInfo() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_GetCategory             (FMOD_EVENT *event, FMOD_EVENTCATEGORY **category);
func (e *Event) GetCategory() error {
  return errors.New("Event.GetCategory() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_SetVolume         (FMOD_EVENT *event, float volume);
func (e *Event) SetVolume(volume float64) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_Event_SetVolume(e.event, C.float(volume))
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_Event_GetVolume         (FMOD_EVENT *event, float *volume);
func (e *Event) GetVolume() (float64, error) {
  var ferr C.FMOD_RESULT
  var volume C.float
  base.Thread(func() {
    ferr = C.FMOD_Event_GetVolume(e.event, &volume)
  })
  return float64(volume), base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_Event_SetPitch                (FMOD_EVENT *event, float pitch, FMOD_EVENT_PITCHUNITS units);
func (e *Event) SetPitch() error {
  return errors.New("Event.SetPitch() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_GetPitch                (FMOD_EVENT *event, float *pitch, FMOD_EVENT_PITCHUNITS units);
func (e *Event) GetPitch() error {
  return errors.New("Event.GetPitch() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_SetPaused         (FMOD_EVENT *event, FMOD_BOOL paused);
func (event *Event) SetPaused(paused bool) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_Event_SetPaused(event.event, makeFmodBool(paused))
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_Event_GetPaused         (FMOD_EVENT *event, FMOD_BOOL *paused);
func (event *Event) GetPaused() (bool, error) {
  var ferr C.FMOD_RESULT
  var paused C.FMOD_BOOL
  base.Thread(func() {
    ferr = C.FMOD_Event_GetPaused(event.event, &paused)
  })
  return paused != 0, base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_Event_SetMute           (FMOD_EVENT *event, FMOD_BOOL mute);
func (event *Event) SetMute(mute bool) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_Event_SetMute(event.event, makeFmodBool(mute))
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_Event_GetMute           (FMOD_EVENT *event, FMOD_BOOL *mute);
func (event *Event) GetMute() (bool, error) {
  var ferr C.FMOD_RESULT
  var mute C.FMOD_BOOL
  base.Thread(func() {
    ferr = C.FMOD_Event_GetMute(event.event, &mute)
  })
  return mute != 0, base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_Event_Set3DAttributes         (FMOD_EVENT *event, const FMOD_VECTOR *position, const FMOD_VECTOR *velocity, const FMOD_VECTOR *orientation);
func (e *Event) Set3DAttributes() error {
  return errors.New("Event.Set3DAttributes() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_Get3DAttributes         (FMOD_EVENT *event, FMOD_VECTOR *position, FMOD_VECTOR *velocity, FMOD_VECTOR *orientation);
func (e *Event) Get3DAttributes() error {
  return errors.New("Event.Get3DAttributes() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_Set3DOcclusion          (FMOD_EVENT *event, float directocclusion, float reverbocclusion);
func (e *Event) Set3DOcclusion() error {
  return errors.New("Event.Set3DOcclusion() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_Get3DOcclusion          (FMOD_EVENT *event, float *directocclusion, float *reverbocclusion);
func (e *Event) Get3DOcclusion() error {
  return errors.New("Event.Get3DOcclusion() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_SetReverbProperties     (FMOD_EVENT *event, const FMOD_REVERB_CHANNELPROPERTIES *props);
func (e *Event) SetReverbProperties() error {
  return errors.New("Event.SetReverbProperties() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_GetReverbProperties     (FMOD_EVENT *event, FMOD_REVERB_CHANNELPROPERTIES *props);
func (e *Event) GetReverbProperties() error {
  return errors.New("Event.GetReverbProperties() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_SetUserData             (FMOD_EVENT *event, void *userdata);
func (e *Event) SetUserData() error {
  return errors.New("Event.SetUserData() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_GetUserData             (FMOD_EVENT *event, void **userdata);
func (e *Event) GetUserData() error {
  return errors.New("Event.GetUserData() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_GetMemoryInfo           (FMOD_EVENT *event, unsigned int memorybits, unsigned int event_memorybits, unsigned int *memoryused, FMOD_MEMORY_USAGE_DETAILS *memoryused_details);
func (e *Event) GetMemoryInfo() error {
  return errors.New("Event.GetMemoryInfo() has not been implemented yet.")
}
