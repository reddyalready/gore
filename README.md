# gore

Simple golang command line utility that, using goroutines:

 * Watches the current directory for file system changes; and
 * Reruns a given command if files changed

## Installing

Get the source:

`git clone https://github.com/reddyalready/gore.git`

Get dependencies (you must have `dep` - Go's dependency manager)

`dep ensure`

Install gore

`go build`

## Running

Simply run by passing the program as command line arguments, e.g.:

`gore echo restarted`

Make a change to a file in the current directory and see the magic!

## MORE!

Okay, I exaggerated with the "magic". 

This was just a small side project for me to understand goroutines 
and channels. If you'd like to use it, and would like some features,
please file an issue and I can learn more!
