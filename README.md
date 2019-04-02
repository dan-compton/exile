## FUNK (pronounced FUNC) is a _simple_ code-generation framework built on go templates.  

[x]  Pluggable (plugins are collections of functions)
[x]  Simple interface for plugins to implement
[x]  Examples
[x]  Great at parties


## Funk's core argument and ultimate goal is to demonstrate that code-generation is tool that can be leveraged not only to reduce the cost of duplicated effort, but also to:

> Of course, this is the argument made by all metaprogramming languages -- like *jsonnet*.  Unlike jsonnet, Funk uses go-templates to provide a clear transformation and 
> dependency hierarchy with clear seperation between compile-time and runtime transformations.  Unlike C metaprogramming, no overhead is added to compilation outside of 
> plugin compilation -- nor are your source files littered with metastatements.  Note that this would be possible to implement in jsonnet, and the results would be catastrophic.

* **Enforce standards** across a code-base through the [principle of least effort](https://en.wikipedia.org/wiki/Principle_of_least_effort).
* **Reduce time spent** on discussions around implementation details and choices of standard libraries by moving those discussions to the template-level.
* **Codify and convey knowledge** around how a team thinks a solution should look.
* **Allow technical decision-owners to hide behind-the-scenes changes to standard libraries**
* **Encourage a team to define a set of core-competencies** through templating solutions to problems they solve frequently.

## Generating The Examples

A handful of examples are provided along with the repo.  These include string transformation functions like camel_to_snake and 
transformation functions for go packages like exported. You'll pick it up quickly -- it's easy. 

```
$ cd $GOPATH/src/github.com/dan-compton/Funk
$ Funk 
$ RENDERING TEMPLATE => /home/dc/work/src/github.com/dan-compton/Funk/examples/main.go
$ DONE
user@host>
```


## Interfaces to implement...

All plugged funcs must implement the Caller interface,
which simply defines the function's namesspace and possible
return values.  Since we're operating on strings, the retval will
always be a string.

```
type Caller interface {
	Call(...interface{}) (string, error)
	Namespace() string
}
```

You'll probably want to define your own namespace as well. 
You can do this with the mapper interface and `NewMappers(namespace string)` helper.
These allow you to bundle your callables (that implement Caller) into a single, referencable
namespace :)

A Mapper defines a mapping from a string transformation function to a template.FuncMap.
type Mapper interface {
	Map(t template.FuncMap)
}

## Questions/Concerns/Comments/Improvements

Simply open an issue.  I'm not picky on stylistic guidelines for issues
nor contributed code.  Just ensure that what you contribute conforms
to go's [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)!

If it does not, then I will not review it.  If you find violations in this codebase,
please open an issue.
