![screenshot](./assets/repoImages/img1.bmp)

## Set-Up

Need to create new file called **``.env``** in the main folder that contains :
```bash 
GEMINI_API_KEY="Ur_gemini_api_key"
```

Installing Gemini api dependencies using:
```bash 
$ go get github.com/google/generative-ai-go
```
To install [Templ](https://github.com/a-h/templ)
```bash 
$ go install github.com/a-h/templ/cmd/templ@latest
```

* For running the app is recommended to use **``'air'``** but if u not gonna do any changes its not necessary to have air, its just more easy and convenient to use it than, manualy closeing, rebuilding the project and starting it again etc.


#### Using ``air`` [repo link](https://github.com/air-verse/air)

To install **``'air'``** just run this in the console or look up the repo for installation instructions
```bash 
$ go install github.com/air-verse/air@latest 
```
After that just using **``'air'``** it will run your code, while automatically running the code when saved and automatically building the .temp files
```bash 
$ air
```
#### Without ``air`` altenativ  #TODO

* need to build templ each time.
* run main.go and whit all other files
*
*


## Adding Questions 

Adding questions to the program is straightforward, although there is currently no way to create questions directly within the app. Questions can be added by modifying the files in the [assets](https://github.com/aleks20905/testsWeb_templ/tree/main/assets) directory.

Depending on the type of question you want to add, you can categorize them into two types: multiple-choice questions and open-ended questions. For multiple-choice questions with predefined answers, use the [questions.json](https://github.com/aleks20905/testsWeb_templ/blob/main/assets/questions.json) file. For open-ended questions, use the [open_questions.json](https://github.com/aleks20905/testsWeb_templ/blob/main/assets/open_questions.json) file.

The structure of the data is as follows:

### Multiple-Choice Questions
```go
type Question struct {
	ID       int      // ID is dynamically generated when the file is accessed
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   []string `json:"answer"`
}
```

### Open-Ended Questions
```go
type OpenQuestion struct {
	ID       int    // ID is dynamically generated when the file is accessed
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
```

## Project File Structure

```
├─ .air.toml                # Configuration file for the Air live reloading tool
├─ README.md                
├─ assets
│  ├─ open_questions.json   # JSON file containing open-ended questions
│  └─ questions.json        # JSON file containing multiple-choice questions
├─ css
│  └─ main.css              # Main stylesheet for the project
├─ go.mod                   # Go module file defining the module and dependencies
├─ go.sum                   # Go checksum file for dependencies
├─ handler
│  ├─ submit.go             # Handler for submitting data
│  ├─ user.go               # Handler for user-related operations
│  └─ utils.go              # Utility functions for handlers
├─ jsonthing
│  ├─ model.go              # Go struct definitions for JSON data
│  └─ readFile.go           # Functions for reading JSON files
├─ main.go                  # Main entry point for the Go application
├─ tmp
│  ├─ build-errors.log      # Log file for build errors
│  └─ main.exe              # Compiled executable of the project
└─ view
   ├─ base
   │  └─ base.templ         # Base template for the views
   ├─ components
   │  ├─ card
   │  │  └─ card.templ      # Template for the card component
   │  ├─ question
   │  │  ├─ openQuestion.templ # Template for open questions
   │  │  └─ question.templ  # Template for multiple-choice questions
   │  └─ submitquestion
   │     └─ submitquestion  # Template for the submit question component
   └─ userView
      └─ show.templ         # Template for user view
```

This should provide clear explanations for the purpose of each file in your project.