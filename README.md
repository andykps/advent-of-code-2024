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

## Day 5
I've learnt to use a struct to hold data. It still feels like I'm missing a lot of knowledge about the capabilities of Go. I should read a book sometime...

## Day 6
Part one was fun. I just brute forced part two by running the simulation until it went over a limit (gridSize\*gridSize). It ran quickly enough. It was easier to just copy paste part one and work differently than trying to do in the same file.

## Day 7
I had to read up on permutations again. Initially my function to generate the perms didn't work because I wasn't copying the slice each time. Seems I've still not got my head around how slices are backed by arrays. Gemini was able to point out my error though. It attempted to optimise it for me too but I didn't really understand the result so left as is. ≈15s to run part 2. I did wonder if the longer run time was because of all the conversion between int and string and back again but I tried out multiplying by 10 to the power of number of digits before adding the next operand and it made little difference. The longer time must be the increased number of permutations.
I've found out about the `range` keyword which made iterating over slices easier.

## Day 8
I used a Go map type to store the positions of the antennas and a struct to hold the coordinate data, that hopefully makes things more readable. I had ended up with a type definition of `map[byte][][2]int` which wasn't. It did lead me to search for conventions on naming of custom types in Go. I didn't really find anything conclusive.

I've tried to split up more into functions and it feels like some of these should be imported from a shared package(?) since I keep reusing them. I think that it helped me generalise the solution for part 2.

## Day 9
Ok, I gave up with Neovim... temporarily, at least for Go development. I can see how it will make editing easier and require less hand movement but it's slowing me down when I don't understand Go well enough either. I switched to VSCode because it much easier to set up the debugger (I clicked the install button).

I also found that in VSCode there is less type hinting that whilst nice to see some of the time, really clutters up what I'm looking at. I need to find out which plugin is causing that and tweak the options or turn it off.

Part 1 went quickly enough today but I got really stuck on part 2 because there was a bug in my code that only became apparent when using the real data. If the only available gap was just before the file being moved then it didn't move it. It was because in my defrag loop I was only finding available space up to, but not including the block before the first block of the file to be moved and this didn't happen in the test data.

Rendering the blocks was helpful with test data but once the blocks numbers went into the extended ASCII codes over 127 then I started getting unrenderable characters. I stopped using `byte`s to render and switched to `rune`s instead as well as adding `0x7F600` so that I got emojis. Not because it was useful but it created some fun output.

<img src="day9/emojis.png">