# Teppan

CLI tool to generate text from template with .env

## Usage

Template format is [Go's text/template package](https://golang.org/pkg/text/template/).

```bash
$ teppan [--base64] <template file>
```

You can try teppan as below:

```bash
$ cd $GOPATH/github.com/dtan4/teppan

$ cat sample.tmpl
Hello, my name is {{ .NAME }}.

$ cp .env.sample .env
$ vi .env
NAME=dtan4

$ teppan sample.tmpl
Hello, my name is dtan4.

# Embed Base64-encoded environment variables
$ ./teppan --base64 sample.tmpl
Hello, my name is ZHRhbjQ=.
```

## Author

Daisuke Fujita (@dtan4)
