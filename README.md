# q

> Work in progress. Things **will** break and change significantly.

**q** is a queue-based task runner that simply runs tasks that you schedule
sequentially.

The **motivation** behind building **q** is simple. I run long-running tasks, and while doing so I prepare new tasks to be run when the ones already running finish. I can't run the in parallel most of the time.

Some features I'm planning:

- Jobs should have a way to report progress.
- Jobs should have a way to report other things in a simple non-intrusive way.
- Jobs should be able to pause/resume/be stopped/killed/etc.
- All stdout and stderr is saved.
- A simple CLI tool to add and manage jobs. If you run your task with `./hello`, you should queue it with `q ./hello`.

I make no promise I'll build all these. I'll probably stop when I have something that's good enough for my use cases.


## Architecture

How I imagine the project architecture now. It might change as I write code.

1. There is a job runner daemon that holds a job queue. This daemon runs as a *server*.
2. The client(s) communicate to the daemon (*server*) to add jobs, get their status, etc.

There should be some persistence layer for the daemon to store the job queue in case of failure, but it's not a priority for now. Everything will be in memory for the sake of simplicity. Persistence can come later on.
