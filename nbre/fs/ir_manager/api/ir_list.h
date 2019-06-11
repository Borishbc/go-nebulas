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
#include "fs/ir_manager/api/ir_item_list_base.h"

namespace neb {
namespace fs {
class ir_list {
public:
  ir_list(rocksdb_storage *storage);

  virtual void write_ir(const nbre::NBREIR &raw_ir,
                        const nbre::NBREIR &compiled_ir);

  virtual nbre::NBREIR get_raw_ir(const std::string &ir_name, version_t v);
  virtual nbre::NBREIR get_ir(const std::string &ir_name, version_t v);

protected:
  std::unordered_map<std::string,
                     std::shared_ptr<internal::ir_item_list_interface>>
      m_ir_item_list;
};
} // namespace fs
} // namespace neb
