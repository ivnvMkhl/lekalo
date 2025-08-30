# Lekalo - File and Project Structure Generator

Lekalo is a Go-based command-line utility for generating files and project structures using YAML templates with Jinja2-like syntax support.

## Features

>  * 🚀 File and folder generation from templates
>  * 📝 YAML template configuration
>  * 🔍 Automatic template discovery (global and local)
>  * ✨ Jinja2-style template syntax
>  * ⚡ Fast generation through native compilation

## Installation

Local in project
```bash
npm i -D lekalo
```

Global in system
```bash
npm i -g lekalo
```

## Usage

### Basic Commands

```bash
# List available templates
npx lekalo list

# Generate files from template
npx lekalo gen <template-name> [key=value...]

# Show help
npx lekalo --help
```

### Localization

CLI automatically detects the language installed in the system. To manually redefine the language, set the system env variable `LEKALO_LANG=en`

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
        multiple: false         # Multiple mode (optional default false)
        default: "default"      # Default value (optional)
      - name: "param2"          # Parameter name
        prompt: "Prompt text"   # Prompt text (optional)
        multiple: true          # Multiple mode
        default:                # Default values (optional)
            - "default"         # Default value 1
            - "default2"        # Default value 2

    # Dynamically created folders (optional)
    folders:
      root: "./{{ param1 }}"    # Root folder
      components: "{{ folders.root }}/src"  # Subfolder

    # Files to generate
    files:
      file1:
        mode: "default" # Generate mode
        path: "{{ folders.components }}/{{ param1 }}.tsx"  # Output path
        template: |  # File content
          // Jinja2 template with parameter substitution
          export const {{ param1 }} = () => null;
```

### Generate mode

* `default` (or unset) - When the file exists, overwrite confirmation will be required
* `always_overwrite` - File will be overwritten (no confirmation)
* `skip_if_exist` - File generation will be skipped if the file exists
* `append` - The file will be appended to or created if it doesn't exist

## Template Locations

Lekalo searches for templates in this order:

  1. Local `.lekalo_templates.yml` in current directory
  2. Global `~/.lekalo/templates.yml`

Local templates take precedence over global ones.

## License

Lekalo is distributed under the [MIT License](https://github.com/ivnvMkhl/lekalo/blob/master/LICENCE).
