package cipher

import (
	"io"
)

/*
Task 3: Rot 13

This task is taken from http://tour.golang.org.

A common pattern is an io.Reader that wraps another io.Reader, modifying the
stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of
compressed data) and returns a *gzip.Reader that also implements io.Reader (a
stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader,
modifying the stream by applying the rot13 substitution cipher to all
alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing
its Read method.
*/

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	if err == nil {
		var i int
		// 'A' = 65, 'Z' = 64 + 26 = 90
		// 'a' = 97, 'z' = 96 + 26 = 122
		// go through the byte list. If num is between A and Z the letter is ... else if (check for small)
		// calculation letter = (num - firstletter)+rotnum % alphabetlength + firstletter
		const A, Z, a, z = 65, 90, 97, 122
		const alphabetLen = 26
		for i = 0; i < len(p); i++ {
			if p[i] >= A && p[i] < Z {
				p[i] = (((p[i] - A) + 13) % alphabetLen) + A
			} else if p[i] >= a && p[i] <= z {
				p[i] = (((p[i] - a) + 13) % alphabetLen) + a
			}
		}
	}
	return n, err
}
