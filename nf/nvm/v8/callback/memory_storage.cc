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

#include "memory_storage.h"

#include <atomic>
#include <mutex>
#include <string>
#include <unordered_map>
#include <stdio.h>
#include <stdlib.h>

static std::mutex mapMutex;
static std::unordered_map<std::string, std::string> memoryMap;

char* StorageGet(V8Engine* engine, void* handler, const char *key, size_t *cnt){
  NVMCallbackResponse *res = new NVMCallbackResponse();
  res->set_func_name(std::string(STORAGE_GET));
  res->add_func_params(std::string(key));

  const NVMCallbackResult *callback_res = SNVM::DataExchangeCallback(engine, handler, res);
  *cnt = (size_t)std::stoull(callback_res->extra(0));
  std::string resString = callback_res->result();
  bool not_null_flag = callback_res->not_null();
  if(callback_res != nullptr)
    delete callback_res;

  if(!not_null_flag)
    return NULL;
  char* cStr = (char*)calloc(resString.length()+1, sizeof(char));
  strcpy(cStr, resString.c_str());
  return cStr;
}

int StoragePut(V8Engine* engine, void* handler, const char* key, const char *value, size_t *cnt){
  NVMCallbackResponse *res = new NVMCallbackResponse();
  res->set_func_name(std::string(STORAGE_PUT));
  res->add_func_params(std::string(key));
  res->add_func_params(std::string(value));

  const NVMCallbackResult* callback_res = SNVM::DataExchangeCallback(engine, handler, res);
  *cnt = (size_t)std::stoull(callback_res->extra(0));
  int resCode = std::stoi(callback_res->result());
  if(callback_res != nullptr)
    delete callback_res;

  return resCode;
}

int StorageDel(V8Engine* engine, void* handler, const char* key, size_t *cnt){
  NVMCallbackResponse *res = new NVMCallbackResponse();
  res->set_func_name(std::string(STORAGE_DEL));
  res->add_func_params(std::string(key));

  const NVMCallbackResult* callback_res = SNVM::DataExchangeCallback(engine, handler, res);
  *cnt = (size_t)std::stoull(callback_res->extra(0));
  int resCode = std::stoi(callback_res->result());
  if(callback_res != nullptr)
    delete callback_res;

  return resCode;
}