#' @useDynLib goslice
#' @export
sum_int <- function(x) {
  stopifnot(is.numeric(x))
  .Call("sum_int", as.integer(x), PACKAGE = "goslice")
}

#' @export
sum_double <- function(x) {
   stopifnot(is.numeric(x))
  .Call("sum_double", as.double(x), PACKAGE = "goslice")
}

#' @export
numbers <- function(n){
  stopifnot(is.numeric(n) && length(n) == 1)
  .Call("numbers", as.integer(n), PACKAGE = "goslice")
}

#' @export
foobar <- function() {
  .Call("foobar", PACKAGE = "goslice")
}

#' @export
nbytes <- function(s){
  stopifnot(length(s) == 1)
  .Call("nbytes", as.character(s), PACKAGE = "goslice")
}

#' Doubles an integer using go
#'
#' @param x an integer
#' @useDynLib gotest
#' @export
godouble <- function(x) {
  stopifnot(is.numeric(x) && length(x)==1)
  .Call("godouble", as.integer(x), PACKAGE = "goslice")
}
