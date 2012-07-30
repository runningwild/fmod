package fmod_ex

// #include "fmod.h"
import "C"

type DspConn struct {
  conn *C.FMOD_DSPCONNECTION
}
