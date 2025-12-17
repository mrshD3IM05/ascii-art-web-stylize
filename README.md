# ASCII Art Web

A lightweight, high-performance web application built with Go that transforms plain text into stylized ASCII art. Choose from multiple banner styles and generate beautiful ASCII representations instantly through a clean, browser-based interface.

![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-green)

## ✨ Features

- **Three Banner Styles**: Standard, Shadow, and Thinkertoy fonts
- **Real-time Generation**: Instant ASCII art rendering
- **Input Validation**: Robust error handling for invalid characters and edge cases
- **Clean UI**: Intuitive, responsive web interface
- **Modular Architecture**: Well-organized codebase for easy maintenance and extension
- **Performance Optimized**: Built on Go's efficient concurrency model

## 📋 Table of Contents

- [Demo](#-demo)
- [Installation](#-installation)
- [Usage](#-usage)
- [Project Structure](#-project-structure)
- [How It Works](#-how-it-works)
- [API Reference](#-api-reference)
- [Contributing](#-contributing)
- [Authors](#-authors)
- [License](#-license)

## 🎯 Demo

```
Input: "Hello"
Banner: standard

Output:
 _    _          _   _        
| |  | |        | | | |       
| |__| |   ___  | | | |   ___ 
|  __  |  / _ \ | | | |  / _ \
| |  | | |  __/ | | | | | (_) |
|_|  |_|  \___| |_| |_|  \___/
```

## 🚀 Installation

### Prerequisites

- **Go 1.18 or later** - [Download Go](https://golang.org/dl/)
- **Git** (optional) - For cloning the repository
- **Modern web browser** - Chrome, Firefox, Safari, or Edge

### Setup

1. **Clone the repository**
   ```bash
   git clone https://learn.zone01oujda.ma/git/ahaddou/ascii-art-web.git
   cd ascii-art-web
   ```

2. **Verify directory structure**
   ```
   ascii-art-web/
   ├── banners/
   │   ├── standard.txt
   │   ├── shadow.txt
   │   └── thinkertoy.txt
   ├── templates/
   │   ├── index.html
   │   ├── result.html
   │   ├── error.html
   │   └── style.css
   ├── server/
   │   ├── ascii/
   │   └── handlers/
   └── main.go
   ```

3. **Build the application**
   ```bash
   go build -o ascii-art-web
   ```

4. **Run the server**
   ```bash
   ./ascii-art-web
   ```
   
   Or run directly without building:
   ```bash
   go run main.go
   ```

5. **Access the application**
   
   Open your browser and navigate to:
   ```
   http://localhost:8080
   ```

## 💡 Usage

### Web Interface

1. Navigate to `http://localhost:8080`
2. Enter your text in the input field
3. Select a banner style (Standard, Shadow, or Thinkertoy)
4. Click **Generate**
5. View your ASCII art instantly

### Input Guidelines

- **Supported Characters**: ASCII printable characters (32-126)
- **Special Characters**: Newlines supported for multi-line output
- **Character Limit**: Maximum 2000 characters
- **Line Breaks**: Use Enter/Return for multi-line ASCII art

### Example Usage

**Single Line:**
```
Input: "Go!"
Banner: shadow
```

**Multi-line:**
```
Input: "Hello\nWorld"
Banner: thinkertoy
```

## 📁 Project Structure

```
ascii-art-web/
│
├── main.go                      # Application entry point
├── go.mod                       # Go module definition
│
├── banners/                     # ASCII art font files
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
│
├── server/
│   ├── ascii/
│   │   └── run.go              # ASCII generation logic
│   └── handlers/
│       ├── handlers.go          # Route registration
│       ├── homehandler.go       # Home page handler
│       ├── asciihandler.go      # ASCII generation handler
│       ├── errorhandler.go      # Error page handler
│       └── tempParser.go        # Template parser & validator
│
└── templates/                   # HTML templates & styles
    ├── index.html              # Main input form
    ├── result.html             # ASCII output display
    ├── error.html              # Error page
    └── style.css               # Application styles
```

## 🔧 How It Works

### Complete Request Flow (Chronological Order)

#### Phase 1: Application Startup

**Step 1: Template Loading and Validation** (`handlers/tempParser.go`)
- Parse all HTML templates: `index.html`, `result.html`, `error.html`
- Validate each template by executing with test data
- Pre-compile templates into memory for performance
- Store validated templates in `HandlersStruct`
- Exit gracefully if any template is invalid

**Step 2: Route Registration** (`handlers/handlers.go`)
- Register route handlers:
  - `GET /` → HomeHandler (serves input form)
  - `POST /ascii-art` → AsciiHandler (generates ASCII art)
  - `/static/*` → Static file server (CSS, assets)
- Routes are now configured and ready

**Step 3: Server Initialization** (`main.go`)
- Application calls `handlers.Handlers()` (Steps 1-2 complete)
- Server binds to port 8080
- Listens for incoming HTTP requests
- Server is now running and ready to accept connections

---

#### Phase 2: User Request Processing

**Step 4: Home Page Request** (`handlers/homehandler.go`)

When user visits `http://localhost:8080`:

1. Validate request method is GET
2. Verify URL path is exactly `/`
3. Execute `index.html` template
4. Send rendered HTML form to browser
5. User sees input form with text area and banner selection

**Step 5: Form Submission** (`handlers/asciihandler.go`)

When user clicks "Generate":

1. **Request Validation**
   - Verify HTTP method is POST
   - Extract form values: `text` and `banner`
   - Check neither field is empty
   - If validation fails → render 400 error page

2. **Call ASCII Generator**
   - Pass `text` and `banner` to `ascii.Run()`
   - Receive ASCII result and HTTP status code

3. **Response Handling**
   - If status is not 200 → render error page with status code
   - If successful → escape HTML characters for safety
   - Execute `result.html` template with ASCII output
   - Buffer the output
   - Send buffered response to browser

---

#### Phase 3: ASCII Art Generation

**Step 6: Input Processing** (`server/ascii/run.go`)

The core algorithm executes in this exact order:

1. **Empty Input Check**
   ```
   If input is empty → return status 400
   ```

2. **Line Break Normalization**
   ```
   Replace all "\r\n" (Windows) with "\n" (Unix)
   This ensures consistent newline handling across platforms
   ```

3. **Length Validation**
   ```
   If input length > 2000 characters → return status 400
   ```

4. **Banner Validation**
   ```
   Check if banner is one of: "standard", "shadow", "thinkertoy"
   If not → return status 404
   ```

5. **Banner File Loading**
   ```
   Construct file path: "banners/{banner}.txt"
   Read entire file into memory
   If file not found → return status 404
   ```

6. **Banner File Parsing**
   ```
   Remove all "\r" characters from file content
   Split file content by "\n" into array of lines
   Result: fontLines array containing all ASCII art rows
   ```

**Step 7: Banner File Structure**

Understanding the banner format:
```
- 95 printable ASCII characters (32 to 126)
- Each character = 8 lines of ASCII art + 1 blank line
- Total: 95 characters × 9 lines = 855 lines

Character mapping formula:
index = (character_code - 32) × 9 + row_number

Example for 'A' (ASCII 65):
- Base position: (65 - 32) × 9 = 297
- Row 1 is at index: 297 + 1 = 298
- Row 2 is at index: 297 + 2 = 299
- ... and so on through row 8
```

**Step 8: Special Case - Only Newlines**
```
If input contains only newline characters:
  Return the input as-is with status 200
  (Renders blank lines in output)
```

**Step 9: Main Rendering Loop**

```
Split input text by "\n" into array of lines

Initialize empty result string

For each line in input:
  
  If line is empty:
    Append single "\n" to result
    Continue to next line
  
  For row 1 through 8:  // Each character is 8 rows tall
    
    For each character in line:
      
      1. Character Validation:
         If character < 32 OR character > 126:
           Return status 400 (unsupported character)
      
      2. Calculate Index:
         index = (character - 32) × 9 + row
      
      3. Bounds Check:
         If index out of fontLines array bounds:
           Return status 500 (internal error)
      
      4. Append ASCII Row:
         Add fontLines[index] to result
    
    Append "\n" after completing one row across all characters
  
  // After 8 rows, one complete line of ASCII art is generated

Return final result with status 200
```

**Step 10: Visual Example**

For input "Hi" with standard banner:

```
Row 1: H_row1 + i_row1 → " _    _   _ "
Row 2: H_row2 + i_row2 → "| |  | | (_)"
Row 3: H_row3 + i_row3 → "| |__| |  _ "
Row 4: H_row4 + i_row4 → "|  __  | | |"
Row 5: H_row5 + i_row5 → "| |  | | | |"
Row 6: H_row6 + i_row6 → "|_|  |_| |_|"
Row 7: H_row7 + i_row7 → "            "
Row 8: H_row8 + i_row8 → "            "
```

**Step 11: Multi-line Processing**

For input "Hello\nWorld":

```
Process "Hello":
  → Generate 8 rows of ASCII art
  
Process "World":
  → Generate 8 rows of ASCII art
  
Final output: Two blocks of ASCII art separated by newlines
```

---

#### Phase 4: Response Delivery

**Step 12: HTML Rendering** (`handlers/asciihandler.go`)

1. Escape special HTML characters in result (`<`, `>`, `&`, etc.)
2. Create `PageData` struct with escaped result
3. Execute `result.html` template
4. Wrap output in `<pre>` tag (preserves formatting and spacing)
5. Buffer complete HTML page
6. Write buffer to HTTP response
7. Browser displays ASCII art with monospace font

**Step 13: Error Handling** (if errors occur at any step)

Error rendering process (`handlers/errorhandler.go`):

1. Select appropriate HTTP status code:
   - 400: Bad Request (invalid input, empty fields)
   - 404: Not Found (invalid banner, missing file)
   - 500: Internal Server Error (template errors, unexpected issues)

2. Create error data struct with status and message
3. Execute `error.html` template
4. Set HTTP response status code
5. Send error page to browser
6. User sees formatted error with "Home" link

---

### Performance Optimizations

**Memory Management:**
- Templates pre-compiled at startup (not per-request)
- Banner files read once per request (not cached globally)
- String concatenation using Go's efficient string builder
- Minimal allocations in hot paths

**Request Handling:**
- Input validation before expensive operations
- Early returns on errors (fail fast)
- Buffered response writing (single write operation)
- Static file serving via Go's optimized file server

**Concurrency:**
- Each request handled in separate goroutine (Go's default)
- No shared mutable state between requests
- Thread-safe template execution
- Scalable to handle multiple concurrent users

## 📡 API Reference

### POST `/ascii-art`

Generates ASCII art from input text.

**Request Parameters:**
- `text` (string, required): Input text to convert
- `banner` (string, required): Banner style (`standard`, `shadow`, or `thinkertoy`)

**Response:**
- Success: HTML page with ASCII art in `<pre>` tag
- Error: HTML error page with status code and message

**Status Codes:**
- `200 OK`: Successful generation
- `400 Bad Request`: Invalid input
- `404 Not Found`: Banner not found
- `500 Internal Server Error`: Server error

## 🤝 Contributing

We welcome contributions! Here's how you can help:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go conventions and best practices
- Add tests for new features
- Update documentation as needed
- Ensure all existing tests pass

## 👥 Authors

- **ahaddou** - Handlers and formatting
  - [GitHub](https://github.com/mrshD3IM05)
  - [Gitea](https://learn.zone01oujda.ma/git/ahaddou)

- **achent** - Front-end development
  - [GitHub](https://github.com/chentaymane)
  - [Gitea](https://learn.zone01oujda.ma/git/achent)

- **halhyane** - ASCII generator
  - [GitHub](https://github.com/Houssam-Alhyane)
  - [Gitea](https://learn.zone01oujda.ma/git/halhyane)

## 📄 License

This project is part of the Zone01 Oujda curriculum.

## 🙏 Acknowledgments

- Zone01 Oujda for project guidance
- The Go community for excellent documentation
- ASCII art enthusiasts worldwide

## 🔮 Future Enhancements

Potential features for future versions:

- [ ] Additional banner styles
- [ ] Export to file (PNG, TXT)
- [ ] RESTful API endpoint
- [ ] Color support
- [ ] Custom font upload
- [ ] ASCII art gallery
- [ ] Rate limiting
- [ ] User sessions

---

**Made with ❤️ at Zone01 Oujda**
