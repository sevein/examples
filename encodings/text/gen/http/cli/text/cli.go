// Code generated by goa v3.7.2, DO NOT EDIT.
//
// text HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/encodings/text/design -o encodings/text

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	textc "goa.design/examples/encodings/text/gen/http/text/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `text (concatstrings|concatbytes|concatstringfield|concatbytesfield)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` text concatstrings --a "Suscipit et quae soluta reiciendis quo totam." --b "Numquam aut tempora itaque nulla quo."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		textFlags = flag.NewFlagSet("text", flag.ContinueOnError)

		textConcatstringsFlags = flag.NewFlagSet("concatstrings", flag.ExitOnError)
		textConcatstringsAFlag = textConcatstringsFlags.String("a", "REQUIRED", "Left operand")
		textConcatstringsBFlag = textConcatstringsFlags.String("b", "REQUIRED", "Right operand")

		textConcatbytesFlags = flag.NewFlagSet("concatbytes", flag.ExitOnError)
		textConcatbytesAFlag = textConcatbytesFlags.String("a", "REQUIRED", "Left operand")
		textConcatbytesBFlag = textConcatbytesFlags.String("b", "REQUIRED", "Right operand")

		textConcatstringfieldFlags = flag.NewFlagSet("concatstringfield", flag.ExitOnError)
		textConcatstringfieldAFlag = textConcatstringfieldFlags.String("a", "REQUIRED", "Left operand")
		textConcatstringfieldBFlag = textConcatstringfieldFlags.String("b", "REQUIRED", "Right operand")

		textConcatbytesfieldFlags = flag.NewFlagSet("concatbytesfield", flag.ExitOnError)
		textConcatbytesfieldAFlag = textConcatbytesfieldFlags.String("a", "REQUIRED", "Left operand")
		textConcatbytesfieldBFlag = textConcatbytesfieldFlags.String("b", "REQUIRED", "Right operand")
	)
	textFlags.Usage = textUsage
	textConcatstringsFlags.Usage = textConcatstringsUsage
	textConcatbytesFlags.Usage = textConcatbytesUsage
	textConcatstringfieldFlags.Usage = textConcatstringfieldUsage
	textConcatbytesfieldFlags.Usage = textConcatbytesfieldUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "text":
			svcf = textFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "text":
			switch epn {
			case "concatstrings":
				epf = textConcatstringsFlags

			case "concatbytes":
				epf = textConcatbytesFlags

			case "concatstringfield":
				epf = textConcatstringfieldFlags

			case "concatbytesfield":
				epf = textConcatbytesfieldFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "text":
			c := textc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "concatstrings":
				endpoint = c.Concatstrings()
				data, err = textc.BuildConcatstringsPayload(*textConcatstringsAFlag, *textConcatstringsBFlag)
			case "concatbytes":
				endpoint = c.Concatbytes()
				data, err = textc.BuildConcatbytesPayload(*textConcatbytesAFlag, *textConcatbytesBFlag)
			case "concatstringfield":
				endpoint = c.Concatstringfield()
				data, err = textc.BuildConcatstringfieldPayload(*textConcatstringfieldAFlag, *textConcatstringfieldBFlag)
			case "concatbytesfield":
				endpoint = c.Concatbytesfield()
				data, err = textc.BuildConcatbytesfieldPayload(*textConcatbytesfieldAFlag, *textConcatbytesfieldBFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// textUsage displays the usage of the text command and its subcommands.
func textUsage() {
	fmt.Fprintf(os.Stderr, `The text service performs operations on strings
Usage:
    %[1]s [globalflags] text COMMAND [flags]

COMMAND:
    concatstrings: Concatstrings implements concatstrings.
    concatbytes: Concatbytes implements concatbytes.
    concatstringfield: Concatstringfield implements concatstringfield.
    concatbytesfield: Concatbytesfield implements concatbytesfield.

Additional help:
    %[1]s text COMMAND --help
`, os.Args[0])
}
func textConcatstringsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] text concatstrings -a STRING -b STRING

Concatstrings implements concatstrings.
    -a STRING: Left operand
    -b STRING: Right operand

Example:
    %[1]s text concatstrings --a "Suscipit et quae soluta reiciendis quo totam." --b "Numquam aut tempora itaque nulla quo."
`, os.Args[0])
}

func textConcatbytesUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] text concatbytes -a STRING -b STRING

Concatbytes implements concatbytes.
    -a STRING: Left operand
    -b STRING: Right operand

Example:
    %[1]s text concatbytes --a "Qui ut voluptatum culpa." --b "Dicta vel quo."
`, os.Args[0])
}

func textConcatstringfieldUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] text concatstringfield -a STRING -b STRING

Concatstringfield implements concatstringfield.
    -a STRING: Left operand
    -b STRING: Right operand

Example:
    %[1]s text concatstringfield --a "Magnam quis explicabo magnam eius et." --b "Culpa dolor eligendi recusandae aliquam."
`, os.Args[0])
}

func textConcatbytesfieldUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] text concatbytesfield -a STRING -b STRING

Concatbytesfield implements concatbytesfield.
    -a STRING: Left operand
    -b STRING: Right operand

Example:
    %[1]s text concatbytesfield --a "Quas nobis illum accusamus repudiandae." --b "Est voluptatem omnis sapiente nihil molestiae."
`, os.Args[0])
}
