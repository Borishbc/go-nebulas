// Copyright (C) 2018 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or
// modify
// it under the terms of the GNU General Public License as published by
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
#pragma once
#include "common/common.h"
#include "util/quitable_thread.h"
#include <boost/process/child.hpp>
#include <chrono>

namespace neb {
namespace core {
class ipc_client_watcher : public util::quitable_thread {
public:
  ipc_client_watcher(const std::string &path);

  virtual ~ipc_client_watcher();

  inline bool is_client_alive() { return m_b_client_alive; };
  void kill_client();

protected:
  virtual void thread_func();

protected:
  std::string m_path;
  uint32_t m_restart_times;
  std::chrono::system_clock::time_point m_last_start_time;
  std::atomic_bool m_b_client_alive;
  std::unique_ptr<boost::process::child> m_client;
  std::mutex m_mutex;
  bool m_killed_already;

}; // end class ipc_client_watcher
} // namespace core
} // namespace neb
