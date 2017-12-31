library(goslice)

print("sum_int")
vec <- c(1,2,3)
result <- sum_int(vec)
print(paste("got result", result))
stopifnot(result == 6L)

print("sum_double")
dub <- c(1.2, 3.4)
result <- sum_double(dub)
print(paste("got result", result))
stopifnot(result == (1.2 + 3.4))

print("numbers")
n <- 22
res <- numbers(n)
print("got result")
print(res)
stopifnot(res == c(0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42))

print("foobar")
res <- foobar()
print(paste("got result", res))
stopifnot(res == "foobar")

print("nbytes")
s <- "hello world"
res <- nbytes(s)
print(paste("got result", res))
stopifnot(res == nchar(s))

print("godouble")
x <- 42
res <- godouble(x)
print(paste("got result", res))
stopifnot(res == 84L)
