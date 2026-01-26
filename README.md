# Numerical Analysis Class Repository

This repository contains assignments for the numerical analysis class, where mathematical problems are solved through coding.

## Structure

Each assignment is located in its own directory (`program1`, `program2`, etc.). The code is written in Go and is intended to be run via the terminal. Each program directory includes a separate `README.md` with specific instructions and details.

### README locations:
- **[Program 1](program1/README.md)** - Bisection Method Programming Assignment

---

## Setup
### 1. Clone the Repository

To set up the project, clone the repo:

```bash
git clone https://github.com/rdYards/numerical-analysis.git
cd numerical-analysis
```

### 2. Install Go
If needed download and install go onto your system.

### Windows
1. **Download Go**
   - Visit the official Go downloads page: [https://go.dev/dl/](https://go.dev/dl/)
   - Download the Windows `.msi` installer for your architecture (e.g., amd64).

2. **Install Go**
   - Run the installer and follow the setup wizard.
   - By default, Go will be installed to `C:\Program Files\Go`.

3. **Set Environment Variables**
   - Ensure `C:\Program Files\Go\bin` is in your `PATH`.
   - Optionally, set your Go workspace:
     ```bash
     setx GOPATH "%USERPROFILE%\go"
     setx PATH "%PATH%;%GOPATH%\bin"
     ```

4. **Verify Installation**
   - Open Command Prompt or PowerShell and type:
     ```bash
     go version
     ```
   - You should see the installed Go version printed.

---

### macOS
1. **Download Go**
   - Go to [https://go.dev/dl/](https://go.dev/dl/) and download the macOS `.pkg` installer.

2. **Install Go**
   - Open the downloaded package and follow the instructions.
   - By default, Go installs to `/usr/local/go`.

3. **Update Environment Variables**
   - Add the following lines to your shell config (`~/.zshrc` or `~/.bash_profile`):
     ```bash
     export PATH=$PATH:/usr/local/go/bin
     export GOPATH=$HOME/go
     export PATH=$PATH:$GOPATH/bin
     ```
   - Then reload your shell:
     ```bash
     source ~/.zshrc
     ```

4. **Verify Installation**
   ```bash
   go version
   
---

### Linux
#### Ubuntu/Debian
1. **Update Package List**
   Open a terminal and run:
   ```bash
   sudo apt update
   ```

2. **Install Go**
   ```bash
   sudo apt install golang
   ```

3. **Set Environment Variables**
   Add the following to your shell config (`~/.bashrc`, `~/.zshrc`, or `~/.profile`):
   ```bash
   export GOPATH=$HOME/go
   export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
   ```
   Then reload your shell:
   ```bash
   source ~/.bashrc  # or ~/.zshrc, etc.
   ```

4. **Verify Installation**
   ```bash
   go version
   ```

---

#### Fedora
1. **Install Go**
   ```bash
   sudo dnf install golang
   ```

2. **Set Environment Variables**
   Add the following to your shell config:
   ```bash
   export GOPATH=$HOME/go
   export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
   ```
   Reload your shell:
   ```bash
   source ~/.bashrc  # or ~/.zshrc, etc.
   ```

3. **Verify Installation**
   ```bash
   go version
   ```

---

#### Arch Linux
1. **Install Go**
   ```bash
   sudo pacman -S go
   ```

2. **Set Environment Variables**
   Add the following to your shell config:
   ```bash
   export GOPATH=$HOME/go
   export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
   ```
   Reload your shell:
   ```bash
   source ~/.bashrc  # or ~/.zshrc, etc.
   ```

3. **Verify Installation**
   ```bash
   go version
   ```
