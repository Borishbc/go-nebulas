// Copyright (C) 2017-2019 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or
// modify it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see
// <http://www.gnu.org/licenses/>.
//
// Author: Samuel Chen <samuel.chen@nebulas.io>

#pragma once

#include <stddef.h>
#include "../nvm_engine.h"
#include "callback_util.h"

char *Sha256(V8Engine*, const char *data, size_t *counterVal);
char *Sha3256(V8Engine*, const char *data, size_t *counterVal);
char *Ripemd160(V8Engine*, const char *data, size_t *counterVal);
char *RecoverAddress(V8Engine*, int alg, const char *data, const char *sign, size_t *counterVal);
char *Md5(V8Engine*, const char *data, size_t *counterVal);
char *Base64(V8Engine*, const char *data, size_t *counterVal);