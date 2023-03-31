package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {
	// first thing we do is Create a new multiplexer (map)
	mux := http.NewServeMux()

	// declear handlers for each route
	mux.HandleFunc("/", home)
	mux.HandleFunc("/greeting", greeting)
	mux.HandleFunc("/random", random)

	// Create a new server and listen on port 3360
	fmt.Println("Starting server on port 3360...")
	http.ListenAndServe(":3360", mux)
}

// homeHandler sets the root route
func home(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
	<title>Home</title>
		<style>
		body {
		font-family: "Times New Roman", Times, cursive;
		font-size: 12px;
		line-height: 1.5;
		background-color: #4B0082;
		text-align: left;
		padding: 2em;
		}
		h1 
		{
		color:#FFA500;
		font-size: 3em;
		margin-bottom: 1em;
		}
		</style>
	</head>
	<body>
		<h1>What's Up! I am Nolan Lockwood as you probably remember. My hobbies are playing video games, watching anime and 
		watching series from genres like sci-fi, mystery and fantasy mostly. I have successfully managed to take a step from my introverted nature
		and now hang out with friends at school.</h1>			
		</body>
	</html>
	`
	fmt.Fprint(w, html)
}

// getGreeting is a switch that returns the appropriate greeting based on the time of day the system has
func getGreeting(t time.Time) string {
	switch t.Weekday() {
	case time.Monday:
		return "Monday!"
	case time.Tuesday:
		return "Tuesday!"
	case time.Wednesday:
		return "Wednesday!"
	case time.Thursday:
		return "Thursday!"
	case time.Friday:
		return "Friday!"
	case time.Saturday:
		return "Saturday!"
	case time.Sunday:
		return "Sunday!"
	default:
		return "unknownday"
	}
}

// greetingHandler sets up the /greeting route
func greeting(w http.ResponseWriter, r *http.Request) {
	current_time := time.Now()
	current_date := time.Now()

	greeting := getGreeting(current_time)
	html := `
		<!DOCTYPE html>
		<html>
			<head>
			<title>Greeting</title>
			<style>
			body {
			font-family: "Times New Roman", Times, cursive;
			font-size: 12px;
			line-height: 1.5;
			background-color: #4B0082;
			text-align: left;
			padding: 2em;
			}
			
			h1 {
				color:#FFA500;
				font-size: 3em;
				margin-bottom: 1em;
				}
			</style>
			</head>
		<body>
		<h1>Today's date is {{date}} and the time is {{time}}.</h1>
		<p>User enjoy the rest of your {{greeting}}.</p>
		</body>
	</html>
	`
	html = strings.ReplaceAll(html, "{{date}}", current_date.Format("2006-01-02"))
	html = strings.ReplaceAll(html, "{{time}}", current_time.Format("3:04pm"))
	html = strings.ReplaceAll(html, "{{greeting}}", greeting)
	fmt.Fprint(w, html)
}

// randomHandler sets up the /random route
func random(w http.ResponseWriter, r *http.Request) {
	quotes := map[int]string{
		1:  "Appear weak when you are strong, and strong when you are weak. ― Sun Tzu, The Art of War",
		2:  "One may know how to conquer without being able to do it. ― Sun Tzu, The Art of War",
		3:  "Most great people have attained their greatest success just one step beyond their greatest failure. ~ Napoleon Hill",
		4:  "Knowing yourself is the beginning of all wisdom. ― Aristotle",
		5:  "It is the mark of an educated mind to be able to entertain a thought without accepting it. ― Aristotle, Metaphysics",
		6:  "Hell hath no fury like a woman scorned. - William Congreve",
		7:  "Hold fast to dreams, For if dreams die Life is a broken-winged bird, That cannot fly.― Langston Hughes",
		8:  "Do not fear failure but rather fear not trying.― Roy T. Bennett",
		9:  "Technology is a tool, and like any tool, its value depends on how it is used. - David Crystal",
		10: "Life is never made unbearable by circumstances, but only by lack of meaning and purpose.― Victor Frankl",
	}

	//create a random number generator
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes)) + 1

	html := `
        <!DOCTYPE html>
        <html>
            <head>
            <title>Random Quote</title>
			<style>
			body {
			font-family: "Times New Roman", Times, cursive;
			font-size: 12px;
			line-height: 1.5;
			background-color: #4B0082;
			
			text-align: center;
			}
			h1 {
			margin-top: 50px;
			font-size: 28px;
			color: #FF500;
			text-align:center;
			}
			p {
			margin-top: 20px;
			font-size: 20px;
			}
			</style>
            </head>
        <body>
        <h1>Random Quote</h1>
        <p>` + quotes[randomIndex] + `</p>
        </body>
        </html>
    `
	fmt.Fprintf(w, html)
}
