# justify

Change the alignment of the output (center, left, right, justify)

## Author

* [@smslash](https://github.com/smslash)


## Objectives

You must follow the same [instructions](https://01.alem.school/git/root/public/src/branch/master/subjects/ascii-art/README.md) as in the first subject but the alignment can be changed.

```bash
We
        will
                explain!
```

To change the alignment of the output it must be possible to use a **flag** `--align=<type>`, in which `type` can be :

- center

- left

- right

- justify

- You must adapt your representation to the terminal size. If you reduce the terminal window the graphical representation should be adapted to the terminal size.

- Only text that fits the terminal size will be tested.

- The flag must have exactly the same format as above, any other formats must return the following usage message:

```
Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run . --align=right something standard
```

If there are other `ascii-art` optional projects implemented, the program should accept other correctly formatted `[OPTION]` and/or `[BANNER]`.

Additionally, the program must still be able to run with a single `[STRING]` argument.

## Instructions

- Your project must be written in **Go**.

- The code must respect the [good practices](https://01.alem.school/git/root/public/src/branch/master/subjects/good-practices/README.md).

- It is recommended to have **test files** for [unit testing](https://go.dev/doc/tutorial/add-a-test).

## Usage

Assume the bars in the display below are the terminal borders:

```
|$ go run . --align=center "hello" standard                                                                                 |
|                                             _                _    _                                                       |
|                                            | |              | |  | |                                                      |
|                                            | |__      ___   | |  | |    ___                                               |
|                                            |  _ \    / _ \  | |  | |   / _ \                                              |
|                                            | | | |  |  __/  | |  | |  | (_) |                                             |
|                                            |_| |_|   \___|  |_|  |_|   \___/                                              |
|                                                                                                                           |
|                                                                                                                           |
|$ go run . --align=left "Hello There" standard                                                                             |
| _    _           _    _                 _______   _                                                                       |
|| |  | |         | |  | |               |__   __| | |                                                                      |
|| |__| |   ___   | |  | |    ___           | |    | |__      ___    _ __     ___                                           |
||  __  |  / _ \  | |  | |   / _ \          | |    |  _ \    / _ \  | '__|   / _ \                                          |
|| |  | | |  __/  | |  | |  | (_) |         | |    | | | |  |  __/  | |     |  __/                                          |
||_|  |_|  \___|  |_|  |_|   \___/          |_|    |_| |_|   \___|  |_|      \___|                                          |
|                                                                                                                           |
|                                                                                                                           |
|$ go run . --align=right "hello" shadow                                                                                    |
|                                                                                                                           |
|                                                                                          _|                _| _|          |
|                                                                                          _|_|_|     _|_|   _| _|   _|_|   |
|                                                                                          _|    _| _|_|_|_| _| _| _|    _| |
|                                                                                          _|    _| _|       _| _| _|    _| |
|                                                                                          _|    _|   _|_|_| _| _|   _|_|   |
|                                                                                                                           |
|                                                                                                                           |
|$ go run . --align=justify "how are you" shadow                                                                            |
|                                                                                                                           |
|_|                                                                                                                         |
|_|_|_|     _|_|   _|      _|      _|                  _|_|_| _|  _|_|   _|_|                    _|    _|   _|_|   _|    _| |
|_|    _| _|    _| _|      _|      _|                _|    _| _|_|     _|_|_|_|                  _|    _| _|    _| _|    _| |
|_|    _| _|    _|   _|  _|  _|  _|                  _|    _| _|       _|                        _|    _| _|    _| _|    _| |
|_|    _|   _|_|       _|      _|                      _|_|_| _|         _|_|_|                    _|_|_|   _|_|     _|_|_| |
|                                                                                                      _|                   |
|                                                                                                  _|_|                     |
|$       
```

## Allowed packages

- Only the [standard Go](https://pkg.go.dev/std) packages are allowed
This project will help you learn about :

- The Go file system(**fs**) API

- Data manipulation

- Terminal display