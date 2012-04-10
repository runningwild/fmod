package fmod

// #include "fmod.h"
import "C"

type Channel struct {
  channel *C.FMOD_CHANNEL
}

//     'Channel' API
// FMOD_RESULT F_API FMOD_Channel_GetSystemObject       (FMOD_CHANNEL *channel, FMOD_SYSTEM **system);

// FMOD_RESULT F_API FMOD_Channel_Stop                  (FMOD_CHANNEL *channel);

// FMOD_RESULT F_API FMOD_Channel_SetPaused             (FMOD_CHANNEL *channel, FMOD_BOOL paused);
func (c *Channel) SetPaused(paused bool) error {
  var ferr C.FMOD_RESULT
  thread(func() {
    ferr = C.FMOD_Channel_SetPaused(c.channel, makeFmodBool(paused))
  })
  return error_map[ferr]
}

// FMOD_RESULT F_API FMOD_Channel_GetPaused             (FMOD_CHANNEL *channel, FMOD_BOOL *paused);
// FMOD_RESULT F_API FMOD_Channel_SetVolume             (FMOD_CHANNEL *channel, float volume);
// FMOD_RESULT F_API FMOD_Channel_GetVolume             (FMOD_CHANNEL *channel, float *volume);
// FMOD_RESULT F_API FMOD_Channel_SetFrequency          (FMOD_CHANNEL *channel, float frequency);
// FMOD_RESULT F_API FMOD_Channel_GetFrequency          (FMOD_CHANNEL *channel, float *frequency);
// FMOD_RESULT F_API FMOD_Channel_SetPan                (FMOD_CHANNEL *channel, float pan);
// FMOD_RESULT F_API FMOD_Channel_GetPan                (FMOD_CHANNEL *channel, float *pan);
// FMOD_RESULT F_API FMOD_Channel_SetDelay              (FMOD_CHANNEL *channel, FMOD_DELAYTYPE delaytype, unsigned int delayhi, unsigned int delaylo);
// FMOD_RESULT F_API FMOD_Channel_GetDelay              (FMOD_CHANNEL *channel, FMOD_DELAYTYPE delaytype, unsigned int *delayhi, unsigned int *delaylo);
// FMOD_RESULT F_API FMOD_Channel_SetSpeakerMix         (FMOD_CHANNEL *channel, float frontleft, float frontright, float center, float lfe, float backleft, float backright, float sideleft, float sideright);
// FMOD_RESULT F_API FMOD_Channel_GetSpeakerMix         (FMOD_CHANNEL *channel, float *frontleft, float *frontright, float *center, float *lfe, float *backleft, float *backright, float *sideleft, float *sideright);
// FMOD_RESULT F_API FMOD_Channel_SetSpeakerLevels      (FMOD_CHANNEL *channel, FMOD_SPEAKER speaker, float *levels, int numlevels);
// FMOD_RESULT F_API FMOD_Channel_GetSpeakerLevels      (FMOD_CHANNEL *channel, FMOD_SPEAKER speaker, float *levels, int numlevels);
// FMOD_RESULT F_API FMOD_Channel_SetInputChannelMix    (FMOD_CHANNEL *channel, float *levels, int numlevels);
// FMOD_RESULT F_API FMOD_Channel_GetInputChannelMix    (FMOD_CHANNEL *channel, float *levels, int numlevels);
// FMOD_RESULT F_API FMOD_Channel_SetMute               (FMOD_CHANNEL *channel, FMOD_BOOL mute);
// FMOD_RESULT F_API FMOD_Channel_GetMute               (FMOD_CHANNEL *channel, FMOD_BOOL *mute);
// FMOD_RESULT F_API FMOD_Channel_SetPriority           (FMOD_CHANNEL *channel, int priority);
// FMOD_RESULT F_API FMOD_Channel_GetPriority           (FMOD_CHANNEL *channel, int *priority);
// FMOD_RESULT F_API FMOD_Channel_SetPosition           (FMOD_CHANNEL *channel, unsigned int position, FMOD_TIMEUNIT postype);
// FMOD_RESULT F_API FMOD_Channel_GetPosition           (FMOD_CHANNEL *channel, unsigned int *position, FMOD_TIMEUNIT postype);
// FMOD_RESULT F_API FMOD_Channel_SetReverbProperties   (FMOD_CHANNEL *channel, const FMOD_REVERB_CHANNELPROPERTIES *prop);
// FMOD_RESULT F_API FMOD_Channel_GetReverbProperties   (FMOD_CHANNEL *channel, FMOD_REVERB_CHANNELPROPERTIES *prop);
// FMOD_RESULT F_API FMOD_Channel_SetLowPassGain        (FMOD_CHANNEL *channel, float gain);
// FMOD_RESULT F_API FMOD_Channel_GetLowPassGain        (FMOD_CHANNEL *channel, float *gain);

// FMOD_RESULT F_API FMOD_Channel_SetChannelGroup       (FMOD_CHANNEL *channel, FMOD_CHANNELGROUP *channelgroup);
func (c *Channel) SetChannelGroup(group *ChannelGroup) error {
  var ferr C.FMOD_RESULT
  thread(func() {
    ferr = C.FMOD_Channel_SetChannelGroup(c.channel, group.group)
  })
  return error_map[ferr]
}

// FMOD_RESULT F_API FMOD_Channel_GetChannelGroup       (FMOD_CHANNEL *channel, FMOD_CHANNELGROUP **channelgroup);
// FMOD_RESULT F_API FMOD_Channel_SetCallback           (FMOD_CHANNEL *channel, FMOD_CHANNEL_CALLBACK callback);

// /*
//      3D functionality.
// */

// FMOD_RESULT F_API FMOD_Channel_Set3DAttributes       (FMOD_CHANNEL *channel, const FMOD_VECTOR *pos, const FMOD_VECTOR *vel);
// FMOD_RESULT F_API FMOD_Channel_Get3DAttributes       (FMOD_CHANNEL *channel, FMOD_VECTOR *pos, FMOD_VECTOR *vel);
// FMOD_RESULT F_API FMOD_Channel_Set3DMinMaxDistance   (FMOD_CHANNEL *channel, float mindistance, float maxdistance);
// FMOD_RESULT F_API FMOD_Channel_Get3DMinMaxDistance   (FMOD_CHANNEL *channel, float *mindistance, float *maxdistance);
// FMOD_RESULT F_API FMOD_Channel_Set3DConeSettings     (FMOD_CHANNEL *channel, float insideconeangle, float outsideconeangle, float outsidevolume);
// FMOD_RESULT F_API FMOD_Channel_Get3DConeSettings     (FMOD_CHANNEL *channel, float *insideconeangle, float *outsideconeangle, float *outsidevolume);
// FMOD_RESULT F_API FMOD_Channel_Set3DConeOrientation  (FMOD_CHANNEL *channel, FMOD_VECTOR *orientation);
// FMOD_RESULT F_API FMOD_Channel_Get3DConeOrientation  (FMOD_CHANNEL *channel, FMOD_VECTOR *orientation);
// FMOD_RESULT F_API FMOD_Channel_Set3DCustomRolloff    (FMOD_CHANNEL *channel, FMOD_VECTOR *points, int numpoints);
// FMOD_RESULT F_API FMOD_Channel_Get3DCustomRolloff    (FMOD_CHANNEL *channel, FMOD_VECTOR **points, int *numpoints);
// FMOD_RESULT F_API FMOD_Channel_Set3DOcclusion        (FMOD_CHANNEL *channel, float directocclusion, float reverbocclusion);
// FMOD_RESULT F_API FMOD_Channel_Get3DOcclusion        (FMOD_CHANNEL *channel, float *directocclusion, float *reverbocclusion);
// FMOD_RESULT F_API FMOD_Channel_Set3DSpread           (FMOD_CHANNEL *channel, float angle);
// FMOD_RESULT F_API FMOD_Channel_Get3DSpread           (FMOD_CHANNEL *channel, float *angle);
// FMOD_RESULT F_API FMOD_Channel_Set3DPanLevel         (FMOD_CHANNEL *channel, float level);
// FMOD_RESULT F_API FMOD_Channel_Get3DPanLevel         (FMOD_CHANNEL *channel, float *level);
// FMOD_RESULT F_API FMOD_Channel_Set3DDopplerLevel     (FMOD_CHANNEL *channel, float level);
// FMOD_RESULT F_API FMOD_Channel_Get3DDopplerLevel     (FMOD_CHANNEL *channel, float *level);
// FMOD_RESULT F_API FMOD_Channel_Set3DDistanceFilter   (FMOD_CHANNEL *channel, FMOD_BOOL custom, float customLevel, float centerFreq);
// FMOD_RESULT F_API FMOD_Channel_Get3DDistanceFilter   (FMOD_CHANNEL *channel, FMOD_BOOL *custom, float *customLevel, float *centerFreq);

// /*
//      DSP functionality only for channels playing sounds created with FMOD_SOFTWARE.
// */

// FMOD_RESULT F_API FMOD_Channel_GetDSPHead            (FMOD_CHANNEL *channel, FMOD_DSP **dsp);
// FMOD_RESULT F_API FMOD_Channel_AddDSP                (FMOD_CHANNEL *channel, FMOD_DSP *dsp, FMOD_DSPCONNECTION **connection);

// /*
//      Information only functions.
// */

// FMOD_RESULT F_API FMOD_Channel_IsPlaying             (FMOD_CHANNEL *channel, FMOD_BOOL *isplaying);
// FMOD_RESULT F_API FMOD_Channel_IsVirtual             (FMOD_CHANNEL *channel, FMOD_BOOL *isvirtual);
// FMOD_RESULT F_API FMOD_Channel_GetAudibility         (FMOD_CHANNEL *channel, float *audibility);
// FMOD_RESULT F_API FMOD_Channel_GetCurrentSound       (FMOD_CHANNEL *channel, FMOD_SOUND **sound);
// FMOD_RESULT F_API FMOD_Channel_GetSpectrum           (FMOD_CHANNEL *channel, float *spectrumarray, int numvalues, int channeloffset, FMOD_DSP_FFT_WINDOW windowtype);
// FMOD_RESULT F_API FMOD_Channel_GetWaveData           (FMOD_CHANNEL *channel, float *wavearray, int numvalues, int channeloffset);
// FMOD_RESULT F_API FMOD_Channel_GetIndex              (FMOD_CHANNEL *channel, int *index);

// /*
//      Functions also found in Sound class but here they can be set per channel.
// */

// FMOD_RESULT F_API FMOD_Channel_SetMode               (FMOD_CHANNEL *channel, FMOD_MODE mode);
// FMOD_RESULT F_API FMOD_Channel_GetMode               (FMOD_CHANNEL *channel, FMOD_MODE *mode);
// FMOD_RESULT F_API FMOD_Channel_SetLoopCount          (FMOD_CHANNEL *channel, int loopcount);
// FMOD_RESULT F_API FMOD_Channel_GetLoopCount          (FMOD_CHANNEL *channel, int *loopcount);
// FMOD_RESULT F_API FMOD_Channel_SetLoopPoints         (FMOD_CHANNEL *channel, unsigned int loopstart, FMOD_TIMEUNIT loopstarttype, unsigned int loopend, FMOD_TIMEUNIT loopendtype);
// FMOD_RESULT F_API FMOD_Channel_GetLoopPoints         (FMOD_CHANNEL *channel, unsigned int *loopstart, FMOD_TIMEUNIT loopstarttype, unsigned int *loopend, FMOD_TIMEUNIT loopendtype);

// /*
//      Userdata set/get.                                                
// */

// FMOD_RESULT F_API FMOD_Channel_SetUserData           (FMOD_CHANNEL *channel, void *userdata);
// FMOD_RESULT F_API FMOD_Channel_GetUserData           (FMOD_CHANNEL *channel, void **userdata);

// FMOD_RESULT F_API FMOD_Channel_GetMemoryInfo         (FMOD_CHANNEL *channel, unsigned int memorybits, unsigned int event_memorybits, unsigned int *memoryused, FMOD_MEMORY_USAGE_DETAILS *memoryused_details);
