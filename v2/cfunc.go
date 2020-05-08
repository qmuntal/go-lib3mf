/*++

Copyright (C) 2019 3MF Consortium (Original Author)

All rights reserved.

Redistribution and use in source and binary forms, with or without modification,
are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice,
this list of conditions and the following disclaimer in the documentation
and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.7.0-develop.

Abstract: This is an autogenerated Go wrapper file in order to allow an easy
use of the 3MF Library.

Interface version: 2.0.0

*/

// Code generated by Automatic Component Toolkit (ACT); DO NOT EDIT.

package lib3mf

/*
#include "lib3mf_types.h"

extern void progressCallback(bool *, Lib3MF_double, eLib3MFProgressIdentifier, Lib3MF_pvoid);
void Lib3MFProgressCallback_cgo(bool * pAbort, Lib3MF_double dProgressValue, eLib3MFProgressIdentifier eProgressIdentifier, Lib3MF_pvoid pUserData){
  progressCallback(pAbort, dProgressValue, eProgressIdentifier, pUserData);
}

extern void writeCallback(Lib3MF_uint64, Lib3MF_uint64, Lib3MF_pvoid);
void Lib3MFWriteCallback_cgo(Lib3MF_uint64 nByteData, Lib3MF_uint64 nNumBytes, Lib3MF_pvoid pUserData){
  writeCallback(nByteData, nNumBytes, pUserData);
}

extern void readCallback(Lib3MF_uint64, Lib3MF_uint64, Lib3MF_pvoid);
void Lib3MFReadCallback_cgo(Lib3MF_uint64 nByteData, Lib3MF_uint64 nNumBytes, Lib3MF_pvoid pUserData){
  readCallback(nByteData, nNumBytes, pUserData);
}

extern void seekCallback(Lib3MF_uint64, Lib3MF_pvoid);
void Lib3MFSeekCallback_cgo(Lib3MF_uint64 nPosition, Lib3MF_pvoid pUserData){
  seekCallback(nPosition, pUserData);
}
*/
import "C"
