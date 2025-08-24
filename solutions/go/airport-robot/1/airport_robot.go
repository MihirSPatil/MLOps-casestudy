package airportrobot
import "fmt"

type Greeter interface {
    LanguageName() string
    Greet(username string) string
}

type Italian struct {}
type Portuguese struct{}

func (iG Italian) LanguageName() string{return "Italian"}
func (iG Italian) Greet (username string) string {return fmt.Sprintf("Ciao %s!", username)}

func (pG Portuguese) LanguageName() string {return "Portuguese"}
func (pG Portuguese) Greet (username string) string {return fmt.Sprintf("Olá %s!", username)}

func SayHello(user string, g Greeter) string{
    greet_string := g.Greet(user)
    lang_string := g.LanguageName()
    return fmt.Sprintf("I can speak %s: %s", lang_string, greet_string)
    }
    
// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
