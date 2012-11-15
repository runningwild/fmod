package fmod_event

// #include "fmod_event.h"
// #include "stdlib.h"
import "C"

import (
  "unsafe"
  "errors"
  "github.com/MobRulesGames/fmod/base"
)

type MusicSystem struct {
  system *C.FMOD_MUSICSYSTEM
}

// FMOD_RESULT F_API FMOD_MusicSystem_Reset             (FMOD_MUSICSYSTEM *musicsystem);
func (music *MusicSystem) Reset() error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_MusicSystem_Reset(music.system)
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_MusicSystem_SetVolume         (FMOD_MUSICSYSTEM *musicsystem, float volume);
func (music *MusicSystem) SetVolume(volume float64) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_MusicSystem_SetVolume(music.system, C.float(volume))
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_MusicSystem_GetVolume         (FMOD_MUSICSYSTEM *musicsystem, float *volume);
func (music *MusicSystem) GetVolume() (float64, error) {
  var ferr C.FMOD_RESULT
  var volume C.float
  base.Thread(func() {
    ferr = C.FMOD_MusicSystem_GetVolume(music.system, &volume)
  })
  return float64(volume), base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_MusicSystem_SetReverbProperties(FMOD_MUSICSYSTEM *musicsystem, const FMOD_REVERB_CHANNELPROPERTIES *props);
func (music *MusicSystem) SetReverbProperties() error {
  return errors.New("MusicSystem.SetReverbProperties() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_MusicSystem_GetReverbProperties(FMOD_MUSICSYSTEM *musicsystem, FMOD_REVERB_CHANNELPROPERTIES *props);
func (music *MusicSystem) GetReverbProperties() error {
  return errors.New("MusicSystem.GetReverbProperties() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_MusicSystem_SetPaused         (FMOD_MUSICSYSTEM *musicsystem, FMOD_BOOL paused);
func (music *MusicSystem) SetPaused(paused bool) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_MusicSystem_SetPaused(music.system, makeFmodBool(paused))
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_MusicSystem_GetPaused         (FMOD_MUSICSYSTEM *musicsystem, FMOD_BOOL *paused);
func (music *MusicSystem) GetPaused() (bool, error) {
  var ferr C.FMOD_RESULT
  var paused C.FMOD_BOOL
  base.Thread(func() {
    ferr = C.FMOD_MusicSystem_GetPaused(music.system, &paused)
  })
  return paused != 0, base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_MusicSystem_SetMute           (FMOD_MUSICSYSTEM *musicsystem, FMOD_BOOL mute);
func (music *MusicSystem) SetMute(mute bool) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_MusicSystem_SetMute(music.system, makeFmodBool(mute))
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_MusicSystem_GetMute           (FMOD_MUSICSYSTEM *musicsystem, FMOD_BOOL *mute);
func (music *MusicSystem) GetMute() (bool, error) {
  var ferr C.FMOD_RESULT
  var mute C.FMOD_BOOL
  base.Thread(func() {
    ferr = C.FMOD_MusicSystem_GetMute(music.system, &mute)
  })
  return mute != 0, base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_MusicSystem_GetInfo           (FMOD_MUSICSYSTEM *musicsystem, FMOD_MUSIC_INFO *info);
func (music *MusicSystem) GetInfo() error {
  return errors.New("MusicSystem.GetInfo() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_MusicSystem_PromptCue         (FMOD_MUSICSYSTEM *musicsystem, FMOD_MUSIC_CUE_ID id);
func (music *MusicSystem) PromptCue(cue MusicEntity) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_MusicSystem_PromptCue(music.system, C.FMOD_MUSIC_CUE_ID(cue.Id()))
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_MusicSystem_PrepareCue        (FMOD_MUSICSYSTEM *musicsystem, FMOD_MUSIC_CUE_ID id, FMOD_MUSICPROMPT **prompt);
func (music *MusicSystem) PrepareCue(cue MusicEntity) (*Prompt, error) {
  var ferr C.FMOD_RESULT
  var prompt Prompt
  base.Thread(func() {
    ferr = C.FMOD_MusicSystem_PrepareCue(music.system, C.FMOD_MUSIC_CUE_ID(cue.Id()), &prompt.prompt)
  })
  err := base.ResultToError(ferr)
  if err != nil {
    return nil, err
  }
  return &prompt, nil
}

// FMOD_RESULT F_API FMOD_MusicSystem_GetParameterValue (FMOD_MUSICSYSTEM *musicsystem, FMOD_MUSIC_PARAM_ID id, float *parameter);
func (music *MusicSystem) GetParameterValue() error {
  return errors.New("MusicSystem.GetParameterValue() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_MusicSystem_SetParameterValue (FMOD_MUSICSYSTEM *musicsystem, FMOD_MUSIC_PARAM_ID id, float parameter);
func (music *MusicSystem) SetParameterValue() error {
  return errors.New("MusicSystem.SetParameterValue() has not been implemented yet.")
}

type MusicId C.FMOD_MUSIC_ID

type MusicEntity struct {
  name string
  id   MusicId
}

func (e *MusicEntity) Name() string {
  return e.name
}
func (e *MusicEntity) Id() MusicId {
  return e.id
}

// Different than the FMOD definition of GetCues - instead this returns a
// slice of all matching cues.  Pass in the empty string to match all cues.
func (music *MusicSystem) GetCues(filter string) ([]MusicEntity, error) {
  var it C.FMOD_MUSIC_ITERATOR
  var cfilter *C.char
  if filter != "" {
    cfilter = C.CString(filter)
  } else {
    cfilter = nil
  }

  var ents []MusicEntity
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_MusicSystem_GetCues(music.system, &it, cfilter)
    if base.ResultToError(ferr) != nil {
      return
    }
    for unsafe.Pointer(it.value) != unsafe.Pointer(uintptr(0)) {
      ent := MusicEntity{C.GoString(it.value.name), MusicId(it.value.id)}
      ents = append(ents, ent)
      ferr = C.FMOD_MusicSystem_GetNextCue(music.system, &it)
      if base.ResultToError(ferr) != nil {
        return
      }
    }
  })
  if filter != "" {
    C.free(unsafe.Pointer(cfilter))
  }
  return ents, base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_MusicSystem_GetParameters     (FMOD_MUSICSYSTEM *musicsystem, FMOD_MUSIC_ITERATOR *it, const char *filter);
func (music *MusicSystem) GetParameters() error {
  return errors.New("MusicSystem.GetParameters() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_MusicSystem_GetNextParameter  (FMOD_MUSICSYSTEM *musicsystem, FMOD_MUSIC_ITERATOR *it);
func (music *MusicSystem) GetNextParameter() error {
  return errors.New("MusicSystem.GetNextParameter() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_MusicSystem_LoadSoundData     (FMOD_MUSICSYSTEM *musicsystem, FMOD_EVENT_RESOURCE resource, FMOD_EVENT_MODE mode);
func (music *MusicSystem) LoadSoundData(resource Resource, mode Mode) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_MusicSystem_LoadSoundData(music.system, C.FMOD_EVENT_RESOURCE(resource), C.FMOD_EVENT_MODE(mode))
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_MusicSystem_FreeSoundData     (FMOD_MUSICSYSTEM *musicsystem, FMOD_BOOL waituntilready);
func (music *MusicSystem) FreeSoundData() error {
  return errors.New("MusicSystem.FreeSoundData() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_MusicSystem_SetCallback       (FMOD_MUSICSYSTEM *musicsystem, FMOD_MUSIC_CALLBACK callback, void *userdata);
func (music *MusicSystem) SetCallback() error {
  return errors.New("MusicSystem.SetCallback() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_MusicSystem_GetMemoryInfo     (FMOD_MUSICSYSTEM *musicsystem, unsigned int memorybits, unsigned int event_memorybits, unsigned int *memoryused, FMOD_MEMORY_USAGE_DETAILS *memoryused_details);
func (music *MusicSystem) GetMemoryInfo() error {
  return errors.New("MusicSystem.GetMemoryInfo() has not been implemented yet.")
}

type Prompt struct {
  prompt *C.FMOD_MUSICPROMPT
}

// FMOD_RESULT F_API FMOD_MusicPrompt_Release           (FMOD_MUSICPROMPT *musicprompt);
func (p *Prompt) Release() error {
  return errors.New("Prompt.Release() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_MusicPrompt_Begin             (FMOD_MUSICPROMPT *musicprompt);
func (p *Prompt) Begin() error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_MusicPrompt_Begin(p.prompt)
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_MusicPrompt_End               (FMOD_MUSICPROMPT *musicprompt);
func (p *Prompt) End() error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_MusicPrompt_End(p.prompt)
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_MusicPrompt_IsActive          (FMOD_MUSICPROMPT *musicprompt, FMOD_BOOL *active);
func (p *Prompt) IsActive() (bool, error) {
  var ferr C.FMOD_RESULT
  var active C.FMOD_BOOL
  base.Thread(func() {
    ferr = C.FMOD_MusicPrompt_IsActive(p.prompt, &active)
  })
  return active != 0, base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_MusicPrompt_GetMemoryInfo     (FMOD_MUSICPROMPT *musicprompt, unsigned int memorybits, unsigned int event_memorybits, unsigned int *memoryused, FMOD_MEMORY_USAGE_DETAILS *memoryused_details);
func (p *Prompt) GetMemoryInfo() error {
  return errors.New("Prompt.GetMemoryInfo() has not been implemented yet.")
}
