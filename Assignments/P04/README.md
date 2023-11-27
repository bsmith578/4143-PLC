## P04 - Concurrent Image Downloader
### Brett Smith
### Description:

This program test two methods of downloading images from the internet.
The first is downloading each image sequentially and the other downloads
the images concurrently. The program tests each method and prints the
time taken for each to show how much faster concurrent downloading is
using Goroutines. Goroutines are lightweight threads that are managed
during the runtime of a program.

### Files

|   #   | File            | Description                                        |
| :---: | --------------- | -------------------------------------------------- |
|   1   | [main.go]() | Main driver code for the program |
|   2   | [go.mod]() | contains module's properties required to run the program |