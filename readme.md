# gitlab-applet

`gitlab-applet` is a cli wrapper built to make managing and retrieving project information easier from gitlab's v4 api. It's built using [cobra](https://github.com/spf13/cobra), and [go-gitlab](https://github.com/xanzy/go-gitlab).

## Usage
Clone this repo `git clone https://github.com/giancarlopetrini/gitlab-applet` and build locally, or `cd gitlab-applet` and `go run main.go <flags> <command> <commandflags>`

## Top Level Flags
* `token` gitlab auth token
* `giturl` gitlab url (`/api/v4/` will be appended)
* `project` project name (or close match to name)

Example:  
`./gitlab-applet --token yourtokehere --project helloworld --giturl gitlab.com`

## Commands
Commands follow the top level flag declarations. Each command may have arguments that are passed to it as well. 
For Example:
```
gitlab-applet --token yourtokehere --project helloworld --giturl gitlab.com show variables
```

Here's the current list of supported commands:
* show - retrieves information from a project
  * variables
    * gets project level variables. Decodes if in base64 format (used at times if CI integrates with container deployments)