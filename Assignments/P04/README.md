## P04 - Concurrent Image Downloader
### Brett Smith
### Description:

This program test two methods of downloading images from the internet.
The first is downloading each image sequentially and the other downloads
the images concurrently. The program tests each method and prints the
time taken for each to show how much faster concurrent downloading is
using Goroutines. Goroutines are lightweight threads that are managed
during the runtime of a program. After running the program the sequential
download of the five images took 862 ms while the concurrent download
took 296 ms.

### Files

|   #   | File            | Description                                        |
| :---: | --------------- | -------------------------------------------------- |
|   1   | [main.go](https://github.com/bsmith578/4143-PLC/blob/main/Assignments/P04/main.go) | Main driver code for the program |
|   2   | [go.mod](https://github.com/bsmith578/4143-PLC/blob/main/Assignments/P04/go.mod) | contains module's properties required to run the program |