/*
   Copyright (C) 2016 cacahuatl < cacahuatl at autistici dot org >
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package onion

import (
	"encoding/base32"
	"golang.org/x/crypto/sha3"
	"golang.org/x/crypto/ed25519"
)

var b32 *base32.Encoding = base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567")

func Hash(pub ed25519.PublicKey) string {
	value := append([]byte(".onion checksum"), []byte(pub)...)
	value = append(value, 3)
	cksum := sha3.Sum256(value[:])
	rawonion := append([]byte(pub), cksum[:2]...)
	rawonion = append(rawonion, 3)
	return b32.EncodeToString(rawonion)
}
