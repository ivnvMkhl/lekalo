# Lekalo - File and Project Structure Generator

Lekalo is a Go-based command-line utility for generating files and project structures using YAML templates with Jinja2-like syntax support.

## Features

>  * 🚀 File and folder generation from templates
>  * 📝 YAML template configuration
>  * 🔍 Automatic template discovery (global and local)
>  * ✨ Jinja2-style template syntax
>  * ⚡ Fast generation through native compilation

## Installation

### Build from Source

>  1. Ensure you have Go 1.21+ installed
>  2. Clone the repository:
>  ```bash
>  git clone https://github.com/ivnvMkhl/lekalo.git
>  cd lekalo
>  ```
>  3. Build the project:
>  ```bash
>  go build -o lekalo .
>  ```
>  4. (Optional) Install system-wide:
>  ```bash
>  sudo mv lekalo /usr/local/bin/
>  ```

### Download Binaries

Pre-built binaries for various platforms are available in [project releases](https://github.com/ivnvMkhl/lekalo/tree/master/build_bin).

## Usage

### Basic Commands

```bash
# List available templates
lekalo list

# Generate files from template
lekalo gen <template-name> [key=value...]

# Show help
lekalo --help
```

### React Component Example

1. Create a template file `.lekalo_templates.yml` placement on run folder:
```yaml
templates:
  react-component:
    params:
      - name: "name"
        prompt: "Enter component name"
      - name: "path"
        prompt: "Enter path"
        default: "./"
    files:
      component:
        path: "{{ path }}/{{ name }}.tsx"
        template: |
          import React from 'react';

          interface {{ name }}Props {}

          export const {{ name }}: React.FC<{{ name }}Props> = () => {
            return <div>{{ name }}</div>;
          }
      index:
        path: "{{ path }}/index.ts"
        template: |
          export { {{ name }} } from './{{ name }}'
```
2. Run generation:
```bash
lekalo gen react-component name=Button path=./src/components/Button
```

## Configuration Format

Lekalo uses YAML files for template definitions. Full structure example:

```yaml
templates:
  template-name:
    params:
      - name: "param1"          # Parameter name
        prompt: "Prompt text"   # Prompt text (optional)
        default: "default"      # Default value (optional)

    # Dynamically created folders (optional)
    folders:
      root: "./{{ param1 }}"    # Root folder
      components: "{{ folders.root }}/src"  # Subfolder

    # Files to generate
    files:
      file1:
        path: "{{ folders.components }}/{{ param1 }}.tsx"  # Output path
        template: |  # File content
          // Jinja2 template with parameter substitution
          export const {{ param1 }} = () => null;
```

## Template Locations

Lekalo searches for templates in this order:

  1. Local `.lekalo_templates.yml` in current directory
  2. Global `~/.lekalo/templates.yml`

Local templates take precedence over global ones.

## Development

### Project Structure

```bash
├── cmd/                # CLI commands
│   ├── gen.go          # File generation
│   ├── list.go         # Template listing
│   └── cmd.go          # Core CLI logic
├── config/             # Configuration
│   └── config.go       # YAML config loading
├── core/               # Core logic
│   └── generate.go     # File generation
├── render/             # Template rendering
│   ├── engine.go       # Jinja2 rendering
│   └── paths.go        # Path processing
├── npm_package/        # NPM package files
├── go.mod              # Dependencies
├── go.sum              # Dependencies
└── main.go             # Entry point
```

### Building

To build the project:

```bash
# Build for current platform
go build -o lekalo .

# Cross-compilation
./build.sh
```

The build.sh script creates binaries for:

* Linux (amd64, arm64)
* Windows (amd64)
* macOS (amd64, arm64)

## License

Lekalo is distributed under the [MIT License](https://github.com/ivnvMkhl/lekalo/blob/master/LICENCE).
