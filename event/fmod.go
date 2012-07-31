package fmod_event

// #cgo darwin  LDFLAGS: -Llib -lfmodevent
// #cgo CFLAGS: -Iinc -I../ex/inc
// #include "fmod_event.h"
// #include "stdlib.h"
import "C"

import (
  "errors"
  "unsafe"
  "github.com/runningwild/fmod/base"
)

func makeFmodBool(b bool) C.FMOD_BOOL {
  if b {
    return 1
  }
  return 0
}

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
func (sys *EventSystem) Init(max_channels int, init_flags InitFlags, extradriverdata interface{}, event_init_flags EventInitFlags) error {
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

/*
   FEV load/unload.                                 
*/

type LoadInfo struct {
  // Optional. Key, or 'password' to decrypt a bank.  A sound designer may
  // have encrypted the audio data to protect their sound data from 'rippers'.
  Encryption_key string

  // Optional. A value between 0 -> 1 that is multiplied with the number of
  // sound definition entries in each sound definition in the project being
  // loaded in order to programmatically reduce the number of sound definition
  // entries used at runtime.
  Sound_def_entry_limit float64

  // Optional. Length of memory buffer pointed to by name_or_data parameter
  // passed to EventSystem.Load(). If this field is non-zero then the
  // name_or_data parameter passed to EventSystem.Load() will be interpreted
  // as a pointer to a memory buffer containing the .fev data to load.
  // If this field is zero the name_or_data parameter is interpreted as the
  // filename of the .fev file to load.
  Load_from_memory_length uint32

  // Optional. If this member is set to true, newly-loaded categories will
  // impart their properties (volume, pitch etc.) to existing categories of
  // the same name.
  Overrid_category_vals bool

  // Optional. Specify 0 to ignore. If this value is non-zero, FMOD will
  // create an instance pool for simple events with
  // "sizeof_instancepool_simple" entries.
  // Note: Event instance pools currently work for simple events only.
  // Complex events will behave as normal and not be pooled.
  Sizeof_instancepool_simple uint32
}

func (info *LoadInfo) cIfy() *C.FMOD_EVENT_LOADINFO {
  var cinfo C.FMOD_EVENT_LOADINFO
  cinfo.size = C.uint(unsafe.Sizeof(&cinfo))
  cinfo.encryptionkey = C.CString(info.Encryption_key)
  cinfo.sounddefentrylimit = C.float(info.Sound_def_entry_limit)
  cinfo.loadfrommemory_length = C.uint(info.Load_from_memory_length)
  cinfo.override_category_vals = makeFmodBool(info.Overrid_category_vals)
  cinfo.sizeof_instancepool_simple = C.uint(info.Sizeof_instancepool_simple)
  return &cinfo
}

// FMOD_RESULT F_API FMOD_EventSystem_Load              (FMOD_EVENTSYSTEM *eventsystem, const char *name_or_data, FMOD_EVENT_LOADINFO *loadinfo, FMOD_EVENTPROJECT **project);
func (sys *EventSystem) LoadPath(path string, info *LoadInfo) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    cpath := C.CString(path)
    var cinfo *C.FMOD_EVENT_LOADINFO
    if info != nil {
      cinfo = info.cIfy()
    }
    ferr = C.FMOD_EventSystem_Load(sys.system, cpath, cinfo, (**C.FMOD_EVENTPROJECT)(unsafe.Pointer(uintptr(0))))
    if info != nil {
      if cinfo.encryptionkey != nil {
        C.free(unsafe.Pointer(cinfo.encryptionkey))
      }
    }
    C.free(unsafe.Pointer(cpath))
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_EventSystem_Unload            (FMOD_EVENTSYSTEM *eventsystem);

/*
   Event,EventGroup,EventCategory Retrieval.        
*/

// FMOD_RESULT F_API FMOD_EventSystem_GetProject        (FMOD_EVENTSYSTEM *eventsystem, const char *name, FMOD_EVENTPROJECT **project);
// FMOD_RESULT F_API FMOD_EventSystem_GetProjectByIndex (FMOD_EVENTSYSTEM *eventsystem, int index, FMOD_EVENTPROJECT **project);
// FMOD_RESULT F_API FMOD_EventSystem_GetNumProjects    (FMOD_EVENTSYSTEM *eventsystem, int *numprojects);
// FMOD_RESULT F_API FMOD_EventSystem_GetCategory       (FMOD_EVENTSYSTEM *eventsystem, const char *name, FMOD_EVENTCATEGORY **category);
// FMOD_RESULT F_API FMOD_EventSystem_GetCategoryByIndex(FMOD_EVENTSYSTEM *eventsystem, int index, FMOD_EVENTCATEGORY **category);
// FMOD_RESULT F_API FMOD_EventSystem_GetMusicCategory  (FMOD_EVENTSYSTEM *eventsystem, FMOD_EVENTCATEGORY **category);
// FMOD_RESULT F_API FMOD_EventSystem_GetNumCategories  (FMOD_EVENTSYSTEM *eventsystem, int *numcategories);

// FMOD_RESULT F_API FMOD_EventSystem_GetGroup          (FMOD_EVENTSYSTEM *eventsystem, const char *name, FMOD_BOOL cacheevents, FMOD_EVENTGROUP **group);
func (sys *EventSystem) GetGroup(name string, cache_events bool) (*Group, error) {
  var ferr C.FMOD_RESULT
  var group Group
  base.Thread(func() {
    cname := C.CString(name)
    ferr = C.FMOD_EventSystem_GetGroup(sys.system, cname, makeFmodBool(cache_events), &group.group)
    C.free(unsafe.Pointer(cname))
  })
  err := base.ResultToError(ferr)
  if err != nil {
    return nil, err
  }
  return &group, nil
}

// FMOD_RESULT F_API FMOD_EventSystem_GetEvent          (FMOD_EVENTSYSTEM *eventsystem, const char *name, FMOD_EVENT_MODE mode, FMOD_EVENT **event);
func (sys *EventSystem) GetEvent(name string, mode Mode) (*Event, error) {
  var ferr C.FMOD_RESULT
  var event Event
  base.Thread(func() {
    cname := C.CString(name)
    ferr = C.FMOD_EventSystem_GetEvent(sys.system, cname, C.FMOD_EVENT_MODE(mode), &event.event)
    C.free(unsafe.Pointer(cname))
  })
  err := base.ResultToError(ferr)
  if err != nil {
    return nil, err
  }
  return &event, nil
}

// FMOD_RESULT F_API FMOD_EventSystem_GetEventBySystemID(FMOD_EVENTSYSTEM *eventsystem, unsigned int systemid, FMOD_EVENT_MODE mode, FMOD_EVENT **event);
func (sys *EventSystem) GetEventBySystemID(system_id int, mode Mode) (*Event, error) {
  var ferr C.FMOD_RESULT
  var event Event
  base.Thread(func() {
    ferr = C.FMOD_EventSystem_GetEventBySystemID(sys.system, C.uint(system_id), C.FMOD_EVENT_MODE(mode), &event.event)
  })
  err := base.ResultToError(ferr)
  if err != nil {
    return nil, err
  }
  return &event, nil
}

// FMOD_RESULT F_API FMOD_EventSystem_GetEventByGUID    (FMOD_EVENTSYSTEM *eventsystem, const FMOD_GUID *guid, FMOD_EVENT_MODE mode, FMOD_EVENT **event);
// FMOD_RESULT F_API FMOD_EventSystem_GetEventByGUIDString(FMOD_EVENTSYSTEM *eventsystem, const char *guid, FMOD_EVENT_MODE mode, FMOD_EVENT **event);

// FMOD_RESULT F_API FMOD_EventSystem_GetNumEvents      (FMOD_EVENTSYSTEM *eventsystem, int *numevents);
func (sys *EventSystem) GetNumEvents() (int, error) {
  var ferr C.FMOD_RESULT
  var num_events C.int
  base.Thread(func() {
    ferr = C.FMOD_EventSystem_GetNumEvents(sys.system, &num_events)
  })
  return int(num_events), base.ResultToError(ferr)
}

/*
   Reverb interfaces.
*/

// FMOD_RESULT F_API FMOD_EventSystem_SetReverbProperties(FMOD_EVENTSYSTEM *eventsystem, const FMOD_REVERB_PROPERTIES *props);
// FMOD_RESULT F_API FMOD_EventSystem_GetReverbProperties(FMOD_EVENTSYSTEM *eventsystem, FMOD_REVERB_PROPERTIES *props);

// FMOD_RESULT F_API FMOD_EventSystem_GetReverbPreset   (FMOD_EVENTSYSTEM *eventsystem, const char *name, FMOD_REVERB_PROPERTIES *props, int *index);
// FMOD_RESULT F_API FMOD_EventSystem_GetReverbPresetByIndex(FMOD_EVENTSYSTEM *eventsystem, const int index, FMOD_REVERB_PROPERTIES *props, char **name);
// FMOD_RESULT F_API FMOD_EventSystem_GetNumReverbPresets(FMOD_EVENTSYSTEM *eventsystem, int *numpresets);

// FMOD_RESULT F_API FMOD_EventSystem_CreateReverb      (FMOD_EVENTSYSTEM *eventsystem, FMOD_EVENTREVERB **reverb);
// FMOD_RESULT F_API FMOD_EventSystem_SetReverbAmbientProperties(FMOD_EVENTSYSTEM *eventsystem, FMOD_REVERB_PROPERTIES *props);
// FMOD_RESULT F_API FMOD_EventSystem_GetReverbAmbientProperties(FMOD_EVENTSYSTEM *eventsystem, FMOD_REVERB_PROPERTIES *props);

/*
   Event queue interfaces.
*/

// FMOD_RESULT F_API FMOD_EventSystem_CreateEventQueue  (FMOD_EVENTSYSTEM *eventsystem, FMOD_EVENTQUEUE **queue);
// FMOD_RESULT F_API FMOD_EventSystem_CreateEventQueueEntry(FMOD_EVENTSYSTEM *eventsystem, FMOD_EVENT *event, FMOD_EVENTQUEUEENTRY **entry);

/*
   3D Listener interface.
*/

// FMOD_RESULT F_API FMOD_EventSystem_Set3DNumListeners (FMOD_EVENTSYSTEM *eventsystem, int numlisteners);
// FMOD_RESULT F_API FMOD_EventSystem_Get3DNumListeners (FMOD_EVENTSYSTEM *eventsystem, int *numlisteners);
// FMOD_RESULT F_API FMOD_EventSystem_Set3DListenerAttributes(FMOD_EVENTSYSTEM *eventsystem, int listener, const FMOD_VECTOR *pos, const FMOD_VECTOR *vel, const FMOD_VECTOR *forward, const FMOD_VECTOR *up);
// FMOD_RESULT F_API FMOD_EventSystem_Get3DListenerAttributes(FMOD_EVENTSYSTEM *eventsystem, int listener, FMOD_VECTOR *pos, FMOD_VECTOR *vel, FMOD_VECTOR *forward, FMOD_VECTOR *up);

/*
   Get/set user data
*/

// FMOD_RESULT F_API FMOD_EventSystem_SetUserData       (FMOD_EVENTSYSTEM *eventsystem, void *userdata);
// FMOD_RESULT F_API FMOD_EventSystem_GetUserData       (FMOD_EVENTSYSTEM *eventsystem, void **userdata);

/*
   Pre-loading FSB files (from disk or from memory, use FMOD_OPENMEMORY_POINT to point to pre-loaded memory).
*/

// FMOD_RESULT F_API FMOD_EventSystem_PreloadFSB        (FMOD_EVENTSYSTEM *eventsystem, const char *filename, int streaminstance, FMOD_SOUND *sound, FMOD_BOOL unloadprevious);
// FMOD_RESULT F_API FMOD_EventSystem_UnloadFSB         (FMOD_EVENTSYSTEM *eventsystem, const char *filename, int streaminstance);

// FMOD_RESULT F_API FMOD_EventSystem_GetMemoryInfo     (FMOD_EVENTSYSTEM *eventsystem, unsigned int memorybits, unsigned int event_memorybits, unsigned int *memoryused, FMOD_MEMORY_USAGE_DETAILS *memoryused_details);
