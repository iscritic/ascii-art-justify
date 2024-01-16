## ascii-art-justify

### Objectives

ASCII Art Justify is a command-line utility for generating ASCII art text with left, right, and center alignment.

### Instructions

- Some **alignment** files with a specific graphical template representation using ASCII can be used. 

  - `left`
  - `right`
  - `center`
  - `justify`


- Some **banner** files with a specific graphical template representation using ASCII can be used. 

  - [shadow](shadow.txt)
  - [standard](standard.txt)
  - [thinkertoy](thinkertoy.txt)


### Usage

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
|$                                                                                                                          |
```

### Author

- [msarvaro](https://01.alem.school/git/msarvaro)
- [arthkim](https://01.alem.school/git/arthkim)