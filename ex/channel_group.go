package fmod_ex

// #include "fmod.h"
import "C"

import (
  "github.com/runningwild/fmod/base"
)

type ChannelGroup struct {
  group *C.FMOD_CHANNELGROUP
}

// FMOD_RESULT F_API FMOD_ChannelGroup_Release          (FMOD_CHANNELGROUP *channelgroup);
// TODO: Bind this to a finalizer
func (cg *ChannelGroup) Release() error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_ChannelGroup_Release(cg.group)
  })
  cg.group = nil
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_ChannelGroup_GetSystemObject  (FMOD_CHANNELGROUP *channelgroup, FMOD_SYSTEM **system);

// /*
//      Channelgroup scale values.  (changes attributes relative to the channels, doesn't overwrite them)
// */

// FMOD_RESULT F_API FMOD_ChannelGroup_SetVolume        (FMOD_CHANNELGROUP *channelgroup, float volume);
func (cg *ChannelGroup) SetVolume(volume float64) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_ChannelGroup_SetVolume(cg.group, C.float(volume))
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_ChannelGroup_GetVolume        (FMOD_CHANNELGROUP *channelgroup, float *volume);

// FMOD_RESULT F_API FMOD_ChannelGroup_SetPitch         (FMOD_CHANNELGROUP *channelgroup, float pitch);
func (cg *ChannelGroup) SetPitch(pitch float64) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_ChannelGroup_SetPitch(cg.group, C.float(pitch))
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_ChannelGroup_GetPitch         (FMOD_CHANNELGROUP *channelgroup, float *pitch);
// FMOD_RESULT F_API FMOD_ChannelGroup_Set3DOcclusion   (FMOD_CHANNELGROUP *channelgroup, float directocclusion, float reverbocclusion);
// FMOD_RESULT F_API FMOD_ChannelGroup_Get3DOcclusion   (FMOD_CHANNELGROUP *channelgroup, float *directocclusion, float *reverbocclusion);
// FMOD_RESULT F_API FMOD_ChannelGroup_SetPaused        (FMOD_CHANNELGROUP *channelgroup, FMOD_BOOL paused);
// FMOD_RESULT F_API FMOD_ChannelGroup_GetPaused        (FMOD_CHANNELGROUP *channelgroup, FMOD_BOOL *paused);

// FMOD_RESULT F_API FMOD_ChannelGroup_SetMute          (FMOD_CHANNELGROUP *channelgroup, FMOD_BOOL mute);
func (cg *ChannelGroup) SetMute(mute bool) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_ChannelGroup_SetMute(cg.group, makeFmodBool(mute))
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_ChannelGroup_GetMute          (FMOD_CHANNELGROUP *channelgroup, FMOD_BOOL *mute);

// /*
//      Channelgroup override values.  (recursively overwrites whatever settings the channels had)
// */

// FMOD_RESULT F_API FMOD_ChannelGroup_Stop             (FMOD_CHANNELGROUP *channelgroup);
// FMOD_RESULT F_API FMOD_ChannelGroup_OverrideVolume   (FMOD_CHANNELGROUP *channelgroup, float volume);
// FMOD_RESULT F_API FMOD_ChannelGroup_OverrideFrequency(FMOD_CHANNELGROUP *channelgroup, float frequency);
// FMOD_RESULT F_API FMOD_ChannelGroup_OverridePan      (FMOD_CHANNELGROUP *channelgroup, float pan);
// FMOD_RESULT F_API FMOD_ChannelGroup_OverrideReverbProperties(FMOD_CHANNELGROUP *channelgroup, const FMOD_REVERB_CHANNELPROPERTIES *prop);
// FMOD_RESULT F_API FMOD_ChannelGroup_Override3DAttributes(FMOD_CHANNELGROUP *channelgroup, const FMOD_VECTOR *pos, const FMOD_VECTOR *vel);
// FMOD_RESULT F_API FMOD_ChannelGroup_OverrideSpeakerMix(FMOD_CHANNELGROUP *channelgroup, float frontleft, float frontright, float center, float lfe, float backleft, float backright, float sideleft, float sideright);

// /*
//      Nested channel groups.
// */

// FMOD_RESULT F_API FMOD_ChannelGroup_AddGroup         (FMOD_CHANNELGROUP *channelgroup, FMOD_CHANNELGROUP *group);
func (cg *ChannelGroup) AddGroup(group *ChannelGroup) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_ChannelGroup_AddGroup(cg.group, group.group)
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_ChannelGroup_GetNumGroups     (FMOD_CHANNELGROUP *channelgroup, int *numgroups);
// FMOD_RESULT F_API FMOD_ChannelGroup_GetGroup         (FMOD_CHANNELGROUP *channelgroup, int index, FMOD_CHANNELGROUP **group);
// FMOD_RESULT F_API FMOD_ChannelGroup_GetParentGroup   (FMOD_CHANNELGROUP *channelgroup, FMOD_CHANNELGROUP **group);

// /*
//      DSP functionality only for channel groups playing sounds created with FMOD_SOFTWARE.
// */

// FMOD_RESULT F_API FMOD_ChannelGroup_GetDSPHead       (FMOD_CHANNELGROUP *channelgroup, FMOD_DSP **dsp);
func (cg *ChannelGroup) GetDSPHead() (*Dsp, error) {
  var ferr C.FMOD_RESULT
  var dsp Dsp
  base.Thread(func() {
    ferr = C.FMOD_ChannelGroup_GetDSPHead(cg.group, &dsp.dsp)
  })
  err := base.ResultToError(ferr)
  if err != nil {
    return nil, err
  }
  return &dsp, nil
}

// FMOD_RESULT F_API FMOD_ChannelGroup_AddDSP           (FMOD_CHANNELGROUP *channelgroup, FMOD_DSP *dsp, FMOD_DSPCONNECTION **connection);
func (cg *ChannelGroup) AddDSP(dsp *Dsp) (*DspConn, error) {
  var ferr C.FMOD_RESULT
  var conn DspConn
  base.Thread(func() {
    ferr = C.FMOD_ChannelGroup_AddDSP(cg.group, dsp.dsp, &conn.conn)
  })
  err := base.ResultToError(ferr)
  if err != nil {
    return nil, err
  }
  return &conn, nil
}

// /*
//      Information only functions.
// */

// FMOD_RESULT F_API FMOD_ChannelGroup_GetName          (FMOD_CHANNELGROUP *channelgroup, char *name, int namelen);
// FMOD_RESULT F_API FMOD_ChannelGroup_GetNumChannels   (FMOD_CHANNELGROUP *channelgroup, int *numchannels);
// FMOD_RESULT F_API FMOD_ChannelGroup_GetChannel       (FMOD_CHANNELGROUP *channelgroup, int index, FMOD_CHANNEL **channel);
// FMOD_RESULT F_API FMOD_ChannelGroup_GetSpectrum      (FMOD_CHANNELGROUP *channelgroup, float *spectrumarray, int numvalues, int channeloffset, FMOD_DSP_FFT_WINDOW windowtype);
// FMOD_RESULT F_API FMOD_ChannelGroup_GetWaveData      (FMOD_CHANNELGROUP *channelgroup, float *wavearray, int numvalues, int channeloffset);

// /*
//      Userdata set/get.
// */

// FMOD_RESULT F_API FMOD_ChannelGroup_SetUserData      (FMOD_CHANNELGROUP *channelgroup, void *userdata);
// FMOD_RESULT F_API FMOD_ChannelGroup_GetUserData      (FMOD_CHANNELGROUP *channelgroup, void **userdata);

// FMOD_RESULT F_API FMOD_ChannelGroup_GetMemoryInfo    (FMOD_CHANNELGROUP *channelgroup, unsigned int memorybits, unsigned int event_memorybits, unsigned int *memoryused, FMOD_MEMORY_USAGE_DETAILS *memoryused_details);
