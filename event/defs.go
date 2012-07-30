package fmod_event

// #include "fmod_event.h"
import "C"

const VERSION = C.FMOD_EVENT_VERSION

type InitFlags C.FMOD_EVENT_INITFLAGS

const (
  NORMAL                           InitFlags = 0x00000000 /* All platforms - Initialize normally */
  USER_ASSETMANAGER                InitFlags = 0x00000001 /* All platforms - All wave data loading/freeing will be referred back to the programmer through the FMOD_EVENT_CALLBACKTYPE_SOUNDDEF_CREATE/FMOD_EVENT_CALLBACKTYPE_SOUNDDEF_RELEASE callback */
  FAIL_ON_MAXSTREAMS               InitFlags = 0x00000002 /* All platforms - Events will fail if "Max streams" was reached when playing streamed banks, instead of going virtual. */
  DONTUSENAMES                     InitFlags = 0x00000004 /* All platforms - All event/eventgroup/eventparameter/eventcategory/eventreverb names will be discarded on load. Use getXXXByIndex to access them. This may potentially save a lot of memory at runtime. */
  UPPERCASE_FILENAMES              InitFlags = 0x00000008 /* All platforms - All FSB filenames will be translated to upper case before being used. */
  LOWERCASE_FILENAMES              InitFlags = 0x00000080 /* All platforms - All FSB filenames will be translated to lower case before being used. */
  SEARCH_PLUGINS                   InitFlags = 0x00000010 /* All platforms - Search the current directory for dsp/codec plugins on EventSystem::init. */
  USE_GUIDS                        InitFlags = 0x00000020 /* All platforms - Build an event GUID table when loading FEVs so that EventSystem::getEventByGUID can be used. */
  DETAILED_SOUNDDEF_INFO           InitFlags = 0x00000040 /* All platforms - Pass an FMOD_EVENT_SOUNDDEFINFO struct to FMOD_EVENT_CALLBACKTYPE_SOUNDDEF_SELECTINDEX callbacks rather than just the sound definition name (uses more memory for sound definition waveform names). */
  RESETPARAMSTOMINIMUM             InitFlags = 0x00000100 /* All platforms - Reset parameters to minimum value when getting an event instance instead of using the INFO_ONLY event's values. */
  ELEVATION_AFFECTS_LISTENER_ANGLE InitFlags = 0x00000200 /* All platforms - The listener angle event parameters will be affected by elevation, and not just horizontal components. */
  DONTUSELOWMEM                    InitFlags = 0x00000400 /* All platforms - Instruct the event system to NOT use FMOD_LOWMEM when it opens .FSB files. Specify this flag if you need access to the names of individual subsounds in loaded .FSB files. Specifying this flag will make the event system use more memory. */
)
