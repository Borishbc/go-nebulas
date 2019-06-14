// Copyright (C) 2017 go-nebulas authors
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
#include "core/net_ipc/client/client_context.h"
#include "fs/storage.h"

namespace neb {
namespace core {
client_context::client_context() : execution_context() {}
client_context::~client_context() {
  m_bc_storage.reset();
  m_nbre_storage.reset();
}

fs::storage *client_context::blockchain_storage() { return m_bc_storage.get(); }
fs::storage *client_context::nbre_storage() { return m_nbre_storage.get(); }

void client_context::shutdown() {
  m_bc_storage.reset();
  m_nbre_storage.reset();
}

} // namespace core
} // namespace neb
