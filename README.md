Bindings for [FMOD](www.fmod.org) version 4.40.04

These bindings are very incomplete, there is only enough right now for the most basic sound playing requirements.

**Thread Safety**: Fmod functions must be called from the same thread, so this library does lock an os thread and runs all fmod functions on that thread.  This means that you can freely call any fmod function from any go-routine or thread safely.
