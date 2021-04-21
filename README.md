# Workwork

_A simple dictionary for listing and opening URLs for common software development concerns._

## Concept

Software development requires a lot of resources and services. When you're neck deep in a project or a veteran on your team at work, this is no problem. You know where everything is and how all the systems fits together.
Then there are everyone else: New hires on the teams, handoffs to other teams or companies, and legacy projects that no one remembers anything about, let alone where to find logs and how the build pipeline works.

This CLI utility (simply invoked as `ww` in a terminal) simply stores a dictionary of URLs for common software development concerns and services, such as documentation, task board, issues, getting started info, build server, deploy server, logs dashboard, monitoring dashboard, and so on.
You can of course register whatever URLs you want, in addition to - or completely replacing - the default ones.

When these URLs has been saved, the `ww ls` command will make it simple to get a complete listing of all the project's related resources, and the `ww goto` command lets you open a browser to go there.

Example:
```yaml
// .workwork.yaml

urls:
  docs: https://my-documentation.com/my-app
  ci: https://circleci.com
  newbie: https://my-documentation.com/company/onboarding

environments:
  - name: local
    urls:
      live: http://localhost:8080/my-app
      logs: https://localhost:8081/logs

  - name: prod
    urls:
      live: http://my-app.com
      logs: https://my-prod-logs.com
      monitoring: https://grafana.com
```

- `ww goto docs` would open the corresponding URL in a browser.
- `ww goto ci`
- `ww goto local.logs` open environment urls.
- `ww goto prod.live` 

And so on.

Besides being useful for your team, it's also super convenient to just run `ww goto ci` to open the build server in a browser or `ww goto pulls` to open the pull requests instead of manually switching to the browser and finding a bookmark or - god forbid - type the URL yourself.

## Installation

### Download archive

1. Head over to the [Releases page](https://github.com/eaardal/workwork/releases) and download the archive for your OS.
2. Extract the file and put it in a directory that is available on your PATH.

### Go

If you have Go installed, simply installing the repo should make it available on your PATH (it will be put in your $GOPATH/bin).
```
go install github.com/eaardal/workwork
```

_Note: This will install the executable as `workwork` and not `ww` so you'll probably want to make an alias for convenience._

## Getting started

1. In a terminal, `cd` to your repository.
2. Run `ww init`. You will be prompted for several URLs for default concerns built into the app. This will result in a `.workwork.yaml` file in your repository.

That's it for the initial setup. From here you can:

- Open `.workwork.yaml` in any text editor and add/remove/edit URLs and environments as you want.
- Use `ww ls` to explore the contents of `.workwork.yaml`.
- Use `ww get` to see a particular URL (or many at once).
- Use `ww set` to add a new URL or update an existing one.
- Use `ww rm` to remove one or more URLs.

See the `--help` texts for each particular command for more info.

## Commands

```
NAME:
   WorkWork - A simple dictionary for listing and opening URLs for common software development concerns

USAGE:
   ww [global options] command [command options] [arguments...]

COMMANDS:
   init     Create a .workwork file with default content
   ls       List all the registered URLs
   goto     Open the URL for the given key using the default browser
   set      Set a new URL for an existing key, or add a new URL if the key doesn't exist
   get      Show the URL for a specific key
   rm       Remove the URL for the given key
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## TODO

- More options when running `init` to make a blank `.workwork.yaml` without running through the prompts.
- Export markdown and thereby html of the links, making it possible to copy & paste the URLs as markdown to other documentation systems or host the html page or something.
- Better support for `ww goto` command to open many urls at once.