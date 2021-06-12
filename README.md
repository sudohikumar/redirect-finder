### ABOUT

This is a small tool which tries to find the redirect URLs of any URLs.
To run, user can run 

```
$ go run main.go <set of URLs space separated>
```

For example,

```
$ go run main.go "https://bit.ly/3iBDOMM" "https://bit.ly/3wdOmWa" "https://bit.ly/3iBDOMM" "https://bit.ly/3wdOmWa"
```

This will give output as follows.

```
URL: https://bit.ly/3iBDOMM - Redirected URL: www.google.com
URL: https://bit.ly/3iBDOMM - Redirected URL: www.google.com
URL: https://bit.ly/3wdOmWa - Redirected URL: www.microsoft.com
URL: https://bit.ly/3wdOmWa - Redirected URL: www.microsoft.com
```