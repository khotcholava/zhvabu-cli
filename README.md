# React CLI

A command-line tool for generating React TypeScript components with boilerplate code.

## Installation

### Install Script (Recommended)

**Unix/macOS/Linux:**

```bash
curl -fsSL https://raw.githubusercontent.com/khotcholava/zhvabu-cli/main/install.sh | sh
```

**Windows (PowerShell):**

```powershell
iwr -useb https://raw.githubusercontent.com/khotcholava/zhvabu-cli/main/install.ps1 | iex
```

### Homebrew (macOS/Linux)

```bash
brew tap khotcholava/homebrew-tap
brew install react-cli
```

### npm

```bash
npm install -g react-cli
```

### Manual Installation

1. Download the binary for your platform from [Releases](https://github.com/khotcholava/zhvabu-cli/releases)
2. Extract and move to a directory in your PATH
3. Make it executable (Unix): `chmod +x rc`

## Quick Start

1. **Initialize configuration** (optional):

   ```bash
   rc init
   ```

2. **Generate a component**:
   ```bash
   rc generate component UserList "userList, onActionClick"
   ```

This creates:

- `UserList/UserList.tsx` - Component file
- `UserList/UserList.css` - Style file
- `UserList/index.ts` - Export file

## Configuration

Create a `react-cli.json` file in your project root (or run `rc init`):

```json
{
  "defaults": {
    "component": {
      "style": "css",
      "path": ".",
      "skipStyle": false,
      "componentStyle": "functional"
    }
  },
  "project": {
    "prefix": ""
  }
}
```

## Usage

### Generate Component

```bash
rc generate component <ComponentName> ["<props>"]
```

**Options:**

- `--style <type>` - Style file type (css, scss, sass, none)
- `--path <path>` - Path where to create component
- `--prefix <prefix>` - Prefix for component name
- `--skip-style` - Skip style file creation

**Examples:**

```bash
# Basic component
rc generate component Button

# Component with props
rc generate component UserList "userList, onActionClick"

# With custom style and path
rc generate component Card --style scss --path src/components

# With prefix
rc generate component Modal --prefix App
```

### Initialize Config

```bash
rc init
```

Creates `react-cli.json` with default settings.

### Version

```bash
rc version
# or
rc --version
```

## Features

- ✅ Generate React TypeScript components
- ✅ Configurable style files (CSS, SCSS, SASS)
- ✅ Project-level configuration
- ✅ Props type generation
- ✅ Component folder structure
- ✅ Index file exports

## Development

### Build

```bash
./scripts/build.sh [version]
```

Builds binaries for all platforms in the `dist/` directory.

### Requirements

- Go 1.24 or higher

## License

MIT
