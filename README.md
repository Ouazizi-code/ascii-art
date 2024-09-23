# ASCII Art Generator

This Go application converts a given string into ASCII art using character representations stored in a text file. It reads the standard ASCII art from a file called `standard.txt` and prints the corresponding ASCII art representation of the input string.

## Features

- Converts printable ASCII characters (from space to tilde) into their ASCII art           representations.
- Handles multiple lines of input.
- Outputs an error message for non-printable characters.

## Requirements

- Go 1.16 or higher
- A `standard.txt` file containing ASCII art representations of characters.

## Installation

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd ascii-art

2. Ensure you have a standard.txt file in the project root with the ASCII art definitions.

## Usage

Run the program with the desired string as a command-line argument:
`go run main.go "Your string here"`

## Note

The program checks for non-printable characters and will terminate with an error message if any are found.
It will also inform you if the incorrect number of command-line arguments is provided.

## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, feel free to open an issue or submit a pull request.
