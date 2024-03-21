# goSemver

This is a simple Go application that helps manage semantic versioning for your projects.

## Installation

To install this application, you need to have Go installed on your machine. Once you have Go installed, you can clone this repository and build the application.

```bash
go install github.com/ptdave20/goSemver
```

## Usage

You can run the application with different command line arguments to increment the major, minor, or patch version. The application will print the new version to stdout.

```bash
goSemver break    # Increment the major version and print it
goSemver feature  # Increment the minor version and print it
goSemver          # Increment the patch version and print it
```

Using the command inline with git tags can be useful for managing versioning in your projects.

```bash
git tag -a $(goSemver break) -m "Breaking change"    # Create a new git tag with the incremented major version
git tag -a $(goSemver feature) -m "New feature"      # Create a new git tag with the incremented minor version
git tag -a $(goSemver) -m "Bug fix"                  # Create a new git tag with the incremented patch version
```

## Setting up the PATH

The `go install` command places the compiled executable in the `$GOPATH/bin` directory. To run the program from any location on your system, you need to add `$GOPATH/bin` to your system's `PATH`.

### On Linux or Mac

If you haven't done this already, you can add the following line to your shell profile file (like `.bashrc`, `.bash_profile`, or `.zshrc`):

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### On Windows
1. Right-click on Computer and click on Properties.
2. Click on Advanced system settings.
3. Click on Environment Variables.
4. Under System Variables, find and select the Path variable.
5. Click on Edit.
6. In the Edit Environment Variable window, click on New.
7. Paste the path to your Go binary folder, which is usually C:\Go\bin if you have used the default installation settings.
8. Click on OK in all windows to apply the changes.