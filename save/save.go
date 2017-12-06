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
package save

import (
	"encoding/hex"
	"golang.org/x/crypto/ed25519"
	"io/ioutil"
	"log"
)

func SaveKey(priv ed25519.PrivateKey, o string) {
	if e := ioutil.WriteFile(o+".onion", []byte(priv), 0600); e != nil {
		log.Print("failed to save "+o+".onion to file:", e)
		log.Print(o+"\n", hex.EncodeToString([]byte(priv)))
	}
}
