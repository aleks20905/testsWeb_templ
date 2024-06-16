## Test Preparation Platform                

![screenshot](./assets/repoImages/img1.bmp)



## Set-Up

Create a new file called **`.env`** in the main folder that contains:
```bash 
GEMINI_API_KEY="Your_gemini_api_key"
```

Install Gemini API dependencies using:
```bash 
$ go get github.com/google/generative-ai-go
```

### Installing Dependencies

#### Linux and NixOS

For Linux and NixOS, use your package manager to install `templ` and `air`:

- **Linux**: Use your distribution's package manager (e.g., `apt`, `dnf`, `pacman`) to install `templ` and `air`.

- **NixOS**: Add `templ` and `air` to your environment:
  Edit your `configuration.nix` file to include `templ` and `air`:
     ```nix
     environment.systemPackages = with pkgs; [
       templ
       air
     ];
     ```

#### Windows

To install `templ`:
```bash 
$ go install github.com/a-h/templ/cmd/templ@latest
```

To install `air`:
```bash 
$ go install github.com/air-verse/air@latest
```

### Running the App

* It is recommended to use **`air`** for running the app. `Air` makes it easier and more convenient to use by avoiding the need to manually close, rebuild the project, and start it again.

#### Using `air`

After installing `air`, run your code with automatic rebuilding on file save:
```bash 
$ air
```

#### Without `air`

If you choose not to use `air`, you will need to manually build and run the project each time you make changes:

1. Build the Templ files:
   ```bash 
   $ templ generate
   ```

2. Run the main Go application:
   ```bash 
   $ go run main.go
   ```



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
   └─ userView
      └─ show.templ         # Template for user view
```
