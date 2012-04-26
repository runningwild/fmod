package fmod

// #include "fmod.h"
import "C"

type Dsp struct {
  dsp *C.FMOD_DSP
}

