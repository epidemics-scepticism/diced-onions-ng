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
package main

import (
	"crypto/rand"
	"flag"
	"github.com/epidemics-scepticism/diced-onions-ng/onion"
	"github.com/epidemics-scepticism/diced-onions-ng/save"
	"github.com/epidemics-scepticism/diced-onions-ng/search"
	"log"
	"golang.org/x/crypto/ed25519"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"
)

func worker(die chan bool, tid int) {
	defer log.Print("thread ", tid, " stopping")
	defer w.Done()
	for {
		select {
		case <-die:
			return
		default:
			pub, priv, e := ed25519.GenerateKey(rand.Reader)
			if e != nil {
				log.Print("error: ", e)
			}
			m.Lock()
			c++
			m.Unlock()
			o := onion.Hash(pub)
			if match.Search(o, full) {
				log.Print("match: ", o, ".onion")
				save.SaveKey(priv, o)
			}
		}
	}
}

var match *search.Search

var (
	w    sync.WaitGroup
	m    sync.Mutex
	c    int = 0
	full bool
)

func main() {
	wordlist := flag.String("wordlist", "words", "wordlist to use")
	fullf := flag.Bool("full", true, "match only full matches")
	workers := flag.Int("workers", runtime.NumCPU(), "number of workers to spawn")
	flag.Parse()
	full = *fullf
	log.Print("full search: ", full)
	log.Print("wordlist: ", *wordlist)
	log.Print("workers: ", *workers)
	runtime.GOMAXPROCS(*workers)
	die := make(chan bool, 1)
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, os.Kill)
	match = search.NewSearch()
	if e := match.Populate(*wordlist); e != nil {
		log.Print("error: ", e)
		return
	}
	for i := 0; i < *workers; i++ {
		log.Print("thread ", i+1, " starting")
		w.Add(1)
		go worker(die, i+1)
	}
	start := time.Now()
	t := time.NewTicker(1 * time.Minute)
	log.Print("press ctrl-c to exit")
	for {
		select {
		case <-s:
			log.Print("signal received, exiting")
			close(die)
			w.Wait()
			return
		case now := <-t.C:
			delta := now.Sub(start)
			m.Lock()
			count := c
			m.Unlock()
			log.Print(count, " generated in ", int(delta.Seconds()), " seconds. (", count/int(delta.Seconds()), " per sec)")
		}
	}
}
