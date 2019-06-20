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

#include "require_callback.h"
#include "../engine.h"
#include "file.h"
#include "global.h"
#include "logger.h"

#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <iostream>

using namespace v8;

static char source_require_format[] =
    "(function(){\n"
    "return function (exports, module, require) {\n"
    "%s\n"
    "};\n"
    "})();\n";

static RequireDelegateFunc sRequireDelegate = NULL;
static AttachLibVersionDelegateFunc attachLibVersionDelegate = NULL;
static FetchNativeJSLibContentDelegateFunc fetchNativeJSLibContentDelegate = NULL;

inline bool checkLibWhite(char* filename){
  bool res = false;
  if(strncmp(filename, "lib/contract", 12) == 0 || strncmp(filename, "contract", 8) == 0)
    res = true;
  return res;
}

// native_lib_flag indicates if loading native js lib source code
static int readSource(Local<Context> context, const char *filename, char **data, size_t *lineOffset, bool native_lib_flag) {
  if (strstr(filename, "\"") != NULL) {
    return -1;
  }

  *lineOffset = 0;
  std::string content;

  if(native_lib_flag){
     std::cout<<"^^^^^^^^^^^^^^ Before read source content with native lib flag is true"<<std::endl;

    if(fetchNativeJSLibContentDelegate != NULL){
      content = fetchNativeJSLibContentDelegate(filename);
    }

  }else{
    std::cout<<"^^^^^^^^^^^^^^ Before read source content"<<std::endl;

    // try sRequireDelegate.
    if (sRequireDelegate != NULL) {
      V8Engine *e = GetV8EngineInstance(context);
      content = sRequireDelegate(e, filename, lineOffset);
    }
  }

  if(content.length()>0){
    asprintf(data, source_require_format, content.c_str());
    *lineOffset += -2;
    return 0;
  }

  return -1;
}

void attachVersion(char *out, int maxoutlen, Local<Context> context, const char *libname) {

  //const char *verlib;
  //char* verlib = nullptr;
  std::string verlib;
  
  if (attachLibVersionDelegate != NULL) {
    V8Engine *e = GetV8EngineInstance(context);
    verlib = attachLibVersionDelegate(e, libname);
  }

  if (verlib.length()>0) {
    std::cout<<"$$$$$$$ >>>>> verlib is not NULL: "<<verlib.c_str()<<std::endl;
    strncat(out, verlib.c_str(), maxoutlen - strlen(out) - 1);
  }
}

void NewNativeRequireFunction(Isolate *isolate,
                              Local<ObjectTemplate> globalTpl) {
  globalTpl->Set(String::NewFromUtf8(isolate, "_native_require"),
                 FunctionTemplate::New(isolate, RequireCallback),
                 static_cast<PropertyAttribute>(PropertyAttribute::DontDelete |
                                                PropertyAttribute::ReadOnly));
}

void RequireCallback(const v8::FunctionCallbackInfo<v8::Value> &info) {
  Isolate *isolate = info.GetIsolate();
  Local<Context> context = isolate->GetCurrentContext();

  if (info.Length() == 0) {
    isolate->ThrowException(
        Exception::Error(String::NewFromUtf8(isolate, "require missing path")));
    return;
  }

  Local<Value> path = info[0];
  if (!path->IsString()) {
    isolate->ThrowException(Exception::Error(
        String::NewFromUtf8(isolate, "require path must be string")));
    return;
  }

  String::Utf8Value filename(path);
  if (filename.length() >= MAX_PATH_LEN) {
    isolate->ThrowException(Exception::Error(
        String::NewFromUtf8(isolate, "require path length more")));
    return;
  }

  std::cout<<"############## Required filename is: "<<*filename<<std::endl;

  char *abPath = NULL;
  //if (strcmp(*filename, LIB_WHITE)) { // check if it's library js files, if so, read them from local path with version attached.
  bool is_native_lib_flag = false;
  if (!checkLibWhite(*filename)){
    is_native_lib_flag = true;
    char versionlizedPath[MAX_VERSIONED_PATH_LEN] = {0};
    attachVersion(versionlizedPath, MAX_VERSIONED_PATH_LEN, context, *filename);
    std::cout<<"$$$$$---- Attached versionlized path: "<<versionlizedPath<<std::endl;
    abPath = realpath(versionlizedPath, NULL);
    if (abPath == NULL) {
      isolate->ThrowException(Exception::Error(String::NewFromUtf8(
          isolate, "require path is invalid absolutepath")));
      return;
    }
    char curPath[MAX_VERSIONED_PATH_LEN] = {0};
    if (curPath[0] == 0x00 && !getCurAbsolute(curPath, MAX_VERSIONED_PATH_LEN)) {
      isolate->ThrowException(Exception::Error(
          String::NewFromUtf8(isolate, "invalid cwd absolutepath")));
      free(abPath);
      return;
    }
    int curLen = strlen(curPath);
    if (strncmp(abPath, curPath, curLen) != 0) {
      isolate->ThrowException(Exception::Error(
          String::NewFromUtf8(isolate, "require path is not in lib")));
      free(abPath);
      return;
    } 

    //free(abPath);
    if (!isFile(abPath)) {
      isolate->ThrowException(Exception::Error(
          String::NewFromUtf8(isolate, "require path is not file")));
      free(abPath);
      return;
    }
  }

  char *pFile = abPath;
  if (abPath == NULL) {
    pFile = *filename;
  }
  char *data = NULL;
  size_t lineOffset = 0;
  if (readSource(context, (const char*)pFile, &data, &lineOffset, is_native_lib_flag)) {
    char msg[512];
    snprintf(msg, 512, "require cannot find module '%s'", pFile);
    isolate->ThrowException(
        Exception::Error(String::NewFromUtf8(isolate, msg)));
    free(abPath);
    return;
  }
  free(abPath);

  ScriptOrigin sourceSrcOrigin(path, Integer::New(isolate, lineOffset));
  MaybeLocal<Script> script = Script::Compile(
      context, String::NewFromUtf8(isolate, data), &sourceSrcOrigin);
  if (!script.IsEmpty()) {
    MaybeLocal<Value> ret = script.ToLocalChecked()->Run(context);
    if (!ret.IsEmpty()) {
      Local<Value> rr = ret.ToLocalChecked();
      info.GetReturnValue().Set(rr);
    }
  }

  free(static_cast<void *>(data));
}

void InitializeRequireDelegate(
  RequireDelegateFunc delegate, 
  AttachLibVersionDelegateFunc aDelegate, 
  FetchNativeJSLibContentDelegateFunc fDelegate){

  sRequireDelegate = delegate;
  attachLibVersionDelegate = aDelegate;
  fetchNativeJSLibContentDelegate = fDelegate;
}