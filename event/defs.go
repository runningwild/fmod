package fmod_event

// #include "fmod_event.h"
import "C"

const VERSION = C.FMOD_EVENT_VERSION

type InitFlags C.FMOD_INITFLAGS

const (
  INIT_NORMAL                    InitFlags = C.FMOD_INIT_NORMAL
  INIT_STREAM_FROM_UPDATE                  = C.FMOD_INIT_STREAM_FROM_UPDATE
  INIT_3D_RIGHTHANDED                      = C.FMOD_INIT_3D_RIGHTHANDED
  INIT_SOFTWARE_DISABLE                    = C.FMOD_INIT_SOFTWARE_DISABLE
  INIT_OCCLUSION_LOWPASS                   = C.FMOD_INIT_OCCLUSION_LOWPASS
  INIT_HRTF_LOWPASS                        = C.FMOD_INIT_HRTF_LOWPASS
  INIT_DISTANCE_FILTERING                  = C.FMOD_INIT_DISTANCE_FILTERING
  INIT_SOFTWARE_REVERB_LOWMEM              = C.FMOD_INIT_SOFTWARE_REVERB_LOWMEM
  INIT_ENABLE_PROFILE                      = C.FMOD_INIT_ENABLE_PROFILE
  INIT_VOL0_BECOMES_VIRTUAL                = C.FMOD_INIT_VOL0_BECOMES_VIRTUAL
  INIT_WASAPI_EXCLUSIVE                    = C.FMOD_INIT_WASAPI_EXCLUSIVE
  INIT_PS3_PREFERDTS                       = C.FMOD_INIT_PS3_PREFERDTS
  INIT_PS3_FORCE2CHLPCM                    = C.FMOD_INIT_PS3_FORCE2CHLPCM
  INIT_DISABLEDOLBY                        = C.FMOD_INIT_DISABLEDOLBY
  INIT_SYSTEM_MUSICMUTENOTPAUSE            = C.FMOD_INIT_SYSTEM_MUSICMUTENOTPAUSE
  INIT_SYNCMIXERWITHUPDATE                 = C.FMOD_INIT_SYNCMIXERWITHUPDATE
  INIT_GEOMETRY_USECLOSEST                 = C.FMOD_INIT_GEOMETRY_USECLOSEST
  INIT_DISABLE_MYEARS_AUTODETECT           = C.FMOD_INIT_DISABLE_MYEARS_AUTODETECT
)

type EventInitFlags C.FMOD_EVENT_INITFLAGS

const (
  EVENT_INIT_NORMAL                           EventInitFlags = C.FMOD_EVENT_INIT_NORMAL                           /* All platforms - Initialize normally */
  EVENT_INIT_USER_ASSETMANAGER                               = C.FMOD_EVENT_INIT_USER_ASSETMANAGER                /* All platforms - All wave data loading/freeing will be referred back to the programmer through the FMOD_EVENT_CALLBACKTYPE_SOUNDDEF_CREATE/FMOD_EVENT_CALLBACKTYPE_SOUNDDEF_RELEASE callback */
  EVENT_INIT_FAIL_ON_MAXSTREAMS                              = C.FMOD_EVENT_INIT_FAIL_ON_MAXSTREAMS               /* All platforms - Events will fail if "Max streams" was reached when playing streamed banks, instead of going virtual. */
  EVENT_INIT_DONTUSENAMES                                    = C.FMOD_EVENT_INIT_DONTUSENAMES                     /* All platforms - All event/eventgroup/eventparameter/eventcategory/eventreverb names will be discarded on load. Use getXXXByIndex to access them. This may potentially save a lot of memory at runtime. */
  EVENT_INIT_UPPERCASE_FILENAMES                             = C.FMOD_EVENT_INIT_UPPERCASE_FILENAMES              /* All platforms - All FSB filenames will be translated to upper case before being used. */
  EVENT_INIT_LOWERCASE_FILENAMES                             = C.FMOD_EVENT_INIT_LOWERCASE_FILENAMES              /* All platforms - All FSB filenames will be translated to lower case before being used. */
  EVENT_INIT_SEARCH_PLUGINS                                  = C.FMOD_EVENT_INIT_SEARCH_PLUGINS                   /* All platforms - Search the current directory for dsp/codec plugins on EventSystem::init. */
  EVENT_INIT_USE_GUIDS                                       = C.FMOD_EVENT_INIT_USE_GUIDS                        /* All platforms - Build an event GUID table when loading FEVs so that EventSystem::getEventByGUID can be used. */
  EVENT_INIT_DETAILED_SOUNDDEF_INFO                          = C.FMOD_EVENT_INIT_DETAILED_SOUNDDEF_INFO           /* All platforms - Pass an FMOD_EVENT_SOUNDDEFINFO struct to FMOD_EVENT_CALLBACKTYPE_SOUNDDEF_SELECTINDEX callbacks rather than just the sound definition name (uses more memory for sound definition waveform names). */
  EVENT_INIT_RESETPARAMSTOMINIMUM                            = C.FMOD_EVENT_INIT_RESETPARAMSTOMINIMUM             /* All platforms - Reset parameters to minimum value when getting an event instance instead of using the INFO_ONLY event's values. */
  EVENT_INIT_ELEVATION_AFFECTS_LISTENER_ANGLE                = C.FMOD_EVENT_INIT_ELEVATION_AFFECTS_LISTENER_ANGLE /* All platforms - The listener angle event parameters will be affected by elevation, and not just horizontal components. */
  EVENT_INIT_DONTUSELOWMEM                                   = C.FMOD_EVENT_INIT_DONTUSELOWMEM                    /* All platforms - Instruct the event system to NOT use FMOD_LOWMEM when it opens .FSB files. Specify this flag if you need access to the names of individual subsounds in loaded .FSB files. Specifying this flag will make the event system use more memory. */
)

type Mode C.FMOD_EVENT_MODE

const (
  MODE_DEFAULT             Mode = C.FMOD_EVENT_DEFAULT             /* FMOD_EVENT_DEFAULT specifies default loading behaviour i.e. event data for the whole group is NOT cached and the function that initiated the loading process will block until loading is complete. */
  MODE_NONBLOCKING              = C.FMOD_EVENT_NONBLOCKING         /* For loading event data asynchronously. FMOD will use a thread to load the data.  Use Event::getState to find out when loading is complete. */
  MODE_ERROR_ON_DISKACCESS      = C.FMOD_EVENT_ERROR_ON_DISKACCESS /* For EventGroup::getEvent / EventGroup::getEventByIndex.  If EventGroup::loadEventData has accidently been forgotten this flag will return an FMOD_ERR_FILE_UNWANTED if the getEvent function tries to load data. */
  MODE_INFOONLY                 = C.FMOD_EVENT_INFOONLY            /* For EventGroup::getEvent / EventGroup::getEventByIndex.  Don't allocate instances or load data, just get a handle to allow user to get information from the event. */
  MODE_USERDSP                  = C.FMOD_EVENT_USERDSP             /* For EventGroup::getEvent / EventGroup::getEventByIndex.  Tells FMOD that you plan to add your own DSP effects to this event's ChannelGroup at runtime. Omitting this flag will yield a small memory gain. */
  MODE_NONBLOCKING_THREAD0      = C.FMOD_EVENT_NONBLOCKING_THREAD0 /* FMOD_EVENT_NONBLOCKING, execute on thread 0.  See remarks. (default) */
  MODE_NONBLOCKING_THREAD1      = C.FMOD_EVENT_NONBLOCKING_THREAD1 /* FMOD_EVENT_NONBLOCKING, execute on thread 1.  See remarks. */
  MODE_NONBLOCKING_THREAD2      = C.FMOD_EVENT_NONBLOCKING_THREAD2 /* FMOD_EVENT_NONBLOCKING, execute on thread 2.  See remarks. */
  MODE_NONBLOCKING_THREAD3      = C.FMOD_EVENT_NONBLOCKING_THREAD3 /* FMOD_EVENT_NONBLOCKING, execute on thread 3.  See remarks. */
  MODE_NONBLOCKING_THREAD4      = C.FMOD_EVENT_NONBLOCKING_THREAD4 /* FMOD_EVENT_NONBLOCKING, execute on thread 4.  See remarks. */
)

type State C.FMOD_EVENT_STATE

const (
  STATE_READY          State = C.FMOD_EVENT_STATE_READY          /* Event is ready to play. */
  STATE_LOADING              = C.FMOD_EVENT_STATE_LOADING        /* Loading in progress. */
  STATE_ERROR                = C.FMOD_EVENT_STATE_ERROR          /* Failed to open - file not found, out of memory etc.  See return value of Event::getState for what happened. */
  STATE_PLAYING              = C.FMOD_EVENT_STATE_PLAYING        /* Event has been started.  This will still be true even if there are no sounds active.  Event::stop must be called or the event must stop itself using a 'one shot and stop event' parameter mode. */
  STATE_CHANNELSACTIVE       = C.FMOD_EVENT_STATE_CHANNELSACTIVE /* Event has active voices.  Use this if you want to detect if sounds are playing in the event or not. */
  STATE_INFOONLY             = C.FMOD_EVENT_STATE_INFOONLY       /* Event was loaded with the FMOD_EVENT_INFOONLY flag. */
  STATE_STARVING             = C.FMOD_EVENT_STATE_STARVING       /* Event is streaming but not being fed data in time, so may be stuttering. */
  STATE_NEEDSTOLOAD          = C.FMOD_EVENT_STATE_NEEDSTOLOAD    /* Event still needs to load wavebank data. */
)

type Resource C.FMOD_EVENT_RESOURCE

const (
  RESOURCE_STREAMS_AND_SAMPLES Resource = C.FMOD_EVENT_RESOURCE_STREAMS_AND_SAMPLES /* Open all streams and load all banks into memory, under this group (recursive) */
  RESOURCE_STREAMS                      = C.FMOD_EVENT_RESOURCE_STREAMS             /* Open all streams under this group (recursive).  No samples are loaded. */
  RESOURCE_SAMPLES                      = C.FMOD_EVENT_RESOURCE_SAMPLES             /* Load all banks into memory, under this group (recursive).  No streams are opened. */
)
