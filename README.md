# go-tkgtools

source code for project go-tkgtools

this tool is an implementation of f1-f5, f1*, f5* functions defined in 3GPP 35.206.


Use is as:

import "github.com/hyrut/go-tkgtools"

And get start by:

tkg := tkgtools.NewTKGTOOLS()

tkg.F1(&key, &rand, &sqn, &amf, &mac_a, &op, nil)

tkg.F2345(&key, &rand, &res, &ck, &ik, &ak, &op, nil)

