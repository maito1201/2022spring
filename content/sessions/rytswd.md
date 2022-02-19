---
key: rytswd
title: 'Go Module with Microservices and Monorepo: Clear Dependencies with Ease of
  Development'
id: rytswd
format: conference
talkType: short_session
level: All
tags: []
speakers:
- rytswd
videoId: null
presentation: null
draft: false
---
Go Module is certainly an integral part of the Go ecosystem today. But perhaps because of its relatively recent inception, it can be confusing how to integrate with some specific needs. In this talk, we'll see how Go module can work extremely well with microservices and monorepo.
---
Go Module is great. It allows simple development workflow, without too much prior knowledge of `GOPATH` or anything special. When you start your new Go project, the first thing is to run `go mod init github.com/me/myrepo`, and you are all set. It just works. But when you try to do something a bit more involved, such as having a private repository, having a monorepo with many modules, or both, it can be confusing how to wire up all the dependencies.

In this talk, we will take a quick look at some implementation details of Go Module, and some key concepts such as `GOPRIVATE` environment variable and the like.

From there, we will work our way up to setting up a monorepo with handful of microservices, in a private repository, and with our own domain. We will look at how dependency management is made clear and easy with such setup.