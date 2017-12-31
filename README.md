# Espárrago - simplifying R and Go integration

Espárrago means asparagus. Espá**r**ra**go** has R and Go in it.
Espárrago helps connect R (rstats) and Go (golang).

## Overview

[Romain François](http://romain.rbind.io/) has shown, in some lovely
blog posts and GitHub repositories (see [resources section](#resources))
that it's possible to call Go functions from R.

The process requires writing a little bit of C code to bridge the two
together. Espárrago writes that C code, so you don't have to.


<div id="resources"></div>
## Resources

From Romain François:

* [Calling Go from R](https://romain.rbind.io/blog/2017/05/14/calling-go-from-r/)
* Using Go Strings in R [[blog post](https://romain.rbind.io/blog/2017/06/10/using-go-strings-in-r/)] [[GitHub Repo](https://github.com/rstats-go/_playground_string)]
* [Using Go Slices in R](https://github.com/rstats-go/_playground_slice)

From others:

* [Building shared libraries in Go: Part I](https://www.darkcoding.net/software/building-shared-libraries-in-go-part-1/)
* [Building shared libraries in Go: Part II](https://www.darkcoding.net/software/building-shared-libraries-in-go-part-2/)
