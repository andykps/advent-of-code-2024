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

[Lazyvim](https://www.lazyvim.org) is proving to be pretty annoying. I'm comfortable with Vim motions and it's nice to have some hints as to invalid syntax Go from the LSP. However, it has a habit of inserting completions that I just don't want and the non text hints that are in your editor but not really in your source are pretty annoying. Since I don't know what plugins are causing these effects I'm thinking about using a more vanilla [kickstart.nvim](https://github.com/nvim-lua/kickstart.nvim) setup and knowing what I'm enabling. It'd be nice to have something to autosave like IntelliJ does.

## Day 3
Yay regex! Pt 2 took a bit longer because of needing to remember to add a multiline modifier. I should have twigged earlier since my first attempt at part 1 was processing the file line by line because I'd copied from previous days, until I realised that I needed to match on the entire content.

I tried kickstart.nvim but it removed too much. Hopefully given some more time I'll find a good middle ground. An annoyance that came up today is the different keyboard that I use in the office. It's a US keyboard that I have a UK layout mapped onto because I expect `"` to be above the `2` and `@` to be 2 to the left of `L` etc. (and I'm not looking at the keyboard to find them). However, it means I need to press `<Alt-Gr> \` to get a `\` and this for some reason (possibly a combination with Kanata) causes `<Ctrl>` to get locked on causing all sorts of weirdness until I realise and mash `<Ctrl>` to resolve it.

## Day 4
I found out that I don't really understand how Go slices are backed by arrays. I'm guessing the slice is passed by value but it's just a pointer to the actual data in the backing array. I got stuck on part 1 which worked fine on the test data but then failed on the real data. It turns out it is because the slices that I was adding to my grid data were backed by the bufio buffer used for reading from the file, so after reading beyond that buffer size it started changing what I could see in my grid data slices and returning the wrong XMAS count.

I thought I was doing it cleverly by checking rows, cols and diagonals for XMAS or SAMX. It worked fine but then wasn't at all useful for part 2 so they both operate differently.

