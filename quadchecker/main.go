package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	// ── Build subcommand ──────────────────────────────────────────────
	if len(os.Args) > 1 && os.Args[1] == "build" {
		runBuild()
		return
	}

	// ── Original quad-checker logic ───────────────────────────────────
	inputText := readStandardInput()

	if inputText == "" {
		fmt.Println("Not a quad function")
		return
	}

	lines := splitLines(inputText)
	height := len(lines)
	width := len(lines[0])

	if !isRectangle(lines, width) {
		fmt.Println("Not a quad function")
		return
	}

	matches := findMatchingQuads(inputText, width, height)

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	printMatches(matches)
}

/* ===================== BUILD SUBCOMMAND ===================== */

// runBuild replicates the shell script: extracts each quadX function,
// wraps it as a standalone main, compiles it, then builds quadchecker.
func runBuild() {
	source, err := os.ReadFile("main.go")
	if err != nil {
		fmt.Println("Error: main.go not found!")
		os.Exit(1)
	}

	fmt.Println("--- 🛠️  Extracting and Building Quads from main.go ---")

	for _, letter := range []string{"A", "B", "C", "D", "E"} {
		name := "quad" + letter
		fmt.Printf("Building %s...\n", name)

		body := extractFunc(string(source), name)
		if body == "" {
			fmt.Printf("  Warning: func %s not found, skipping\n", name)
			continue
		}

		// Replace "func quadX()" with "func main()" so it compiles as a binary.
		wrapped := "package main\n\nimport (\n\t\"fmt\"\n\t\"os\"\n\t\"strconv\"\n)\n\n" +
			strings.Replace(body, "func "+name+"()", "func main()", 1) + "\n"

		tmp := "temp_build_" + name + ".go"
		if err := os.WriteFile(tmp, []byte(wrapped), 0o644); err != nil {
			fmt.Printf("  Error writing temp file: %v\n", err)
			continue
		}

		if err := gobuild(name, tmp); err != nil {
			fmt.Printf("  Build failed for %s: %v\n", name, err)
		}

		os.Remove(tmp)
	}

	// Build the quadchecker itself.
	fmt.Println("Building quadchecker...")
	if err := gobuild("quadchecker", "main.go"); err != nil {
		fmt.Printf("  Build failed for quadchecker: %v\n", err)
	}

	fmt.Println("--- ✅ All executables are ready! ---")
}

// extractFunc pulls the complete source block of the named function
// using brace-depth counting, handling nested braces correctly.
func extractFunc(src, name string) string {
	sig := "func " + name + "() {"
	start := strings.Index(src, sig)
	if start == -1 {
		return ""
	}

	depth := 0
	i := start
	for i < len(src) {
		switch src[i] {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return src[start : i+1]
			}
		}
		i++
	}
	return ""
}

// gobuild runs `go build -o <out> <file>` and streams output to stdout/stderr.
func gobuild(out, file string) error {
	cmd := exec.Command("go", "build", "-o", out, file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

/* ===================== ORIGINAL QUAD FUNCTIONS ===================== */

func quadA() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	x, _ := strconv.Atoi(args[0])
	y, _ := strconv.Atoi(args[1])
	if x <= 0 || y <= 0 {
		return
	}
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (i == 1 || i == y) && (j == 1 || j == x) {
				fmt.Print("o")
			} else if i == 1 || i == y {
				fmt.Print("-")
			} else if j == 1 || j == x {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func quadB() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	x, _ := strconv.Atoi(args[0])
	y, _ := strconv.Atoi(args[1])
	if x <= 0 || y <= 0 {
		return
	}
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 {
				fmt.Print("/")
			} else if i == 1 && j == x {
				fmt.Print("\\")
			} else if i == y && j == 1 {
				fmt.Print("\\")
			} else if i == y && j == x {
				fmt.Print("/")
			} else if i == 1 || i == y || j == 1 || j == x {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func quadC() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	x, _ := strconv.Atoi(args[0])
	y, _ := strconv.Atoi(args[1])
	if x <= 0 || y <= 0 {
		return
	}
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && (j == 1 || j == x) {
				fmt.Print("A")
			} else if i == y && (j == 1 || j == x) {
				fmt.Print("C")
			} else if i == 1 || i == y || j == 1 || j == x {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func quadD() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	x, _ := strconv.Atoi(args[0])
	y, _ := strconv.Atoi(args[1])
	if x <= 0 || y <= 0 {
		return
	}
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if j == 1 && (i == 1 || i == y) {
				fmt.Print("A")
			} else if j == x && (i == 1 || i == y) {
				fmt.Print("C")
			} else if i == 1 || i == y || j == 1 || j == x {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func quadE() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	x, _ := strconv.Atoi(args[0])
	y, _ := strconv.Atoi(args[1])
	if x <= 0 || y <= 0 {
		return
	}
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (i == 1 && j == 1) || (i == y && j == x && i != 1 && j != 1) {
				fmt.Print("A")
			} else if (i == 1 && j == x) || (i == y && j == 1) {
				fmt.Print("C")
			} else if i == 1 || i == y || j == 1 || j == x {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

/* ===================== CHECKER HELPERS ===================== */

func readStandardInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) == 0 {
		return ""
	}

	return strings.Join(lines, "\n")
}

func normalize(s string) string {
	return strings.TrimRight(s, "\n")
}

func splitLines(input string) []string {
	return strings.Split(input, "\n")
}

func isRectangle(lines []string, width int) bool {
	for _, line := range lines {
		if len(line) != width {
			return false
		}
	}
	return true
}

func findMatchingQuads(input string, width, height int) []string {
	type quad struct {
		name string
		fn   func(int, int) string
	}

	quads := []quad{
		{"quadA", generateQuadA},
		{"quadB", generateQuadB},
		{"quadC", generateQuadC},
		{"quadD", generateQuadD},
		{"quadE", generateQuadE},
	}

	var matches []string

	for _, q := range quads {
		if normalize(input) == normalize(q.fn(width, height)) {
			match := fmt.Sprintf("[%s] [%d] [%d]", q.name, width, height)
			matches = append(matches, match)
		}
	}

	return matches
}

func printMatches(matches []string) {
	for i := 0; i < len(matches); i++ {
		fmt.Print(matches[i])
		if i != len(matches)-1 {
			fmt.Print(" || ")
		}
	}
	fmt.Println()
}

func buildQuadPattern(width, height int,
	tl, tr, bl, br, hChar, vChar string) string {

	if width <= 0 || height <= 0 {
		return ""
	}

	var result strings.Builder

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			switch {
			case row == 0 && col == 0:
				result.WriteString(tl)
			case row == 0 && col == width-1:
				result.WriteString(tr)
			case row == height-1 && col == 0:
				result.WriteString(bl)
			case row == height-1 && col == width-1:
				result.WriteString(br)
			case row == 0 || row == height-1:
				result.WriteString(hChar)
			case col == 0 || col == width-1:
				result.WriteString(vChar)
			default:
				result.WriteString(" ")
			}
		}
		if row != height-1 {
			result.WriteString("\n")
		}
	}

	return result.String()
}

/* ===================== GENERATORS ===================== */

func generateQuadA(w, h int) string {
	return buildQuadPattern(w, h, "o", "o", "o", "o", "-", "|")
}

func generateQuadB(w, h int) string {
	return buildQuadPattern(w, h, "/", "\\", "\\", "/", "*", "*")
}

func generateQuadC(w, h int) string {
	return buildQuadPattern(w, h, "A", "A", "C", "C", "B", "B")
}

func generateQuadD(w, h int) string {
	return buildQuadPattern(w, h, "A", "C", "A", "C", "B", "B")
}

func generateQuadE(w, h int) string {
	return buildQuadPattern(w, h, "A", "C", "C", "A", "B", "B")
}
