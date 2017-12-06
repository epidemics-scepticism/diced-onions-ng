# diced-onions-ng
Generate vanity onions in Go

Generating a specific prefix

    $ go get github.com/epidemics-scepticism/diced-onions-ng
    $ ~/go/bin/diced-onions-ng -wordlist=test -full=false
    2017/12/06 23:31:58 full search: false
    2017/12/06 23:31:58 wordlist: test
    2017/12/06 23:31:58 workers: 4
    2017/12/06 23:31:58 added 1 words
    2017/12/06 23:31:58 thread 1 starting
    2017/12/06 23:31:58 thread 2 starting
    2017/12/06 23:31:58 thread 3 starting
    2017/12/06 23:31:58 thread 4 starting
    2017/12/06 23:31:58 press ctrl-c to exit
    2017/12/06 23:32:00 match: lolpngx2x7ytkdzpby7lhfdwxeh3hvkgdnxfdbwgisqw6pgxjxggtbyd.onion
    2017/12/06 23:32:01 match: lol3hnvejoyhnholv66jewstyrlihwbbjcoyhkw2oi3hhsgigubnmwqd.onion
    2017/12/06 23:32:03 match: loljszvvnrvifr46fz4l5uppfi2ykuyol6sk2qte27xfje3wz7lk7kid.onion
    ^C2017/12/06 23:32:05 signal received, exiting
    2017/12/06 23:32:05 thread 2 stopping
    2017/12/06 23:32:05 thread 1 stopping
    2017/12/06 23:32:05 thread 3 stopping
    2017/12/06 23:32:05 thread 4 stopping

Generating a full onion is likely infeasible, given the keyspace.
