# github.com/dan-compton/exile

Exile (pronounced ex~~L~~ile) is a wholly re-written, language-agnostic, and generic version of [Lile](https://github.com/lileio/lile).  In short, it is a _simple_ code-generation framework built on go templates.

### Installation

**NOTE**  Due to the inclusion of golang plugins, the build process is not ensured to be stable outside of a linux-based build-host.  Plugins are currently not supported on windows ([issue](https://github.com/golang/go/issues/19282)).  

* Install the latest version of go (>=1.9.2).  Older versions may work so give it a shot.
* Use a GOPATH.  If you're doing something weird, it's likely that you will break exile.
* Install via `go get github.com/dan-compton/exile`

### Generating The Examples

To generate the examples, `cd` to the examples directory and run `exile` as shown below.

```
user@host> cd $GOPATH/src/github.com/dan-compton/exile
user@host> exile 
RENDERING TEMPLATE => /home/dc/work/src/github.com/dan-compton/exile/examples/main.go
DONE
user@host>
```

### Template Helpers

For template-metaprogramming to be accessible and language-agnostic, template helpers must be provided for source code transformation.  Go provides some of these functions by default as well as a way to register additional functions in the form of [template.FuncMap](https://golang.org/src/text/template/funcs.go?s=1000:1035#L20).

**TODO** Add additional details regarding the interfaces and helpers exported for pluggable template helpers.

## Motivation

Exile's core argument and ultimate goal is to demonstrate that code-generation is tool that can be leveraged not only to reduce the cost of duplicated effort, but also to:

* **Enforce standards** across a code-base through the [principle of least effort](https://en.wikipedia.org/wiki/Principle_of_least_effort).
* **Reduce time spent** on discussions around implementation details and choices of standard libraries by moving those discussions to the template-level.
* **Codify and convey knowledge** around how a team thinks a solution should look.
* Allow technical decision-owners to **hide behind-the-scenes changes to standard libraries**
* **Encourage a team to define a set of core-competencies** through templating solutions to problems they solve frequently.

## Contributing

Open a PR. 

**TODO** Code and Commit style information.

## Authors

* **Dan Compton** - *Initial work* - [dan-compton](https://github.com/dan-compton)

## License

This project is licensed under the BSD 3-clause "New" or "Revised" License - see the [LICENSE.md](LICENSE.md) file for details
