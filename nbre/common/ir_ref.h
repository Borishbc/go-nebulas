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
#include "util/version.h"

namespace neb {
class ir_ref {
public:
  inline ir_ref() {}
  inline ir_ref(const std::string &name, util::version &v)
      : m_name(name), m_version(v) {}

  inline const std::string &name() const { return m_name; }
  inline std::string &name() { return m_name; }

  inline const util::version &version() const { return m_version; }
  inline util::version &version() { return m_version; }

protected:
  std::string m_name;
  util::version m_version;
};
}
