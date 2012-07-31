package fmod_event

// #include "fmod_event.h"
// #include "stdlib.h"
import "C"

import (
  "errors"
  "unsafe"
  "runtime"
  "github.com/runningwild/fmod/base"
)

type Group struct {
  group *C.FMOD_EVENTGROUP
}

// FMOD_RESULT F_API FMOD_EventGroup_GetInfo            (FMOD_EVENTGROUP *eventgroup, int *index, char **name);
func (g *Group) GetInfo() error {
  return errors.New("Group.GetInfo() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventGroup_LoadEventData      (FMOD_EVENTGROUP *eventgroup, FMOD_EVENT_RESOURCE resource, FMOD_EVENT_MODE mode);
func (g *Group) LoadEventData(resource Resource, mode Mode) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_EventGroup_LoadEventData(g.group, C.FMOD_EVENT_RESOURCE(resource), C.FMOD_EVENT_MODE(mode))
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_EventGroup_FreeEventData      (FMOD_EVENTGROUP *eventgroup, FMOD_EVENT *event, FMOD_BOOL waituntilready);
func (g *Group) FreeEventData(event *Event, wait_until_ready bool) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_EventGroup_FreeEventData(g.group, event.event, makeFmodBool(wait_until_ready))
  })
  runtime.SetFinalizer(event, nil)
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_EventGroup_GetGroup           (FMOD_EVENTGROUP *eventgroup, const char *name, FMOD_BOOL cacheevents, FMOD_EVENTGROUP **group);
func (g *Group) GetGroup() error {
  return errors.New("Group.GetGroup() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventGroup_GetGroupByIndex    (FMOD_EVENTGROUP *eventgroup, int index, FMOD_BOOL cacheevents, FMOD_EVENTGROUP **group);
func (g *Group) GetGroupByIndex() error {
  return errors.New("Group.GetGroupByIndex() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventGroup_GetParentGroup     (FMOD_EVENTGROUP *eventgroup, FMOD_EVENTGROUP **group);
func (g *Group) GetParentGroup() error {
  return errors.New("Group.GetParentGroup() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventGroup_GetParentProject   (FMOD_EVENTGROUP *eventgroup, FMOD_EVENTPROJECT **project);
func (g *Group) GetParentProject() error {
  return errors.New("Group.GetParentProject() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventGroup_GetNumGroups       (FMOD_EVENTGROUP *eventgroup, int *numgroups);
func (g *Group) GetNumGroups() error {
  return errors.New("Group.GetNumGroups() has not been implemented yet.")
}

func finalizeGroupEvent(event *Event) {
  event.group.FreeEventData(event, true)
}

// FMOD_RESULT F_API FMOD_EventGroup_GetEvent           (FMOD_EVENTGROUP *eventgroup, const char *name, FMOD_EVENT_MODE mode, FMOD_EVENT **event);
func (g *Group) GetEvent(name string, mode Mode) (*Event, error) {
  var event Event
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    cname := C.CString(name)
    ferr = C.FMOD_EventGroup_GetEvent(g.group, cname, C.FMOD_EVENT_MODE(mode), &event.event)
    C.free(unsafe.Pointer(cname))
  })
  err := base.ResultToError(ferr)
  if err != nil {
    return nil, err
  }
  event.group = g
  runtime.SetFinalizer(&event, finalizeGroupEvent)
  return &event, nil
}

// FMOD_RESULT F_API FMOD_EventGroup_GetEventByIndex    (FMOD_EVENTGROUP *eventgroup, int index, FMOD_EVENT_MODE mode, FMOD_EVENT **event);
func (g *Group) GetEventByIndex(index int, mode Mode) (*Event, error) {
  var event Event
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_EventGroup_GetEventByIndex(g.group, C.int(index), C.FMOD_EVENT_MODE(mode), &event.event)
  })
  err := base.ResultToError(ferr)
  if err != nil {
    return nil, err
  }
  event.group = g
  runtime.SetFinalizer(&event, finalizeGroupEvent)
  return &event, nil
}

// FMOD_RESULT F_API FMOD_EventGroup_GetNumEvents       (FMOD_EVENTGROUP *eventgroup, int *numevents);
func (g *Group) GetNumEvents() (int, error) {
  var ferr C.FMOD_RESULT
  var num C.int
  base.Thread(func() {
    ferr = C.FMOD_EventGroup_GetNumEvents(g.group, &num)
  })
  return int(num), base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_EventGroup_GetProperty        (FMOD_EVENTGROUP *eventgroup, const char *propertyname, void *value);
func (g *Group) GetProperty() error {
  return errors.New("Group.GetProperty() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventGroup_GetPropertyByIndex (FMOD_EVENTGROUP *eventgroup, int propertyindex, void *value);
func (g *Group) GetPropertyByIndex() error {
  return errors.New("Group.GetPropertyByIndex() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventGroup_GetNumProperties   (FMOD_EVENTGROUP *eventgroup, int *numproperties);
func (g *Group) GetNumProperties() error {
  return errors.New("Group.GetNumProperties() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventGroup_GetState           (FMOD_EVENTGROUP *eventgroup, FMOD_EVENT_STATE *state);
func (g *Group) GetState() error {
  return errors.New("Group.GetState() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventGroup_SetUserData        (FMOD_EVENTGROUP *eventgroup, void *userdata);
func (g *Group) SetUserData() error {
  return errors.New("Group.SetUserData() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventGroup_GetUserData        (FMOD_EVENTGROUP *eventgroup, void **userdata);
func (g *Group) GetUserData() error {
  return errors.New("Group.GetUserData() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventGroup_GetMemoryInfo      (FMOD_EVENTGROUP *eventgroup, unsigned int memorybits, unsigned int event_memorybits, unsigned int *memoryused, FMOD_MEMORY_USAGE_DETAILS *memoryused_details);
func (g *Group) GetMemoryInfo() error {
  return errors.New("Group.GetMemoryInfo() has not been implemented yet.")
}
