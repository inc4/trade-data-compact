# `/cmd`

This directory contains a code for the executable. If the project has only
one executable you can use `/cmd/main.go`. Otherwise, it is necessary to create
`/cmd/<bin>/<bin>.go`. Where "bin" is the name of the executable.

Don't put a lot of code here, use packages from `/pkg` or `/internal`.
