package fmod_event

// #include "fmod_event.h"
// #include "stdlib.h"
import "C"

import (
  "errors"
  "github.com/MobRulesGames/fmod/base"
)

type Param struct {
  param *C.FMOD_EVENTPARAMETER
}

// FMOD_RESULT F_API FMOD_EventParameter_GetInfo        (FMOD_EVENTPARAMETER *eventparameter, int *index, char **name);
func (p *Param) GetInfo() error {
  return errors.New("Param.GetInfo() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventParameter_GetRange       (FMOD_EVENTPARAMETER *eventparameter, float *rangemin, float *rangemax);
func (p *Param) GetRange() (min, max float64, err error) {
  var ferr C.FMOD_RESULT
  var cmin, cmax C.float
  base.Thread(func() {
    ferr = C.FMOD_EventParameter_GetRange(p.param, &cmin, &cmax)
  })
  min = float64(cmin)
  max = float64(cmax)
  err = base.ResultToError(ferr)
  return
}

// FMOD_RESULT F_API FMOD_EventParameter_SetValue       (FMOD_EVENTPARAMETER *eventparameter, float value);
func (p *Param) SetValue(value float64) error {
  var ferr C.FMOD_RESULT
  base.Thread(func() {
    ferr = C.FMOD_EventParameter_SetValue(p.param, C.float(value))
  })
  return base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_EventParameter_GetValue       (FMOD_EVENTPARAMETER *eventparameter, float *value);
func (p *Param) GetValue() (float64, error) {
  var ferr C.FMOD_RESULT
  var value C.float
  base.Thread(func() {
    ferr = C.FMOD_EventParameter_GetValue(p.param, &value)
  })
  return float64(value), base.ResultToError(ferr)
}

// FMOD_RESULT F_API FMOD_EventParameter_SetVelocity    (FMOD_EVENTPARAMETER *eventparameter, float value);
func (p *Param) SetVelocity() error {
  return errors.New("Param.SetVelocity() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventParameter_GetVelocity    (FMOD_EVENTPARAMETER *eventparameter, float *value);
func (p *Param) GetVelocity() error {
  return errors.New("Param.GetVelocity() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventParameter_SetSeekSpeed   (FMOD_EVENTPARAMETER *eventparameter, float value);
func (p *Param) SetSeekSpeed() error {
  return errors.New("Param.SetSeekSpeed() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventParameter_GetSeekSpeed   (FMOD_EVENTPARAMETER *eventparameter, float *value);
func (p *Param) GetSeekSpeed() error {
  return errors.New("Param.GetSeekSpeed() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventParameter_SetUserData    (FMOD_EVENTPARAMETER *eventparameter, void *userdata);
func (p *Param) SetUserData() error {
  return errors.New("Param.SetUserData() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventParameter_GetUserData    (FMOD_EVENTPARAMETER *eventparameter, void **userdata);
func (p *Param) GetUserData() error {
  return errors.New("Param.GetUserData() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventParameter_KeyOff         (FMOD_EVENTPARAMETER *eventparameter);
func (p *Param) KeyOff() error {
  return errors.New("Param.KeyOff() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventParameter_DisableAutomation(FMOD_EVENTPARAMETER *eventparameter, FMOD_BOOL disable);
func (p *Param) DisableAutomation() error {
  return errors.New("Param.DisableAutomation() has not been implemented yet.")
}

// FMOD_RESULT F_API FMOD_EventParameter_GetMemoryInfo  (FMOD_EVENTPARAMETER *eventparameter, unsigned int memorybits, unsigned int event_memorybits, unsigned int *memoryused, FMOD_MEMORY_USAGE_DETAILS *memoryused_details);
func (p *Param) GetMemoryInfo() error {
  return errors.New("Param.GetMemoryInfo() has not been implemented yet.")
}
