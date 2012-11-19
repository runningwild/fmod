package fmod_ex

// #cgo darwin  LDFLAGS: -L../lib/darwin -lfmodevent -lfmodex
// #cgo CFLAGS: -Iinc
// #include "fmod.h"
// #include "fmod_errors.h"
// #include "__wincompat.h"
import "C"

import (
  "github.com/MobRulesGames/fmod/base"
  "unsafe"
)

const null = unsafe.Pointer(uintptr(0))

func makeNullTerminatedBytes(s string) []byte {
  b := make([]byte, len(s)+1)
  copy(b, []byte(s))
  return b
}
func makeFmodBool(b bool) C.FMOD_BOOL {
  if b {
    return 1
  }
  return 0
}

type System struct {
  system *C.FMOD_SYSTEM
}

func CreateSystem() (*System, error) {
  var system System
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_System_Create(&system.system)
  })
  err := base.ResultToError(ferr)
  if err != nil {
    return nil, err
  }
  return &system, err
}

// TODO: Bind this to a finalizer
func (s *System) Release() error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_System_Release(s.system)
  })
  s.system = nil
  return base.ResultToError(ferr)
}

func (s *System) SetOutput(output OutputType) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_System_SetOutput(s.system, C.FMOD_OUTPUTTYPE(output))
  })
  return base.ResultToError(ferr)
}
func (s *System) GetOutput() (OutputType, error) {
  var output_type C.FMOD_OUTPUTTYPE
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_System_GetOutput(s.system, &output_type)
  })
  return OutputType(output_type), base.ResultToError(ferr)
}

func (s *System) GetNumDrivers() {}

// FMOD_RESULT FMOD_System_GetNumDrivers          (FMOD_SYSTEM *system, int *numdrivers);
func (s *System) GetDriverInfo() {}

// FMOD_RESULT FMOD_System_GetDriverInfo          (FMOD_SYSTEM *system, int id, char *name, int namelen, FMOD_GUID *guid);
func (s *System) GetDriverInfoW() {}

// FMOD_RESULT FMOD_System_GetDriverInfoW         (FMOD_SYSTEM *system, int id, short *name, int namelen, FMOD_GUID *guid);
func (s *System) GetDriverCaps() {}

// FMOD_RESULT FMOD_System_GetDriverCaps          (FMOD_SYSTEM *system, int id, FMOD_CAPS *caps, int *controlpaneloutputrate, FMOD_SPEAKERMODE *controlpanelspeakermode);
func (s *System) SetDriver() {}

// FMOD_RESULT FMOD_System_SetDriver              (FMOD_SYSTEM *system, int driver);
func (s *System) GetDriver() {}

// FMOD_RESULT FMOD_System_GetDriver              (FMOD_SYSTEM *system, int *driver);
func (s *System) SetHardwareChannels() {}

// FMOD_RESULT FMOD_System_SetHardwareChannels    (FMOD_SYSTEM *system, int numhardwarechannels);
func (s *System) SetSoftwareChannels() {}

// FMOD_RESULT FMOD_System_SetSoftwareChannels    (FMOD_SYSTEM *system, int numsoftwarechannels);
func (s *System) GetSoftwareChannels() {}

// FMOD_RESULT FMOD_System_GetSoftwareChannels    (FMOD_SYSTEM *system, int *numsoftwarechannels);
func (s *System) SetSoftwareFormat() {}

// FMOD_RESULT FMOD_System_SetSoftwareFormat      (FMOD_SYSTEM *system, int samplerate, FMOD_SOUND_FORMAT format, int numoutputchannels, int maxinputchannels, FMOD_DSP_RESAMPLER resamplemethod);
func (s *System) GetSoftwareFormat() {}

// FMOD_RESULT FMOD_System_GetSoftwareFormat      (FMOD_SYSTEM *system, int *samplerate, FMOD_SOUND_FORMAT *format, int *numoutputchannels, int *maxinputchannels, FMOD_DSP_RESAMPLER *resamplemethod, int *bits);
func (s *System) SetDSPBufferSize() {}

// FMOD_RESULT FMOD_System_SetDSPBufferSize       (FMOD_SYSTEM *system, unsigned int bufferlength, int numbuffers);
func (s *System) GetDSPBufferSize() {}

// FMOD_RESULT FMOD_System_GetDSPBufferSize       (FMOD_SYSTEM *system, unsigned int *bufferlength, int *numbuffers);
func (s *System) SetFileSystem() {}

// FMOD_RESULT FMOD_System_SetFileSystem          (FMOD_SYSTEM *system, FMOD_FILE_OPENCALLBACK useropen, FMOD_FILE_CLOSECALLBACK userclose, FMOD_FILE_READCALLBACK userread, FMOD_FILE_SEEKCALLBACK userseek, FMOD_FILE_ASYNCREADCALLBACK userasyncread, FMOD_FILE_ASYNCCANCELCALLBACK userasynccancel, int blockalign);
func (s *System) AttachFileSystem() {}

// FMOD_RESULT FMOD_System_AttachFileSystem       (FMOD_SYSTEM *system, FMOD_FILE_OPENCALLBACK useropen, FMOD_FILE_CLOSECALLBACK userclose, FMOD_FILE_READCALLBACK userread, FMOD_FILE_SEEKCALLBACK userseek);
func (s *System) SetAdvancedSettings() {}

// FMOD_RESULT FMOD_System_SetAdvancedSettings    (FMOD_SYSTEM *system, FMOD_ADVANCEDSETTINGS *settings);
func (s *System) GetAdvancedSettings() {}

// FMOD_RESULT FMOD_System_GetAdvancedSettings    (FMOD_SYSTEM *system, FMOD_ADVANCEDSETTINGS *settings);
func (s *System) SetSpeakerMode() {}

// FMOD_RESULT FMOD_System_SetSpeakerMode         (FMOD_SYSTEM *system, FMOD_SPEAKERMODE speakermode);
func (s *System) GetSpeakerMode() {}

// FMOD_RESULT FMOD_System_GetSpeakerMode         (FMOD_SYSTEM *system, FMOD_SPEAKERMODE *speakermode);
func (s *System) SetCallback() {}

// FMOD_RESULT FMOD_System_SetCallback            (FMOD_SYSTEM *system, FMOD_SYSTEM_CALLBACK callback);
func (s *System) SetPluginPath() {}

// FMOD_RESULT FMOD_System_SetPluginPath          (FMOD_SYSTEM *system, const char *path);
func (s *System) LoadPlugin() {}

// FMOD_RESULT FMOD_System_LoadPlugin             (FMOD_SYSTEM *system, const char *filename, unsigned int *handle, unsigned int priority);
func (s *System) UnloadPlugin() {}

// FMOD_RESULT FMOD_System_UnloadPlugin           (FMOD_SYSTEM *system, unsigned int handle);
func (s *System) GetNumPlugins() {}

// FMOD_RESULT FMOD_System_GetNumPlugins          (FMOD_SYSTEM *system, FMOD_PLUGINTYPE plugintype, int *numplugins);
func (s *System) GetPluginHandle() {}

// FMOD_RESULT FMOD_System_GetPluginHandle        (FMOD_SYSTEM *system, FMOD_PLUGINTYPE plugintype, int index, unsigned int *handle);
func (s *System) GetPluginInfo() {}

// FMOD_RESULT FMOD_System_GetPluginInfo          (FMOD_SYSTEM *system, unsigned int handle, FMOD_PLUGINTYPE *plugintype, char *name, int namelen, unsigned int *version);
func (s *System) SetOutputByPlugin() {}

// FMOD_RESULT FMOD_System_SetOutputByPlugin      (FMOD_SYSTEM *system, unsigned int handle);
func (s *System) GetOutputByPlugin() {}

// FMOD_RESULT FMOD_System_GetOutputByPlugin      (FMOD_SYSTEM *system, unsigned int *handle);
func (s *System) CreateDSPByPlugin() {}

// FMOD_RESULT FMOD_System_CreateDSPByPlugin      (FMOD_SYSTEM *system, unsigned int handle, FMOD_DSP **dsp);
func (s *System) RegisterCodec() {}

// FMOD_RESULT FMOD_System_RegisterCodec          (FMOD_SYSTEM *system, FMOD_CODEC_DESCRIPTION *description, unsigned int *handle, unsigned int priority);
func (s *System) RegisterDSP() {}

// FMOD_RESULT FMOD_System_RegisterDSP            (FMOD_SYSTEM *system, FMOD_DSP_DESCRIPTION *description, unsigned int *handle);

// FMOD_RESULT F_API FMOD_System_Init (FMOD_SYSTEM *system, int maxchannels, FMOD_INITFLAGS flags, void *extradriverdata);
func (s *System) Init(max_channels int, flags InitFlags, extra_driver_data interface{}) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_System_Init(s.system, C.int(max_channels), C.FMOD_INITFLAGS(flags), unsafe.Pointer(uintptr(0)))
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT FMOD_System_Close                  (FMOD_SYSTEM *system);
func (s *System) Close() error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_System_Close(s.system)
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT FMOD_System_Update                 (FMOD_SYSTEM *system);
func (s *System) Update() error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_System_Update(s.system)
  })
  return base.ResultToError(ferr)
}

func (s *System) Set3DSettings() {}

// FMOD_RESULT FMOD_System_Set3DSettings          (FMOD_SYSTEM *system, float dopplerscale, float distancefactor, float rolloffscale);
func (s *System) Get3DSettings() {}

// FMOD_RESULT FMOD_System_Get3DSettings          (FMOD_SYSTEM *system, float *dopplerscale, float *distancefactor, float *rolloffscale);
func (s *System) Set3DNumListeners() {}

// FMOD_RESULT FMOD_System_Set3DNumListeners      (FMOD_SYSTEM *system, int numlisteners);
func (s *System) Get3DNumListeners() {}

// FMOD_RESULT FMOD_System_Get3DNumListeners      (FMOD_SYSTEM *system, int *numlisteners);
func (s *System) Set3DListenerAttributes() {}

// FMOD_RESULT FMOD_System_Set3DListenerAttributes(FMOD_SYSTEM *system, int listener, const FMOD_VECTOR *pos, const FMOD_VECTOR *vel, const FMOD_VECTOR *forward, const FMOD_VECTOR *up);
func (s *System) Get3DListenerAttributes() {}

// FMOD_RESULT FMOD_System_Get3DListenerAttributes(FMOD_SYSTEM *system, int listener, FMOD_VECTOR *pos, FMOD_VECTOR *vel, FMOD_VECTOR *forward, FMOD_VECTOR *up);
func (s *System) Set3DRolloffCallback() {}

// FMOD_RESULT FMOD_System_Set3DRolloffCallback   (FMOD_SYSTEM *system, FMOD_3D_ROLLOFFCALLBACK callback);
func (s *System) Set3DSpeakerPosition() {}

// FMOD_RESULT FMOD_System_Set3DSpeakerPosition   (FMOD_SYSTEM *system, FMOD_SPEAKER speaker, float x, float y, FMOD_BOOL active);
func (s *System) Get3DSpeakerPosition() {}

// FMOD_RESULT FMOD_System_Get3DSpeakerPosition   (FMOD_SYSTEM *system, FMOD_SPEAKER speaker, float *x, float *y, FMOD_BOOL *active);
func (s *System) SetStreamBufferSize() {}

// FMOD_RESULT FMOD_System_SetStreamBufferSize    (FMOD_SYSTEM *system, unsigned int filebuffersize, FMOD_TIMEUNIT filebuffersizetype);
func (s *System) GetStreamBufferSize() {}

// FMOD_RESULT FMOD_System_GetStreamBufferSize    (FMOD_SYSTEM *system, unsigned int *filebuffersize, FMOD_TIMEUNIT *filebuffersizetype);

func (s *System) GetVersion() (uint, error) {
  var version C.uint
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_System_GetVersion(s.system, &version)
  })
  return uint(version), base.ResultToError(ferr)
}

func (s *System) GetOutputHandle() {}

// FMOD_RESULT FMOD_System_GetOutputHandle        (FMOD_SYSTEM *system, void **handle);

// FMOD_RESULT FMOD_System_GetChannelsPlaying     (FMOD_SYSTEM *system, int *channels);
func (s *System) GetChannelsPlaying() (int, error) {
  var n C.int
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_System_GetChannelsPlaying(s.system, &n)
  })
  err := base.ResultToError(ferr)
  if err != nil {
    return 0, err
  }
  return int(n), nil
}

func (s *System) GetHardwareChannels() {}

// FMOD_RESULT FMOD_System_GetHardwareChannels    (FMOD_SYSTEM *system, int *numhardwarechannels);
func (s *System) GetCPUUsage() {}

// FMOD_RESULT FMOD_System_GetCPUUsage            (FMOD_SYSTEM *system, float *dsp, float *stream, float *geometry, float *update, float *total);
func (s *System) GetSoundRAM() {}

// FMOD_RESULT FMOD_System_GetSoundRAM            (FMOD_SYSTEM *system, int *currentalloced, int *maxalloced, int *total);
func (s *System) GetNumCDROMDrives() {}

// FMOD_RESULT FMOD_System_GetNumCDROMDrives      (FMOD_SYSTEM *system, int *numdrives);
func (s *System) GetCDROMDriveName() {}

// FMOD_RESULT FMOD_System_GetCDROMDriveName      (FMOD_SYSTEM *system, int drive, char *drivename, int drivenamelen, char *scsiname, int scsinamelen, char *devicename, int devicenamelen);
func (s *System) GetSpectrum() {}

// FMOD_RESULT FMOD_System_GetSpectrum            (FMOD_SYSTEM *system, float *spectrumarray, int numvalues, int channeloffset, FMOD_DSP_FFT_WINDOW windowtype);
func (s *System) GetWaveData() {}

// FMOD_RESULT FMOD_System_GetWaveData            (FMOD_SYSTEM *system, float *wavearray, int numvalues, int channeloffset);

// FMOD_RESULT FMOD_System_CreateSound            (FMOD_SYSTEM *system, const char *name_or_data, FMOD_MODE mode, FMOD_CREATESOUNDEXINFO *exinfo, FMOD_SOUND **sound);
type Sound struct {
  sound *C.FMOD_SOUND
}

func (s *System) CreateSound_FromFilename(filename string, mode Mode) (*Sound, error) {
  var sound Sound
  b := makeNullTerminatedBytes(filename)
  cstr_filename := (*C.char)(unsafe.Pointer(&b[0])) // TODO: Why did I have to make this unsafe?
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_System_CreateSound(s.system, cstr_filename, C.FMOD_MODE(mode), (*C.FMOD_CREATESOUNDEXINFO)(null), &sound.sound)
  })
  err := base.ResultToError(ferr)
  if err != nil {
    return nil, err
  }
  return &sound, nil
}

//const char *name_or_data, FMOD_MODE mode, FMOD_CREATESOUNDEXINFO *exinfo, FMOD_SOUND **sound);
// FMOD_RESULT FMOD_System_CreateSound            (FMOD_SYSTEM *system, const char *name_or_data, FMOD_MODE mode, FMOD_CREATESOUNDEXINFO *exinfo, FMOD_SOUND **sound);

func (s *System) CreateStream() {}

// FMOD_RESULT FMOD_System_CreateStream           (FMOD_SYSTEM *system, const char *name_or_data, FMOD_MODE mode, FMOD_CREATESOUNDEXINFO *exinfo, FMOD_SOUND **sound);
func (s *System) CreateDSP() {}

// FMOD_RESULT FMOD_System_CreateDSP              (FMOD_SYSTEM *system, FMOD_DSP_DESCRIPTION *description, FMOD_DSP **dsp);
func (s *System) CreateDSPByType() {}

// FMOD_RESULT FMOD_System_CreateDSPByType        (FMOD_SYSTEM *system, FMOD_DSP_TYPE type, FMOD_DSP **dsp);

// FMOD_RESULT FMOD_System_CreateChannelGroup     (FMOD_SYSTEM *system, const char *name, FMOD_CHANNELGROUP **channelgroup);
func (s *System) CreateChannelGroup(name string) (*ChannelGroup, error) {
  b := makeNullTerminatedBytes(name)
  cstr_name := (*C.char)(unsafe.Pointer(&b[0])) // TODO: Why did I have to make this unsafe?
  var cg ChannelGroup
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_System_CreateChannelGroup(s.system, cstr_name, &cg.group)
  })
  err := base.ResultToError(ferr)
  if err != nil {
    return nil, err
  }
  return &cg, nil
}

func (s *System) CreateSoundGroup() {}

// FMOD_RESULT FMOD_System_CreateSoundGroup       (FMOD_SYSTEM *system, const char *name, FMOD_SOUNDGROUP **soundgroup);
func (s *System) CreateReverb() {}

// FMOD_RESULT FMOD_System_CreateReverb           (FMOD_SYSTEM *system, FMOD_REVERB **reverb);

// FMOD_RESULT FMOD_System_PlaySound              (FMOD_SYSTEM *system, FMOD_CHANNELINDEX channelid, FMOD_SOUND *sound, FMOD_BOOL paused, FMOD_CHANNEL **channel);
func (s *System) PlaySound(channel_id ChannelIndex, sound *Sound, paused bool) (*Channel, error) {
  var channel Channel
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_System_PlaySound(s.system, C.FMOD_CHANNELINDEX(channel_id), sound.sound, makeFmodBool(paused), &channel.channel)
  })
  err := base.ResultToError(ferr)
  if err != nil {
    return nil, err
  }
  return &channel, err
}

func (s *System) PlayDSP() {}

// FMOD_RESULT FMOD_System_PlayDSP                (FMOD_SYSTEM *system, FMOD_CHANNELINDEX channelid, FMOD_DSP *dsp, FMOD_BOOL paused, FMOD_CHANNEL **channel);
func (s *System) GetChannel() {}

// FMOD_RESULT FMOD_System_GetChannel             (FMOD_SYSTEM *system, int channelid, FMOD_CHANNEL **channel);

// FMOD_RESULT FMOD_System_GetMasterChannelGroup  (FMOD_SYSTEM *system, FMOD_CHANNELGROUP **channelgroup);
func (s *System) GetMasterChannelGroup() (*ChannelGroup, error) {
  var cg ChannelGroup
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_System_GetMasterChannelGroup(s.system, &cg.group)
  })
  err := base.ResultToError(ferr)
  if err != nil {
    return nil, err
  }
  return &cg, nil
}

func (s *System) GetMasterSoundGroup() {}

// FMOD_RESULT FMOD_System_GetMasterSoundGroup    (FMOD_SYSTEM *system, FMOD_SOUNDGROUP **soundgroup);
func (s *System) SetReverbProperties() {}

// FMOD_RESULT FMOD_System_SetReverbProperties    (FMOD_SYSTEM *system, const FMOD_REVERB_PROPERTIES *prop);
func (s *System) GetReverbProperties() {}

// FMOD_RESULT FMOD_System_GetReverbProperties    (FMOD_SYSTEM *system, FMOD_REVERB_PROPERTIES *prop);
func (s *System) SetReverbAmbientProperties() {}

// FMOD_RESULT FMOD_System_SetReverbAmbientProperties(FMOD_SYSTEM *system, FMOD_REVERB_PROPERTIES *prop);
func (s *System) GetReverbAmbientProperties() {}

// FMOD_RESULT FMOD_System_GetReverbAmbientProperties(FMOD_SYSTEM *system, FMOD_REVERB_PROPERTIES *prop);
func (s *System) GetDSPHead() {}

// FMOD_RESULT FMOD_System_GetDSPHead             (FMOD_SYSTEM *system, FMOD_DSP **dsp);
func (s *System) AddDSP() {}

// FMOD_RESULT FMOD_System_AddDSP                 (FMOD_SYSTEM *system, FMOD_DSP *dsp, FMOD_DSPCONNECTION **connection);
func (s *System) LockDSP() {}

// FMOD_RESULT FMOD_System_LockDSP                (FMOD_SYSTEM *system);
func (s *System) UnlockDSP() {}

// FMOD_RESULT FMOD_System_UnlockDSP              (FMOD_SYSTEM *system);
func (s *System) GetDSPClock() {}

// FMOD_RESULT FMOD_System_GetDSPClock            (FMOD_SYSTEM *system, unsigned int *hi, unsigned int *lo);
func (s *System) GetRecordNumDrivers() {}

// FMOD_RESULT FMOD_System_GetRecordNumDrivers    (FMOD_SYSTEM *system, int *numdrivers);
func (s *System) GetRecordDriverInfo() {}

// FMOD_RESULT FMOD_System_GetRecordDriverInfo    (FMOD_SYSTEM *system, int id, char *name, int namelen, FMOD_GUID *guid);
func (s *System) GetRecordDriverInfoW() {}

// FMOD_RESULT FMOD_System_GetRecordDriverInfoW   (FMOD_SYSTEM *system, int id, short *name, int namelen, FMOD_GUID *guid);
func (s *System) GetRecordDriverCaps() {}

// FMOD_RESULT FMOD_System_GetRecordDriverCaps    (FMOD_SYSTEM *system, int id, FMOD_CAPS *caps, int *minfrequency, int *maxfrequency);
func (s *System) GetRecordPosition() {}

// FMOD_RESULT FMOD_System_GetRecordPosition      (FMOD_SYSTEM *system, int id, unsigned int *position);
func (s *System) RecordStart() {}

// FMOD_RESULT FMOD_System_RecordStart            (FMOD_SYSTEM *system, int id, FMOD_SOUND *sound, FMOD_BOOL loop);
func (s *System) RecordStop() {}

// FMOD_RESULT FMOD_System_RecordStop             (FMOD_SYSTEM *system, int id);
func (s *System) IsRecording() {}

// FMOD_RESULT FMOD_System_IsRecording            (FMOD_SYSTEM *system, int id, FMOD_BOOL *recording);
func (s *System) CreateGeometry() {}

// FMOD_RESULT FMOD_System_CreateGeometry         (FMOD_SYSTEM *system, int maxpolygons, int maxvertices, FMOD_GEOMETRY **geometry);
func (s *System) SetGeometrySettings() {}

// FMOD_RESULT FMOD_System_SetGeometrySettings    (FMOD_SYSTEM *system, float maxworldsize);
func (s *System) GetGeometrySettings() {}

// FMOD_RESULT FMOD_System_GetGeometrySettings    (FMOD_SYSTEM *system, float *maxworldsize);
func (s *System) LoadGeometry() {}

// FMOD_RESULT FMOD_System_LoadGeometry           (FMOD_SYSTEM *system, const void *data, int datasize, FMOD_GEOMETRY **geometry);
func (s *System) GetGeometryOcclusion() {}

// FMOD_RESULT FMOD_System_GetGeometryOcclusion   (FMOD_SYSTEM *system, const FMOD_VECTOR *listener, const FMOD_VECTOR *source, float *direct, float *reverb);
func (s *System) SetNetworkProxy() {}

// FMOD_RESULT FMOD_System_SetNetworkProxy        (FMOD_SYSTEM *system, const char *proxy);
func (s *System) GetNetworkProxy() {}

// FMOD_RESULT FMOD_System_GetNetworkProxy        (FMOD_SYSTEM *system, char *proxy, int proxylen);
func (s *System) SetNetworkTimeout() {}

// FMOD_RESULT FMOD_System_SetNetworkTimeout      (FMOD_SYSTEM *system, int timeout);
func (s *System) GetNetworkTimeout() {}

// FMOD_RESULT FMOD_System_GetNetworkTimeout      (FMOD_SYSTEM *system, int *timeout);
func (s *System) SetUserData() {}

// FMOD_RESULT FMOD_System_SetUserData            (FMOD_SYSTEM *system, void *userdata);
func (s *System) GetUserData() {}

// FMOD_RESULT FMOD_System_GetUserData            (FMOD_SYSTEM *system, void **userdata);
func (s *System) GetMemoryInfo() {}

// FMOD_RESULT FMOD_System_GetMemoryInfo          (FMOD_SYSTEM *system, unsigned int memorybits, unsigned int event_memorybits, unsigned int *memoryused, FMOD_MEMORY_USAGE_DETAILS *memoryused_details);
