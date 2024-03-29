
<h1 align="center">
  <br>
  next
  <br>
</h1>

<h4 align="center">A simple todo CLI app, using .md files in your project directories. Written in Go.</h4>

![next3](https://user-images.githubusercontent.com/2995732/188331439-32f49b88-49bd-4bd3-b519-9e41ef9bb2e2.gif)

## Features

* Simple, intuitive commands
  - ```next``` - show current task
  - ```next done``` - complete current task
  - ```next todo``` - show all tasks in todo list
  - ```next completed``` - show all completed tasks
  - ```next add "make fancy github readme"``` - create new task in todo list
* Saved locally in markdown format
* ***Fancy styling***


## Installation

Simply clone this repo and run ```go install```, or if you prefer you can ```go build``` and place the binary in one of your PATH directories. 

## How to use

```bash
# In your project directory, initialize the markdown file
$ next init

# Add your todo tasks
$ next add "make fancy github readme"

# Start working on the task (optionally opens vs code from current directory also)
$ next start t1

# Complete your current task
$ next done
```

