package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/TeamMomentum/bs-url-normalizer/lib/urls"
)

var (
	version         = ""
	showVersion     = flag.Bool("version", false, "show version")
	crawlingURL     = flag.Bool("crawling-url", false, "convert to crawling URL")
	secondNormalize = flag.Bool("second", false, "convert to 2nd normalized URL")
	urlEscape       = flag.Bool("escape", false, "escape output URL")
	file            = flag.String("file", "", "input a file of URLs list")
)

func main() {
	flag.Parse()

	if *showVersion {
		println("bs-url-normalizer", version)
		os.Exit(0)
	}

	var in io.Reader
	if len(*file) > 0 {
		f, err := os.Open(*file)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		defer func() {
			if err := f.Close(); err != nil {
				println(err.Error())
			}
		}()

		in = f
	} else if len(flag.Args()) > 0 {
		in = strings.NewReader(flag.Arg(0))
	} else {
		in = os.Stdin
	}

	err := output(os.Stdout, in, *crawlingURL, *secondNormalize, *urlEscape)
	if err != nil {
		println(err.Error())
		os.Exit(2)
	}
}

func output(out io.Writer, in io.Reader, crawling, secondNorm, escape bool) error {
	s := bufio.NewScanner(in)
	for s.Scan() {
		uri, err := convert(s.Text(), crawling, secondNorm, escape)
		if err != nil {
			return err
		}

		if _, err := fmt.Fprintf(out, "%s\n", uri); err != nil {
			return err
		}
	}

	return s.Err()
}

func convert(uri string, crawling, secondNorm, escape bool) (string, error) {
	if len(uri) == 0 {
		return "", errors.New("URL is empty")
	}

	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}

	n, err := urls.NewNormalizer(u)
	if err != nil {
		return "", err
	}

	var convertedURI string
	switch {
	case crawling:
		convertedURI = n.CrawlingURL()
	case secondNorm:
		convertedURI = n.SecondNormalizedURL()
	default:
		convertedURI = n.FirstNormalizedURL()
	}

	if escape {
		convertedURI = url.QueryEscape(convertedURI)
	}

	return convertedURI, nil
}
