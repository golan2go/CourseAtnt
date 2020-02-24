Graphs
We have files where every line is a message sent from one user to another.

```
1 -> 2
7 -> 3
8 -> 9
```

Write a program that will go over all the files and will print the user who got most messages.

1. Extract the zip file to a directory
1. Think about error handling (bad line …)
1. Make sure to close a file once you’re done with it
1. Use `filepath.Glob` to get the list of files
1. Use `bufio.Scanner` to iterate over lines
1. Use `strings.Split` and `strings.TrimSpace` to parse a line