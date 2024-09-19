# Randonamer

Randonamer is a versatile coolname generator with support for multiple languages and custom configuration files. Born out of the need for a reliable random name generator for terminal sessions, this CLI tool leverages concepts from Formal Grammar and Automata theory to produce unique and interesting names.

## Features

- Generate random cool names based on predefined or custom grammar rules
- Support for multiple languages (currently English and Brazilian Portuguese)
- Customizable configuration
- Cross-platform compatibility (Linux, macOS, Windows)

## Installation

Currently, Randonamer is not available on package managers. You can build it from source or download pre-built binaries from the [releases page](https://github.com/pietrobondioli/randonamer/releases).

To install Randonamer locally:

1. Clone the repository:
   ```
   git clone https://github.com/pietrobondioli/randonamer
   ```
2. Navigate to the project directory and run:
   ```
   go install .
   ```

This will compile and install the `randonamer` binary in your Go bin directory.

Future plans include publishing Randonamer on various package managers, including:

- Linux (with a focus on AUR)
- macOS
- Windows

## Usage

Basic usage:

```
randonamer
```

This will generate a cool name using the default configuration.

### Command-line Options

Randonamer supports the following command-line flags:

- `-c, --cfg-file <path>`: Specify a custom configuration file path
- `-D, --DEBUG`: Enable debug logging
- `-l, --language <lang>`: Set the language for name generation (e.g., "en" or "pt_br")
- `-d, --data-path <path>`: Set the path to the data directory
- `-g, --grammar-file <file>`: Specify the grammar file to use
- `-s, --start-point <point>`: Set the starting point for the grammar

Example:

```
randonamer -l pt_br -d /custom/data/path -g custom_grammar
```

This command generates a name using Brazilian Portuguese language rules, with a custom data path and grammar file.

## Configuration

Randonamer uses a YAML configuration file. The default configuration is located at:

- Linux: `$HOME/.config/randonamer/config.yaml`
- macOS: `$HOME/Library/Application Support/randonamer/config.yaml`
- Windows: `%APPDATA%\randonamer\config.yaml`

You can override the default configuration by creating a custom YAML file and specifying it with the `-c` flag.

Example configuration:

```yaml
language: en
data-path: $HOME/.config/randonamer/data
grammar-file: _grammar
start-point: "start"
```

## How It Works

Randonamer uses a combination of configuration options and grammar files to generate cool names. Here's how the main components interact:

### Configuration Options

- `language`: Specifies the language subset to use (e.g., "en" for English, "pt_br" for Brazilian Portuguese).
- `data-path`: The base directory where language-specific data is stored.
- `grammar-file`: The name of the grammar file to use (default is "\_grammar").

### File Structure

The generator uses the following path to locate the grammar file:

```go
realDataPath := filepath.Join(cfg.DataPath, cfg.Language)
grammarPath := filepath.Join(realDataPath, cfg.GrammarFile)
```

For example, if your configuration is:

```yaml
language: pt_br
data-path: /home/user/.config/randonamer/data
grammar-file: _grammar
```

The generator will look for the grammar file at:
`/home/user/.config/randonamer/data/pt_br/_grammar`

### Grammar Structure

The grammar is defined in JSON format and consists of rules. Each rule can be either a "terminal" (directly generates words) or a "non-terminal" (refers to other rules).

Here's a simplified example of a grammar structure:

```json
{
  "name_separator": "-",
  "rules": {
    "start": {
      "type": "non-terminal",
      "generates": [["1"], ["2"], ["3"]]
    },
    "1": {
      "type": "non-terminal",
      "generates": [["animal"], ["profissao"], ["objeto"]]
    },
    "animal": {
      "type": "terminal",
      "generates": [["gato"], ["cachorro"], ["elefante"]]
    }
  }
}
```

### Creating Custom Grammars

To create a custom grammar:

1. Create a new directory under your `data-path` for your language (e.g., `custom_lang`).
2. Create a `_grammar.json` file in this directory with your main grammar structure.
3. Create additional JSON files for specific categories (e.g., `animal.json`, `adjetivo_composto.json`).

Example of `adjetivo_composto.json`:

```json
{
  "type": "terminal",
  "generates": [
    ["tigre-dentes-de-sabre"],
    ["ar-condicionado"],
    ["pronta-entrega"]
  ]
}
```

4. In your `_grammar.json`, you can reference these files as non-terminal rules.

To use your custom grammar, update your configuration to point to the new language directory and use the appropriate command-line flags:

```
randonamer -l custom_lang -d /path/to/custom/data
```

This flexible structure allows you to create complex, multi-layered grammars for generating diverse and interesting names.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by the need for a reliable coolname generator for terminal sessions
- Built using concepts from Formal Grammar and Automata theory
