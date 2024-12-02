# Advent of Code 2024

My attempts at AOC 2024. I'm trying to use Go which so far is proving exceptionally frustrating...

Please don't treat this code as examples of good Go!

I'm also trying to use [Neovim](https://neovim.io) from my terminal without needing to touch my mouse. (I'm still having to use the mouse when I switch to the browser to read all the Go docs I need to learn how to solve the task).
To save having too much keyboard hand contortion I've enabled home row mods with [Kanata](https://github.com/jtroo/kanata). I was going to try switching my layout to [Colmak Mod-DH](https://colemakmods.github.io/mod-dh/) but honestly I've got enough problems.
## Day 1
No idea if I'm structuring things well for a Go project. Currently each day is a Go module and inside that there are folders for each star's task containing a file with the main method. I'm putting input files (but not committing) in the day folder and then calling with e.g.
```
go run day1-1/day1-1.go
```
Let's see if I refine this over time. It'd be nice to be able to pass in the input to be used rather than have to edit the source to use example input but I haven't figured out command line args yet!

## Day 2
I've done both solutions in the same file today. Just pass a `--dampener` switch to enable the "Problem Dampener". You can also now specify the file to read on the command line but it defaults to `input.txt`

It took me a while to get through the second part. I had a horrible unreadable mess of whether to skip characters which forced me into actually reading how to write functions in Go. It wasn't that hard! It'd be interesting to know whether I'm doing anything dangerous with copying arrays inside functions and then returning them. I haven't really got my head around slices yet. It feels like I should be able to have a slice that is backed by an array without me needing to copy it to have a slice without a specific element.

[Lazyvim](https://www.lazyvim.org) is proving to be pretty annoying. I'm comfortable with Vim motions and it's nice to have some hints as to invalid syntax Go from the LSP. However, it has a habit of inserting completions that I just don't want and the non text hints that are in your editor but not really in your source are pretty annoying. Since I don't know what plugins are causing these effects I'm thinking about using a more vanilla [kickstart.vim](https://github.com/nvim-lua/kickstart.nvim) setup and knowing what I'm enabling. It'd be nice to have something to autosave like IntelliJ does.
