# runegen #

Random string generator.

### What is this repository for? ###

* This is the main repository for random string generator library and command line.
* Version 1.1.1

### How do I get set up? ###

* import the library in go lang code:
  
```bash
$# go get github.com/berrypay/runegen
```

```go
import (
  "fmt"
  "github.com/berrypay/runegen"
)

func main() {
  fmt.Println(runegen.GetRandom(1, 6))
}
```

There's also a command line tool 'runegen' which can be used to generate random string. To compile:

```bash
$# make
$# cd bin
$# ./runegen -c 7 -l 8 --strict=true -startalpha=true

```
Then you can find the binary in the 'bin' folder.

### License ###

**MIT License**

Copyright (c) 2022 Sallehuddin Abdul Latif <sallehuddin@berrypay.com>

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

### Contribution guidelines ###

* Writing tests
* Code review
* Other guidelines

### Who do I talk to? ###

* Any issues? Contact sallehuddin[AT]gmail.com
