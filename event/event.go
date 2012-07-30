package fmod_event

// #include "fmod_event.h"
// #include "stdlib.h"
import "C"

import (
  "errors"
  "github.com/runningwild/fmod/base"
)

type Event struct {
  event *C.FMOD_EVENT
}

// FMOD_RESULT F_API FMOD_Event_Release                 (FMOD_EVENT *event, FMOD_BOOL freeeventdata, FMOD_BOOL waituntilready);
func (e *Event) Release(free_event_data, wait_until_ready bool) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_Event_Release(e.event, makeFmodBool(free_event_data), makeFmodBool(wait_until_ready))
  })
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
func (e *Event) GetInfo() error {
  return errors.New("Event.GetInfo() has not been implemented yet.")
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
func (e *Event) GetParameter() error {
  return errors.New("Event.GetParameter() has not been implemented yet.")
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

// FMOD_RESULT F_API FMOD_Event_SetVolume               (FMOD_EVENT *event, float volume);
func (e *Event) SetVolume() error {
  return errors.New("Event.SetVolume() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_GetVolume               (FMOD_EVENT *event, float *volume);
func (e *Event) GetVolume() error {
  return errors.New("Event.GetVolume() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_SetPitch                (FMOD_EVENT *event, float pitch, FMOD_EVENT_PITCHUNITS units);
func (e *Event) SetPitch() error {
  return errors.New("Event.SetPitch() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_GetPitch                (FMOD_EVENT *event, float *pitch, FMOD_EVENT_PITCHUNITS units);
func (e *Event) GetPitch() error {
  return errors.New("Event.GetPitch() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_SetPaused               (FMOD_EVENT *event, FMOD_BOOL paused);
func (e *Event) SetPaused() error {
  return errors.New("Event.SetPaused() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_GetPaused               (FMOD_EVENT *event, FMOD_BOOL *paused);
func (e *Event) GetPaused() error {
  return errors.New("Event.GetPaused() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_SetMute                 (FMOD_EVENT *event, FMOD_BOOL mute);
func (e *Event) SetMute() error {
  return errors.New("Event.SetMute() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_Event_GetMute                 (FMOD_EVENT *event, FMOD_BOOL *mute);
func (e *Event) GetMute() error {
  return errors.New("Event.GetMute() has not been implemented yet.")
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
