# go-tkgtools

TKGTOOLS stands for Tgpp Key Generator Tools.

It implemente f1 - f5, f1*, f5* functions defined in 3GPP TS 35.205/35.206.

These functions are also known as MILENAGE Algorithm Set.

Test data could be find in TS 35.208.

Specification is here: https://www.3gpp.org/ftp/Specs/archive/35_series

[![GoDoc](https://godoc.org/github.com/hyrut/go-tkgtools?status.svg)](https://godoc.org/github.com/hyrut/go-tkgtools)

### Status
* The current implementation is solid and works fine for general purpose.
* Additional, besides default r1-r5, c1-c5. R and C values modification is supported by current implementation.
* Check examples folder for detail.

[![Build Status](https://secure.travis-ci.org/hyrut/go-tkgtools.png)](http://travis-ci.org/hyrut/go-tkgtools)

### Get started

Just go with command to get:
```
go get github.com/hyrut/go-tkgtools
```
And import it as:
```
import "github.com/hyrut/go-tkgtools"
```

New instance as:
```
tkg = tkgtools.NewTKGTOOLS()
tkg.F1(&key, &rand, &sqn , &amf, &mac_a, &op, nil)
tkg.F2345(&key, &rand, &res, &ck, &ik, &ak, &op, nil)
```
Build examples for detail.

### Performance

With default R and C value, functions calculation performance is below:

* F1 is around 1400ns.

* F2345 is around 3500ns.

Tested on a Intel(R) Core(TM) i5-8350U CPU @ 1.70GHz 1.90GHz, Windows 10.
