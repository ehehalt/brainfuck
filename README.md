# Brainfuck in Go

Brainf\*ck Interpreter in Go. Based on the article [A Virtual Brainfuck Machine in Go](https://thorstenball.com/blog/2017/01/04/a-virtual-brainfuck-machine-in-go) from [Thorsten Ball](https://thorstenball.com).

Times to calculate the mandelbrot sample on a MacBook Pro with Go 1.12.1:

| type              | time   |
|:------------------|-------:|
| bf-interpreter    | 52.16s |
| bf-virtualmachine | 10.41s |
| bf-compiler       |  2.94s |

The compiler variant (not in the article) generates Go code from the mandelbrot sample. This could be compiled as executable:

```bash
bf-compiler mandelbrot.bf > mandelbrot.go && go build mandelbrot.go
time ./mandelbrot
```

