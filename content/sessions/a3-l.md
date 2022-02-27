---
key: a3-l
title: Dissecting Slices, Maps and Channels in Go
id: jespinog
format: conference
talkType: long_session
level: Intermediate
tags: []
speakers:
- jespinog
partner: ""
videoId: null
presentation: null
draft: false
---
My talk is about how slices, maps and channels work in the Go runtime. The idea is to use unsafe to extract the memory state on runtime and analyze how that is changing over time when you operate with the slices, maps and channels. We will see how channel buffers or maps buckets work under the hood.
---
Slices, Maps and Channels in go, are first class citizens of the language, they are used widely in our day to day work, but... do you know how they works under the hood? Do you know the implications of adding elements to an slice, or new keys to a map? Do you know why you can't relay in maps order? Do you know how channels handle the buffer or the blocked goroutines? If you don't know about that, this is your talk. I going to access the go runtime memory state of the maps, slices and channels, and show you how they evolve over time while we change them.