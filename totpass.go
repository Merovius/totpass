// Copyright 2019 Axel Wagner
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"time"
)

func main() {
	log.SetFlags(0)

	period := flag.Duration("period", 30*time.Second, "Validity-period of OTPs")
	digits := flag.Int("digits", 6, "Number of digits to use")

	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("Usage: totpass <name>")
	}

	cmd := exec.Command("pass", "totp/"+flag.Arg(0))
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	secret, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(string(bytes.TrimSpace(out)))
	if err != nil {
		log.Fatal("Invalid TOTP secret")
	}
	h := hmac.New(sha1.New, secret)

	c := uint64(time.Now().Unix() / int64(*period/time.Second))
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], c)
	h.Write(b[:])

	s := h.Sum(nil)
	offs := int(s[len(s)-1] & 0x0F)
	m := binary.BigEndian.Uint32(s[offs:offs+4]) &^ (1 << 31)
	fmt.Println(m % modulus[*digits])
}

var modulus = []uint32{1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9, math.MaxUint32}
