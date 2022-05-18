// Template generated code from trgen 0.16.4

package main
import (
    "fmt"
    "os"
    "io"
    "time"
    "net/http"
	"sync"
	"github.com/antlr/antlr4/runtime/Go/antlr"
    "example.com/myparser/parser"
    _ "net/http/pprof"
)

type CustomErrorListener struct {
    errors int
}

func NewCustomErrorListener() *CustomErrorListener {
    return new(CustomErrorListener)
}

func (l *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
    l.errors += 1
    antlr.ConsoleErrorListenerINSTANCE.SyntaxError(recognizer, offendingSymbol, line, column, msg, e)
}

func (l *CustomErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
    antlr.ConsoleErrorListenerINSTANCE.ReportAmbiguity(recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

func (l *CustomErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
    antlr.ConsoleErrorListenerINSTANCE.ReportAttemptingFullContext(recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

func (l *CustomErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
    antlr.ConsoleErrorListenerINSTANCE.ReportContextSensitivity(recognizer, dfa, startIndex, stopIndex, prediction, configs)
}

var cpuprofile = "cpuprofile.txt"
var memprofile = "memprofile"

func main() {
    var show_tree = false
    var show_tokens = false
    var file_name = ""
    var input = ""
    var str antlr.CharStream = nil
    for i := 0; i < len(os.Args); i = i + 1 {
        if os.Args[i] == "-tokens" {
            show_tokens = true
            continue
        } else if os.Args[i] == "-tree" {
            show_tree = true
            continue
        } else if os.Args[i] == "-input" {
            i = i + 1
            input = os.Args[i]
        } else if os.Args[i] == "-file" {
            i = i + 1
            file_name = os.Args[i]
        }
    }
    if input == "" && file_name == "" {
        var b []byte = make([]byte, 1)
        var st = ""
        for {
            _, err := os.Stdin.Read(b)
            if err == io.EOF {
                break
            }
            st = st + string(b)
        }
        str = antlr.NewInputStream(st)
    } else if input != "" {
        str = antlr.NewInputStream(input)
    } else if file_name != "" {
        str, _ = antlr.NewFileStream(file_name);        
    }
    var lexer = parser.NewSimpleBooleanLexer(str);
    if show_tokens {
        j := 0
        for {
            t := lexer.NextToken()
            fmt.Print(j)
            fmt.Print(" ")
            // missing fmt.Println(t.String())
            fmt.Println(t.GetText())
            if t.GetTokenType() == antlr.TokenEOF {
                break
            }
            j = j + 1
        }
        // missing lexer.Reset()
    }
    // Requires additional 0??
    var tokens = antlr.NewCommonTokenStream(lexer, 0)
    var parser = parser.NewSimpleBooleanParser(tokens)

    lexerErrors := &CustomErrorListener{}
    lexer.RemoveErrorListeners()
    lexer.AddErrorListener(lexerErrors)

    parserErrors := &CustomErrorListener{}
    parser.RemoveErrorListeners()
    parser.AddErrorListener(parserErrors)


	go func() {
        fmt.Println(http.ListenAndServe("localhost:6060", nil))
    }()

	var wg sync.WaitGroup
	wg.Add(1) // pprof - so we won't exit prematurely
    wg.Add(1) // for the hardWork
	defer wg.Done()

    // mutated name--not lowercase.
    start := time.Now()
    var tree = parser.Parse()
    elapsed := time.Since(start)
    fmt.Printf("Time: %.3f s", elapsed.Seconds())
    fmt.Println()
    if parserErrors.errors > 0 || lexerErrors.errors > 0 {
        fmt.Println("Parse failed.");
    } else {
        fmt.Println("Parse succeeded.")
    }
    if show_tree {
        ss := tree.ToStringTree(parser.RuleNames, parser)
        fmt.Println(ss)
    }

    wg.Wait()

    if parserErrors.errors > 0 || lexerErrors.errors > 0 {
        os.Exit(1)
    } else {
        os.Exit(0)
    }
}
