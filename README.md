# npwut??
> A quick way to return desired npm run methods to terminal.

# Background
> I built this for my own quality of life, I found it was helpful while working with many projects at once. I also wanted to work with Go on more simple CLI projects. Yes, I know you can easily run a few UNIX comamnds with a pipe to get the same result.

# Steps
* [X] Read a folder
* [X] Create read buffer for package.json
* [X] Build structs based on commonly used phrases for npm run save to local variable
* [X] Decode package json file for any key names in 'scripts' object save to local variable
* [X] Parse given args for the npm command the user is looking for
* [X] Return the command to output
* [X] Scan specific directories (server/client)
* [ ] If not found, return the entire list. User is able to hit a key to select desired command to paste to output.

## How to build from source and run anywhere
```
$ git clone https://github.com/ncolletti/npwut
$ cd npwut
$ go build
$ mkdir ~/bin
$ mv npwut ~/bin
BASH: export PATH=$PATH:/home/<user>/bin
ZSH: path+=('/home/<user>/bin')
```

# TODO
* [ ] Environment preset (dev/prod)
* [ ] Write tests