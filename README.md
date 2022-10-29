# h3x

A simple hex viewer/editor!
As of right now it only supports outputting a files hexdump.

Examplary Usage:

```bash
go run h3x.go input.bin
```

Example output:

```bash
$ go run h3x.go input.bin                                                                                                                       [0]
13 1c 9c a9 6c 05 57 bc 61 38 37 26 ce d9 21 f8 | ....l.W.a87&..!.|
61 32 15 04 d9 0b 7a 74 ce 66 dc 85 e6 7c ce b8 | a2....zt.f...|..|
68 2c a4 5e e1 8f 11 81 98 23 06 f3 0e cc ed 82 | h,.^.....#......|
1c b8 06 66 e8 08 f9 2b d6 1c 83 a6 c2 ef 2f 35 | ...f...+....../5|
8d 9b a1 0b f5 19 f6 cc bd 69 70 3a e1 5b b8 60 | .........ip:.[.`|
e4 16 5f 6a 33 b5 c7 53 d4 25 18 b2 15 14 05 29 | .._j3..S.%.....)|
a4 db f1 f2 02 24 68 1c 45 b3 c8 e2 72 73 80 aa | .....$h.E...rs..|
f3 61 dd 27 c4 51 29 0d e3 9a 9e ea 80 93 df 55 | .a.'.Q)........U|
8a dc f4 32 fd 8c 4e bc 75 8a 97 43 69 aa a8 53 | ...2..N.u..Ci..S|
f9 ce b2 a5 11 92 3f 20 6c f6 f1 fc c7 ec cc 82 | ......? l.......|
ea 63 36 11 75 0a bb bc e7 76 88 e4 9e d8 5c 92 | .c6.u....v....\.|
79 73 24 16 a2 0a 8f c4 a4 88 66 3b 52 1b 0e 7b | ys$.......f;R..{|
c0 da 6b a2 e7 3f 3d 10 29 46 7d 7e 4f ed 9f 04 | ..k..?=.)F}~O...|
dd e9 83 90 9f 85 b0 34 3b e4 15 d7 f7 89 8a 82 | .......4;.......|
d3 bc c5 71 0a c1 89 64 3c 7b e1 04 f7 d9 c6 f1 | ...q...d<{......|
3b 74 0b 64 51 ec bf e9 a6 81 d8 d7 d5 e2 23 9f | ;t.dQ.........#.|
```
